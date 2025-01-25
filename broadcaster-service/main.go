package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hibiken/asynqmon"
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/monitoring"
	"github.com/renanmedina/dcp-broadcaster/task_queue"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	ctx := context.Background()
	monitoringProvider := monitoring.InitTracer()
	defer monitoringProvider.Shutdown(ctx)
	defer task_queue.GetTasksScheduler().Close()

	mode := setup()
	if mode == utils.MODE_WORKER {
		daily_questions.StartWorker(1 * time.Hour)
		return
	} else if mode == utils.MODE_QUEUE_WORKER {
		startQueueWork()
		return
	}

	startServer()
}

func setup() string {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	utils.MigrateDb("up")
	return utils.GetModeFlag()
}

func startQueueWork() {
	queueWorker := task_queue.InitializeQueueServer()
	// Register tasks to be processed in queue worker
	queueWorker.RegisterTaskProcessor(daily_questions.TypeSolveQuestionTask, daily_questions.SolveQuestionTaskProcessor)

	if err := queueWorker.Run(); err != nil {
		utils.GetApplicationLogger().Fatal(err.Error())
	}
}

func startServer() {
	logger := utils.GetApplicationLogger()
	configs := utils.GetConfigs()

	http.HandleFunc("/saveSolutionFile", func(w http.ResponseWriter, r *http.Request) {
		solutionId := r.URL.Query().Get("id")
		uc := daily_questions.NewStoreQuestionSolutionFile()
		uc.Execute(solutionId)
	})

	http.HandleFunc("/saveSolutionFiles", func(w http.ResponseWriter, r *http.Request) {
		repo := daily_questions.NewQuestionSolutionsRepository()
		solutions, err := repo.GetAll()
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		uc := daily_questions.NewStoreQuestionSolutionFile()
		for _, solution := range solutions {
			uc.Execute(solution.Id)
		}
	})

	http.HandleFunc("/solveQuestions", func(w http.ResponseWriter, r *http.Request) {
		repo := daily_questions.NewQuestionsRepository()
		questions, err := repo.GetAll()
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		handler := daily_questions.NewSolveQuestionEventHandler()
		for _, question := range questions {
			logger.Info(fmt.Sprintf("Solving question %s", question.Id))
			event := daily_questions.NewQuestionCreatedEvent(question)
			go handler.Handle(event)
		}
	})

	http.HandleFunc("/solveQuestion", func(w http.ResponseWriter, r *http.Request) {
		questionId := r.URL.Query().Get("id")
		repo := daily_questions.NewQuestionsRepository()
		question, err := repo.GetById(questionId)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		handler := daily_questions.NewSolveQuestionEventHandler()
		event := daily_questions.NewQuestionCreatedEvent(*question)
		handler.Handle(event)
	})

	asynqmonConfig := asynqmon.New(asynqmon.Options{
		RootPath:     "/tasks-queue", // RootPath specifies the root for asynqmon app
		RedisConnOpt: task_queue.GetQueueClientOptions(),
	})

	http.Handle(asynqmonConfig.RootPath()+"/", asynqmonConfig)

	logger.Info(fmt.Sprintf("Started webserver at http://localhost:%s", configs.WEBSERVER_PORT))
	addr := fmt.Sprintf(":%s", configs.WEBSERVER_PORT)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		panic(err)
	}
}
