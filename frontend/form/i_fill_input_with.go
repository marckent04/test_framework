package form

import (
	"cucumber/frontend/common"
)

var iFillTheInputWith = common.FrontStep{
	Sentences: []string{`^I fill the ([^"]*) input with "([^"]*)"$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(inputLabel, value string) error {
			input := common.GetInputElement(ctx.GetCurrentPage(), inputLabel)
			return input.Input(value)
		}
	},
}
