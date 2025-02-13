package actions

import (
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/logger"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Validate(appConfig *config.App) {
	logger.Info("Validate gherkin files ...")

	const concurrency = 5
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         concurrency,
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                appConfig.Tags,
		Paths:               []string{appConfig.GherkinLocation},
	}

	ctx := core.ValidatorContext{}
	testSuite := godog.TestSuite{
		Name:                 "validate",
		Options:              &opts,
		ScenarioInitializer:  validateScenarioInitializer(&ctx),
		TestSuiteInitializer: validateTestSuiteInitializer(&ctx),
	}

	status := testSuite.Run()
	if status != 0 {
		logger.Fatal(fmt.Sprintf("zero status code expected, %d received", status), nil)
	}
}

func validateScenarioInitializer(ctx *core.ValidatorContext) func(*godog.ScenarioContext) {
	logger.Info("Initializing scenarios for validation ...")

	return func(sc *godog.ScenarioContext) {
		frontend.InitValidationScenarios(sc, ctx)
	}
}

func validateTestSuiteInitializer(validatorCtx *core.ValidatorContext) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.AfterSuite(func() {
			if validatorCtx.HasErrors() {
				logger.Error("Elements validation failed", []string{
					"Elements variables malformed in gherkin files",
					"Elements variables must be defined in the config file",
				}, []string{
					"Verify the elements variables in the gherkin files",
				})
				logger.Info(validatorCtx.GetElementsErrorsFormatted())
			}
		})
	}
}

