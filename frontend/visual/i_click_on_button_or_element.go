package visual

import (
	"cucumber/frontend/common"
	"github.com/go-rod/rod/lib/proto"
)

var iClickOnButtonOrElement = common.FrontStep{
	Sentences: []string{`^I click on "{string}" {string}$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(label string) error {
			button := common.GetElement(ctx.GetCurrentPage(), label)
			return button.Click(proto.InputMouseButtonLeft, 1)
		}
	},
}
