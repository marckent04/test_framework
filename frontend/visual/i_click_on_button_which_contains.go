package visual

import (
	"cucumber/frontend/common"
	"fmt"
)

var iClickOnElementWhichContains = common.FrontStep{
	Sentences: []string{`^I click on element which contains "([^"]*)"$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(text string) error {
			element, err := ctx.GetCurrentPage().ElementX(fmt.Sprintf("//*[contains(text(),%s)]", text))
			//element, err := ctx.page.ElementX(fmt.Sprintf("//*[contains(text(),'%s')]", text))
			if err != nil {
				return fmt.Errorf("no element with text containing %s found", text)
			}
			element.MustClick()
			return nil
		}
	},
}
