package main

import (
	"context"
	"cucumber/frontend"
	"cucumber/report"
	"github.com/cucumber/godog"
	"log"
	"os"
	"time"
)

func main() {
	const concurrency = 2

	file, err := os.Create("report.html")
	if err != nil {
		log.Panicln("cannot create report file in this folder")
	}
	defer file.Close()

	var opts = godog.Options{
		Output:              file,
		Concurrency:         concurrency,
		Format:              "pretty",
		NoColors:            true,
		ShowStepDefinitions: false,
		Paths:               []string{"features"},
	}

	testReport := report.New()

	testSuite := godog.TestSuite{
		Name:    "App",
		Options: &opts,
		TestSuiteInitializer: func(suiteContext *godog.TestSuiteContext) {
			suiteContext.BeforeSuite(func() {
				testReport.SetStartDate(time.Now())
			})

			suiteContext.AfterSuite(func() {

			})
		},
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			frontend.InitializeScenario(sc)

			sc.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
				log.Println("step :", st.Text, "status", status.String(), "error : ", err)
				return ctx, err
			})
			sc.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
				log.Println("laaa")
				/*			log.Println("----------------------------------")
							log.Println("scenario name ", sc.Name)
							godog.ScenarioContext{}.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
								log.Println("step : ", st.Text, "status", status.String(), "error : ", err)
								return ctx, err
							})

							log.Println("----------------------------------")
				*/
				return ctx, err

			})

		},
	}

	status := testSuite.Run()

	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
	}
}
