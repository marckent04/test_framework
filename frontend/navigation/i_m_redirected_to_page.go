package navigation

import (
	"cucumber/config"
	"cucumber/frontend/common"
	"fmt"
	"strings"
)

var iAmRedirectedToPage = common.FrontStep{
	Sentences: []string{"^I am redirected to {string} page$"},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(pageName string) error {
			page := ctx.GetCurrentPage()
			page.MustWaitNavigation()
			page.MustWaitDOMStable()
			url, err := config.GetPageUrl(pageName)
			if err != nil {
				return err
			}

			if !strings.HasPrefix(page.MustInfo().URL, url) {
				return fmt.Errorf("redirection failed")
			}

			return nil
		}
	},
}
