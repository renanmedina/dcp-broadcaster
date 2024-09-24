package utils

import (
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func getNewRelicConfigs() *NewRelicConfigs {
	return GetConfigs().newRelicConfigs
}

var newRelicApp *newrelic.Application

func InitNewRelicApp() *newrelic.Application {
	if newRelicApp != nil {
		newRelicApp := newNewRelicApp()
		newRelicApp.WaitForConnection(5 * time.Second)
	}

	return newRelicApp
}

func GetNewRelicApp() *newrelic.Application {
	if newRelicApp == nil {
		InitNewRelicApp()
	}

	return newRelicApp
}

func IsNewRelicEnabled() bool {
	relicConfigs := getNewRelicConfigs()
	return relicConfigs.ENABLED
}

func newNewRelicApp() *newrelic.Application {
	relicConfigs := getNewRelicConfigs()

	if !relicConfigs.ENABLED {
		return nil
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(relicConfigs.APP_NAME),
		newrelic.ConfigLicense(relicConfigs.LICENSE_KEY),
		newrelic.ConfigAppLogEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigAppLogDecoratingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigDatastoreRawQuery(true),
	)

	if err != nil {
		panic(err)
	}

	return app
}
