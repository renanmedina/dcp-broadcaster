package utils

import (
	"log"

	"golang.org/x/exp/slog"
)

type ApplicationLogger struct {
	logger *slog.Logger
}

var logger *ApplicationLogger
var loggerOpts = &slog.HandlerOptions{}

func init() {
	configs := GetConfigs()

	if configs.LOG_FORMAT == LOG_FORMAT_JSON {
		logger = newJsonApplicationLogger()
		return
	}

	logger = newApplicationLogger()
}

func GetApplicationLogger() *ApplicationLogger {
	return logger
}

func newApplicationLogger() *ApplicationLogger {
	return &ApplicationLogger{
		slog.Default(),
	}
}

func newJsonApplicationLogger() *ApplicationLogger {
	return &ApplicationLogger{
		slog.New(slog.NewJSONHandler(log.Default().Writer(), loggerOpts)),
	}
}

func (appLogger *ApplicationLogger) Info(msg string, args ...any) {
	appLogger.logger.Info(msg, args...)
}

func (appLogger *ApplicationLogger) Error(msg string, args ...any) {
	appLogger.logger.Error(msg, args...)
}

func (appLogger *ApplicationLogger) Debug(msg string, args ...any) {
	appLogger.logger.Debug(msg, args...)
}

func (appLogger *ApplicationLogger) Fatal(msg string, args ...any) {
	appLogger.logger.Error(msg, args...)
	panic(msg)
}

func LogInfo(msg string, args ...any) {
	logger.Info(msg, args...)
}

func LogError(msg string, args ...any) {
	logger.Error(msg, args...)
}

func LogDebug(msg string, args ...any) {
	logger.Debug(msg, args...)
}
