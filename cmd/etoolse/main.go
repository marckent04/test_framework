package main

import (
	"etoolse/internal/actions"
	"etoolse/internal/config"
	"log"
)

func main() {
	cliConfig := config.Init()
	if cliConfig == nil {
		logger.Fatal("no mode specified")
	}

	modes := map[config.Mode]actions.Type{
		config.RunMode:        runMode,
		config.InitMode:       actions.Init,
		config.ValidationMode: actions.Validate,
	}

	if mode, ok := modes[cliConfig.Mode]; ok {
		mode(cliConfig)
	} else {
		logger.Fatalf("unknown mode: %s", cliConfig.Mode)
	}
}

func runMode(appConfig *config.App) {
	logger.Info("--- configuration resume ---")

	logger.Info("app name: ", appConfig.AppName)
	logger.Info("app description: ", appConfig.AppDescription)
	logger.Info("app version: ", appConfig.AppVersion)
	logger.Info("app tags: ", appConfig.Tags)
	logger.Info("app gherkin location: ", appConfig.GherkinLocation)
	logger.Info("app reporters format: ", appConfig.ReportFormat)
	logger.Info("app concurrency: ", appConfig.GetConcurrency())
	logger.Info("app slow motion: ", appConfig.GetSlowMotion())
	logger.Info("app test suite timeout: ", appConfig.Timeout)
	logger.Info("app headless mode: ", appConfig.IsHeadlessModeEnabled())

	log.Print("--- configuration resume end ---\n\n")

	actions.Run(appConfig)
}
