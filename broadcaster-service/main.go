package main

import (
	"context"
	"flag"
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

	logger.Info("Started webserver at http://localhost:3551")
	err := http.ListenAndServe(":3551", nil)

	if err != nil {
		panic(err)
	}
}
