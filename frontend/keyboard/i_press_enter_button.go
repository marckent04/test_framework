package keyboard

import (
	"cucumber/frontend/common"
	"github.com/go-rod/rod/lib/input"
)

var iPressTheEnterButton = common.FrontStep{
	Sentences: []string{`I press the enter button`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func() error {
			return ctx.GetCurrentPage().Keyboard.Press(input.Enter)
		}
	},
}
