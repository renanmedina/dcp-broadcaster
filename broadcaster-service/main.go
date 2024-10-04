package main

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	if utils.IsNewRelicEnabled() {
		utils.InitNewRelicApp()
		// defer app.Shutdown(10 * time.Second)
	}

	setup()
	utils.MigrateDb("up")
	daily_questions.StartWorker(1 * time.Hour)
}

func setup() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}
