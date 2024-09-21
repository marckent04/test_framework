package frontend

import (
	"cucumber/config"
	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	// TODO:  le mettre un poil plus haut car doit etre exucuté qu'une seule fois
	config.InitializeFrontConfig()

	ctx.Step(`I open a new private browser tab`, iOpenAPrivateBrowserTab)
	ctx.Step(`^I am redirected to ([^"]*) page$`, iAmRedirectedToPage)
	ctx.Step(`^I click on "([^"]*)" ([^"]*)$`, iClickOnButtonOrElement)
	ctx.Step(`^I click on "([^"]*)" ([^"]*) if exists$`, iClickOnButtonOrElementIfExists)
	ctx.Step(`^I fill the ([^"]*) input with "([^"]*)"$`, iFillTheInputWith)
	ctx.Step(`^I navigate to ([^"]*) page$`, iNavigateToPage)
	ctx.Step(`^I must see ([^"]*) on the page$`, iMustSeeOnThePage)
	ctx.Step(`^I press the enter button$`, iPressTheEnterButton)
	ctx.Step(`^I must see on page (\d+) ([^"]*)$`, iMustSeeOnPageXElements)
	ctx.Step(`^I click on element which contains "([^"]*)"$`, iClickOnElementWhichContains)
	ctx.Step(`^I must see on page a (link|button|element) with text "([^"]*)"$`, iMustSeeOnPageAnElementWithText)

}
