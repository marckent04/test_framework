package common

import (
	"github.com/go-rod/rod"
)

type Context struct {
	browser *rod.Browser
	page    *rod.Page
}

func (fc *Context) InitBrowser() {
	fc.browser = CreateInstance()
}

func (fc *Context) OpenNewPage(url string) {
	fc.page = fc.browser.MustPage(url)
	fc.page.MustWaitNavigation()
	fc.page.MustWaitIdle()
}

func (fc *Context) GetCurrentPage() *rod.Page {
	return fc.page
}

func NewFrontendContext() *Context {
	return &Context{
		page:    nil,
		browser: nil,
	}
}
