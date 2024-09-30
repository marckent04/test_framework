package form

import (
	"cucumber/frontend/common"
)

var iFillTheInputWith = common.FrontStep{
	Sentences: []string{`^I fill the {string} input with "{string}"$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(inputLabel, value string) error {
			input := common.GetInputElement(ctx.GetCurrentPage(), inputLabel)
			return input.Input(value)
		}
	},
}
