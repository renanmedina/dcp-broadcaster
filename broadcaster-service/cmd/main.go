package main

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	logger := utils.GetApplicationLogger()
	setup()
	migrate(logger)
	startWorker(30*time.Second, logger)
}

func startWorker(every time.Duration, logger *utils.ApplicationLogger) {
	receiver, err := daily_questions.NewQuestionsReceiver()

	if err != nil {
		logger.Fatal(err.Error())
	}

	receiver.Work(every)
}

func setup() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}

func migrate(logger *utils.ApplicationLogger) {
	db := utils.GetDatabase()
	logger.Info("Migrating database")
	err := db.Migrate("up", utils.GetConfigs().MIGRATIONS_PATH)

	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Migration success")
}
