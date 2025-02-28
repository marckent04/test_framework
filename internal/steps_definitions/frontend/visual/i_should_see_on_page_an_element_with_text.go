package visual

import (
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
	"fmt"
)

func (s steps) iShouldSeeElementWhichContains() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I should see a (link|button|element) which contains "{string}"$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(elementLabel, text string) error {
				cases := map[string]string{
					"link":    "a",
					"button":  "button",
					"element": "*",
				}

				xPath := fmt.Sprintf("//%s[contains(text(),\"%s\")]", cases[elementLabel], text)
				element, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
				cErr := fmt.Errorf("no %s is visible with text \"%s\"", elementLabel, text)
				if err != nil {
					return cErr
				}

				visible := element.IsVisible()
				if !visible {
					return cErr
				}

				return nil
			}
		},
		nil,
		core.StepDefDocParams{
			Description: "checks if a link, button or element is visible and contains a specific text.",
			Variables: []shared.StepVariable{
				{Name: "elementLabel", Description: "The label of the element to check.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text that the element should contain.", Type: shared.DocVarTypeString},
			},
			Example:  "Then I should see a button which contains \"Submit\"",
			Category: shared.Visual,
		},
	)
}
