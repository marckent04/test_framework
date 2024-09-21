package frontend

import (
	"cucumber/config"
	Browser "cucumber/utils/frontend"
	"fmt"
	"github.com/go-rod/rod"
	"log"
	"strings"
)

var testBrowser *rod.Browser
var Page *rod.Page

func iOpenAPrivateBrowserTab() {
	Browser.DestroyCurrentInstance()
	testBrowser = Browser.GetInstance()
}

func iNavigateToPage(page string) {
	url, err := config.GetPageUrl(page)
	if err != nil {
		log.Fatalf("Url for page %s not configured", page)
	}
	Page = testBrowser.MustPage(url)
	Page.MustWaitNavigation()
	Page.MustWaitIdle()

	if err != nil {
		log.Fatal(err)
	}

}

func iAmRedirectedToPage(page string) error {
	Page.MustWaitNavigation()
	Page.MustWaitDOMStable()
	url, err := config.GetPageUrl(page)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(Page.MustInfo().URL, url) {
		return fmt.Errorf("redirection failed")
	}

	return nil
}
