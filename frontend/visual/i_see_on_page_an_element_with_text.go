package visual

import (
	"cucumber/frontend/common"
	"fmt"
	"time"
)

var iMustSeeOnPageAnElementWithText = common.FrontStep{
	Sentences: []string{`^I must see on page a (link|button|element) with text "([^"]*)"$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(elementLabel, text string) error {
			cases := map[string]string{
				"link":    "a",
				"button":  "button",
				"element": "*",
			}

			xPath := fmt.Sprintf("//%s[contains(text(),\"%s\")]", cases[elementLabel], text)
			element, err := ctx.GetCurrentPage().Timeout(2 * time.Second).ElementX(xPath)

			cErr := fmt.Errorf("no %s is visible with text \"%s\"", elementLabel, text)
			if err != nil {
				return cErr
			}

			visible, _ := element.Visible()
			if !visible {
				return cErr
			}

			return nil
		}
	},
}
