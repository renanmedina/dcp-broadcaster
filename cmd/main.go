package main

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

func main() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	logger := utils.GetApplicationLogger()
	logger.Info("Starting Database connection")
	db := utils.GetDatabase()
	logger.Info("Database connection started successfully")

	logger.Info("Migrating database")
	err := db.Migrate("up")

	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Migration success")

	logger.Info("Starting questions receiver worker")
	receiver, err := daily_questions.NewQuestionsReceiver()

	if err != nil {
		logger.Error(err.Error())
	}

	receiver.Work()
	logger.Info("Questions receiver worker finished successfully")
}
