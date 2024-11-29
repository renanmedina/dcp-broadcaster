package utils

import (
	"log"

	"log/slog"
)

type ApplicationLogger struct {
	envName string
	logger  *slog.Logger
}

var logger *ApplicationLogger
var loggerOpts = &slog.HandlerOptions{}

func init() {
	configs := GetConfigs()

	if configs.LOG_FORMAT == LOG_FORMAT_JSON {
		logger = newJsonApplicationLogger(configs.ENVIRONMENT)
		return
	}

	logger = newApplicationLogger(configs.ENVIRONMENT)
}

func GetApplicationLogger() *ApplicationLogger {
	return logger
}

func newApplicationLogger(envName string) *ApplicationLogger {
	return &ApplicationLogger{
		envName,
		slog.Default(),
	}
}

func newJsonApplicationLogger(envName string) *ApplicationLogger {
	jsonHandler := slog.NewJSONHandler(log.Default().Writer(), loggerOpts)

	return &ApplicationLogger{
		envName,
		slog.New(jsonHandler),
	}
}

func (appLogger *ApplicationLogger) addEnv(args []any) []any {
	args = append(args, "environment")
	args = append(args, appLogger.envName)
	return args
}

func (appLogger *ApplicationLogger) Info(msg string, args ...any) {
	appLogger.logger.Info(msg, appLogger.addEnv(args)...)
}

func (appLogger *ApplicationLogger) Error(msg string, args ...any) {
	appLogger.logger.Error(msg, appLogger.addEnv(args)...)
}

func (appLogger *ApplicationLogger) Debug(msg string, args ...any) {
	appLogger.logger.Debug(msg, appLogger.addEnv(args)...)
}

func (appLogger *ApplicationLogger) Fatal(msg string, args ...any) {
	appLogger.logger.Error(msg, appLogger.addEnv(args)...)
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
