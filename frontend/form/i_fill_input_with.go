package form

import (
	"cucumber/frontend/common"
	"cucumber/frontend/common/browser"
)

var iFillTheInputWith = common.FrontStep{
	Sentences: []string{`^I fill the {string} input with "{string}"$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(inputLabel, value string) error {
			input := browser.GetInputElement(ctx.GetCurrentPage(), inputLabel)
			return input.Input(value)
		}
	},
}
