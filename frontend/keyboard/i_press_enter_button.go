package keyboard

import (
	"cucumber/frontend/common"
)

var iPressTheEnterButton = common.FrontStep{
	Sentences: []string{`I press the enter button`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func() error {
			return ctx.GetCurrentPageKeyboard().PressEnter()
		}
	},
}
