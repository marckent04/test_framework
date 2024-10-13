package main

import (
	"cucumber/frontend"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"log"
	"os"
)

func main() {

	var opts = godog.Options{
		Output:      colors.Colored(os.Stdout),
		Concurrency: 2,
		Format:      "pretty",
		Paths:       []string{"features"},
	}

	testSuite := godog.TestSuite{
		Name:    "App",
		Options: &opts,
		ScenarioInitializer: func(context *godog.ScenarioContext) {
			frontend.InitializeScenario(context)
		},
	}

	status := testSuite.Run()

	if status != 0 {
		log.Fatalf("zero status code expected, %d received", status)
	}

}
