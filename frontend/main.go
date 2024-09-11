package frontend

import (
	"cucumber/config"
	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	config.InitializeFrontConfig()

	ctx.Step(`I open a new private browser tab`, iOpenAPrivateBrowserTab)
	ctx.Step(`^I am redirected to ([^"]*) page$`, iAmRedirectedToPage)
	ctx.Step(`^I click on "([^"]*)" ([^"]*)$`, iClickOnButtonOrElement)
	ctx.Step(`^I fill the ([^"]*) input with "([^"]*)"$`, iFillTheInputWith)
	ctx.Step(`^I navigate to ([^"]*) page$`, iNavigateToPage)
	ctx.Step(`^I must see ([^"]*) on the page$`, iMustSeeOnThePage)
	ctx.Step(`^I click on element which contains "([^"]*)"$`, iClickOnElementWhichContains)
}
