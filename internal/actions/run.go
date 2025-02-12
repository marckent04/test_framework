package actions

import (
	"context"
	"etoolse/internal/config"
	"etoolse/internal/steps_definitions/frontend"
	"etoolse/pkg/gherkinparser"
	"etoolse/pkg/reporters"
	"log"
	"time"

	"github.com/cucumber/godog"
	"github.com/tdewolff/parse/buffer"
)

func Run(appConfig *config.App) {
	log.Println("Starting tests execution ...")

	parsedFeatures := gherkinparser.Parse(appConfig.GherkinLocation)
	features := make([]godog.Feature, len(parsedFeatures))
	for i, f := range parsedFeatures {
		features[i] = godog.Feature{
			Name:     f.Name,
			Contents: f.Contents,
		}
	}

	testReport := reporters.New(appConfig.AppName, appConfig.AppDescription, appConfig.ReportFormat)
	var opts = godog.Options{
		Output:              &buffer.Writer{},
		Concurrency:         appConfig.GetConcurrency(),
		Format:              "pretty",
		ShowStepDefinitions: false,
		Tags:                appConfig.Tags,
		FeatureContents:     features,
	}

	testSuite := godog.TestSuite{
		Name:                 appConfig.AppName,
		Options:              &opts,
		TestSuiteInitializer: testSuiteInitializer(&testReport),
		ScenarioInitializer:  scenarioInitializer(appConfig, &testReport),
	}

	log.Println("Running tests ...")
	status := testSuite.Run()
	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
	}
}

func testSuiteInitializer(testReport *reporters.Report) func(*godog.TestSuiteContext) {
	return func(suiteContext *godog.TestSuiteContext) {
		suiteContext.BeforeSuite(func() {
			testReport.Start()
		})

		suiteContext.AfterSuite(func() {
			if testReport.HasScenarios() {
				testReport.Write()
				log.Println("Tests execution finished")
			} else {
				log.Println("No scenarios executed")
			}
		})
	}
}
func scenarioInitializer(config *config.App, testReport *reporters.Report) func(*godog.ScenarioContext) {
	log.Println("Initializing scenario for test running ...")
	return func(sc *godog.ScenarioContext) {
		frontend.InitTestRunnerScenarios(sc, config)
		myCtx := newScenarioCtx()
		sc.StepContext().Before(beforeStepHookInitializer(&myCtx))
		sc.StepContext().After(afterStepHookInitializer(&myCtx))
		sc.After(afterScenarioHookInitializer(testReport, &myCtx))
	}
}
func afterStepHookInitializer(myCtx *myScenarioCtx) godog.AfterStepHook {
	return func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		myCtx.addStep(st.Text, status, err)
		return ctx, err
	}
}

func beforeStepHookInitializer(myCtx *myScenarioCtx) godog.BeforeStepHook {
	return func(ctx context.Context, _ *godog.Step) (context.Context, error) {
		myCtx.currentStepStartTime = time.Now()
		return ctx, nil
	}
}

func afterScenarioHookInitializer(testReport *reporters.Report, myCtx *myScenarioCtx) godog.AfterScenarioHook {
	return func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		myCtx.scenarioReport.SetTitle(sc.Name)

		myCtx.scenarioReport.End()
		testReport.AddScenario(myCtx.scenarioReport)
		return ctx, err
	}
}

func newScenarioCtx() myScenarioCtx {
	return myScenarioCtx{
		scenarioReport: reporters.NewScenario(),
	}
}

type myScenarioCtx struct {
	currentStepStartTime time.Time
	scenarioReport       reporters.Scenario
}

func (c *myScenarioCtx) addStep(title string, status godog.StepResultStatus, err error) {
	c.scenarioReport.AddStep(title, status, time.Since(c.currentStepStartTime), err)
}
