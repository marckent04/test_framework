package visual

import (
	"errors"
	"etoolse/internal/steps_definitions/core"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

func (s steps) iShouldSeeDetailsOnPage() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I should see "{string}" details on the page$`},
		func(ctx *core.TestSuiteContext) func(string, *godog.Table) error {
			return func(elementName string, table *godog.Table) error {
				data, parseErr := assistdog.NewDefault().ParseMap(table)
				if parseErr != nil {
					return errors.New("details malformed please go to the doc")
				}

				var errMsgs []string
				for name, value := range data {
					xPath := fmt.Sprintf("//*[contains(text(),\"%s\")]", value)
					elt, err := ctx.GetCurrentPage().GetOneByXPath(xPath)
					if err != nil {
						errMsgs = append(errMsgs, fmt.Sprintf("%s %s not found", elementName, name))
						continue
					}

					if !elt.IsVisible() {
						errMsgs = append(errMsgs, fmt.Sprintf("%s %s is found but is no visible", elementName, name))
					}
				}

				if len(errMsgs) > 0 {
					return errors.New(strings.Join(errMsgs, "\n"))
				}

				return nil
			}
		},
		nil,
	)
}
