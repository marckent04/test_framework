package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldInitializeAppConfig(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			GherkinLocation:    "features",
			ClIConfigPath:      "cli.yml",
			FrontendConfigPath: "frontend.yml",
			Tags:               "tags",
			Parallel:           10,
			Timeout:            15 * time.Second,
			Headless:           true,
			AppVersion:         "1.0",
		},
	}

	appConfigFile := appFileConfig{
		AppName:         "appName",
		AppDescription:  "appDescription",
		Timeout:         "10s",
		SlowMotion:      "2s",
		GherkinLocation: "features",
		reportingConfig: reportingConfig{
			ReportFormat: "html",
		},
	}

	appConfig := initAppConfig(appArgs, appConfigFile)

	assert.Equal(t, "appName", appConfig.AppName)
	assert.Equal(t, "appDescription", appConfig.AppDescription)
	assert.Equal(t, "1.0", appConfig.AppVersion)
	assert.Equal(t, "html", appConfig.ReportFormat)
	assert.Equal(t, "15s", appConfig.Timeout)
	assert.Equal(t, "features", appConfig.GherkinLocation)
	assert.Equal(t, "tags", appConfig.Tags)
	assert.Equal(t, 10, appConfig.Parallel)
	assert.True(t, appConfig.Headless)
	assert.Equal(t, "2s", appConfig.SlowMotion)
}

func TestShouldDefineConcurrencyTo0BecauseHeadlessIsDisabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Parallel: 10,
			Headless: false,
		},
	}

	appConfigFile := appFileConfig{}

	assert.False(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
	assert.Equal(t, 0, initAppConfig(appArgs, appConfigFile).GetConcurrency())
}

func TestShouldDefineConcurrencyTo10BecauseHeadlessIsEnabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Parallel: 10,
			Headless: true,
		},
	}

	appConfigFile := appFileConfig{}

	assert.True(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
	assert.Equal(t, 10, initAppConfig(appArgs, appConfigFile).GetConcurrency())
}

func TestShouldDefineSlowMotionTo0BecauseHeadlessIsEnabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Headless: true,
		},
	}

	appConfigFile := appFileConfig{
		SlowMotion: "2s",
	}

	assert.True(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
	assert.Equal(t, time.Duration(0), initAppConfig(appArgs, appConfigFile).GetSlowMotion())
}

func TestShouldDefineSlowMotionTo20sBecauseHeadlessIsDisabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Headless: false,
		},
	}

	appConfigFile := appFileConfig{
		SlowMotion: "20s",
	}

	assert.False(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
	assert.Equal(t, 20*time.Second, initAppConfig(appArgs, appConfigFile).GetSlowMotion())
}

func TestShouldHeadlessModeEnabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Headless: true,
		},
	}

	appConfigFile := appFileConfig{}

	assert.True(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
}

func TestShouldHeadlessModeDisabled(t *testing.T) {
	appArgs := appArgsConfig{
		Run: &RunCmd{
			Headless: false,
		},
	}

	appConfigFile := appFileConfig{}

	assert.False(t, initAppConfig(appArgs, appConfigFile).IsHeadlessModeEnabled())
}
