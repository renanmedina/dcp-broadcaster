package utils

import (
	"context"
	"log"

	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type ApplicationLogger struct {
	envName        string
	logger         *slog.Logger
	currentContext context.Context
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
		nil,
	}
}

func newJsonApplicationLogger(envName string) *ApplicationLogger {
	jsonHandler := slog.NewJSONHandler(log.Default().Writer(), loggerOpts)

	return &ApplicationLogger{
		envName,
		slog.New(jsonHandler),
		nil,
	}
}

func (appLogger *ApplicationLogger) addAppAttributes(args []any) []any {
	configs := GetConfigs()
	args = append(args, "environment", appLogger.envName)
	args = append(args, "service_name", configs.SERVICE_NAME)

	if appLogger.currentContext != nil {
		spanCtx := trace.SpanContextFromContext(appLogger.currentContext)

		if spanCtx.HasTraceID() {
			args = append(args, "trace_id", spanCtx.TraceID())
		}

		if spanCtx.HasSpanID() {
			args = append(args, "span_id", spanCtx.SpanID())
		}
	}

	return args
}

func (appLogger *ApplicationLogger) WithContext(ctx context.Context) *ApplicationLogger {
	appLogger.currentContext = ctx
	return appLogger
}

func (appLogger *ApplicationLogger) Info(msg string, args ...any) {
	appLogger.logger.Info(msg, appLogger.addAppAttributes(args)...)
}

func (appLogger *ApplicationLogger) Error(msg string, args ...any) {
	appLogger.logger.Error(msg, appLogger.addAppAttributes(args)...)
}

func (appLogger *ApplicationLogger) Debug(msg string, args ...any) {
	appLogger.logger.Debug(msg, appLogger.addAppAttributes(args)...)
}

func (appLogger *ApplicationLogger) Fatal(msg string, args ...any) {
	appLogger.logger.Error(msg, appLogger.addAppAttributes(args)...)
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
