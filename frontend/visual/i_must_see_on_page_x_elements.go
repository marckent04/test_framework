package visual

import (
	"cucumber/frontend/common"
	"fmt"
)

var iMustSeeOnPageXElements = common.FrontStep{
	Sentences: []string{`^I must see on page {number} {string}$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(expectedCount int, elementName string) error {
			elementCount := common.GetElementCount(ctx.GetCurrentPage(), elementName)
			if elementCount != expectedCount {
				return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
			}
			return nil
		}
	},
}
