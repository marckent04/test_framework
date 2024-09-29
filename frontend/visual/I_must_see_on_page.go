package visual

import (
	"cucumber/frontend/common"
	"fmt"
	"strings"
)

var iMustSeeOnThePage = common.FrontStep{
	Sentences: []string{`^I must see ([^"]*) on the page$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(word string) error {
			if !strings.Contains(ctx.GetCurrentPage().MustElement("body").String(), word) {
				return fmt.Errorf("%s not found", word)
			}
			return nil
		}
	},
}
