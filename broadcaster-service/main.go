package main

import (
	"context"
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/monitoring"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	ctx := context.Background()
	monitoringProvider := monitoring.InitTracer()
	defer monitoringProvider.Shutdown(ctx)

	setup()
	utils.MigrateDb("up")
	daily_questions.StartWorker(1 * time.Hour)
}

func setup() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}
