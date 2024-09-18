package utils

import "github.com/newrelic/go-agent/v3/newrelic"

func getNewRelicConfigs() *NewRelicConfigs {
	return GetConfigs().newRelicConfigs
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
		newrelic.ConfigFromEnvironment(),
	)

	if err != nil {
		panic(err)
	}

	return app
}
