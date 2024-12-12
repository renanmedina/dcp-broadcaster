package main

import (
	"context"
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/monitoring"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	ctx := context.Background()
	monitoringProvider := monitoring.InitTracer()
	defer monitoringProvider.Shutdown(ctx)

	setup()
	utils.MigrateDb("up")
	// daily_questions.StartWorker(1 * time.Hour)

	id := "18f62eec-a601-4d13-8bae-ed044de868e2"
	uc := daily_questions.NewSolveQuestion(
		questions_solver.NewOllamaService(),
	)

	uc.Execute(id, "golang")
}

func setup() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}
