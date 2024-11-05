package form

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
	"fmt"
	"reflect"
)

func (s steps) checkCheckboxStatus() common.FrontStep {
	return common.FrontStep{
		Sentences: []string{`^the {string} checkbox should be (checked|unchecked)`},
		Definition: func(ctx *common.TestSuiteContext) common.FrontStepDefinition {
			return func(checkboxId, status string) error {
				input := browser.GetElement(ctx.GetCurrentPage(), fmt.Sprintf("%s_checkbox", checkboxId))

				checkValue, isBoolean := input.GetPropertyValue("checked", reflect.Bool).(bool)

				if isBoolean && checkValue && status == "checked" || !checkValue && status == "unchecked" {
					return nil
				}

				return fmt.Errorf("%s checkbox is not %s", checkboxId, status)
			}
		},
	}
}
