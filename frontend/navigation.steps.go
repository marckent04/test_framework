package frontend

import (
	"cucumber/config"
	Browser "cucumber/utils/frontend"
	"github.com/go-rod/rod"
	"log"
	"strings"
	"time"
)

var Tab *rod.Browser
var Page *rod.Page

func iOpenAPrivateBrowserTab() {
	Tab = Browser.GetInstance()
}

func iNavigateToPage(page string) {
	url, err := config.GetPageUrl(page)
	if err != nil {
		log.Fatalf("Url for page %s not configured", page)
	}
	Page = Tab.MustPage(url)
	Page.MustWaitStable()
	Page.MustWaitIdle()
	Page.MustWaitNavigation()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
}

func iAmRedirectedToPage(page string) {
	Page.MustWaitDOMStable()
	url, err := config.GetPageUrl(page)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasPrefix(Page.MustInfo().URL, url) {
		log.Fatal("redirection failed")
	}
}
