package main

import "github.com/renanmedina/cdp-broadcaster/utils"

func main() {
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
}
