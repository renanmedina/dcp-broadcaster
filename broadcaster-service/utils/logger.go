package utils

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"

	"log/slog"
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
	newRelicApp := newNewRelicApp()

	if newRelicApp != nil {
		nrJsonHandler := nrslog.JSONHandler(newRelicApp, os.Stdout, &slog.HandlerOptions{})
		slog.SetDefault(slog.New(nrJsonHandler))
		return &ApplicationLogger{
			slog.New(nrJsonHandler),
		}
	}

	return &ApplicationLogger{
		slog.New(slog.NewJSONHandler(log.Default().Writer(), loggerOpts)),
	}
}

func newNewRelicApp() *newrelic.Application {
	relicConfigs := GetNewRelicConfigs()

	if !relicConfigs.ENABLED {
		return nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(relicConfigs.APP_NAME),
		newrelic.ConfigLicense(relicConfigs.LICENSE_KEY),
		newrelic.ConfigAppLogDecoratingEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(false),
	)

	if err != nil {
		panic(err)
	}

	return app
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
