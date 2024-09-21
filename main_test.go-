package main_test

import (
	"cucumber/frontend"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var opts = godog.Options{
	Output:      colors.Colored(os.Stdout),
	Concurrency: 1,
	Format:      "pretty",
	Paths:       []string{"features"},
}

func initializeScenarios(ctx *godog.ScenarioContext) {
	frontend.InitializeScenario(ctx)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                "App",
		Options:             &o,
		ScenarioInitializer: initializeScenarios,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}
