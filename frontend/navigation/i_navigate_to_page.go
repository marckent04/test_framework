package navigation

import (
	"cucumber/config"
	"cucumber/frontend/common"
	"log"
)

var iNavigateToPage = common.FrontStep{
	Sentences: []string{`^I navigate to ([^"]*) page$`},
	Definition: func(ctx *common.Context) common.FrontStepDefinition {
		return func(page string) error {
			url, err := config.GetPageUrl(page)
			if err != nil {
				log.Fatalf("Url for page %s not configured", page)
				return err
			}
			ctx.OpenNewPage(url)
			if err != nil {
				log.Fatal(err)
			}

			return nil
		}
	},
}
