package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/monitoring"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	ctx := context.Background()
	monitoringProvider := monitoring.InitTracer()
	defer monitoringProvider.Shutdown(ctx)

	mode := setup()
	if mode == "worker" {
		daily_questions.StartWorker(1 * time.Hour)
		return
	}
	startServer()
}

func setup() string {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	utils.MigrateDb("up")
	mode := flag.String("mode", "worker", "worker or webserver")
	flag.Parse()
	return *mode
}

func startServer() {
	logger := utils.GetApplicationLogger()

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

		handler := daily_questions.NewSolveQuestionHandler()
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

		handler := daily_questions.NewSolveQuestionHandler()
		event := daily_questions.NewQuestionCreatedEvent(*question)
		handler.Handle(event)
	})

	logger.Info("Started webserver at http://localhost:3551")
	err := http.ListenAndServe(":3551", nil)

	if err != nil {
		panic(err)
	}
}
