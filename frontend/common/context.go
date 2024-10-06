package common

import (
	"cucumber/frontend/common/browser"
)

type Context struct {
	browser browser.Browser
	page    browser.Page
}

func (fc *Context) InitBrowser() {
	fc.browser = browser.CreateInstance()
}

func (fc *Context) OpenNewPage(url string) {
	fc.page = fc.browser.NewPage(url)
}

func (fc *Context) GetCurrentPage() browser.Page {
	return fc.page
}

func (fc *Context) GetCurrentPageKeyboard() browser.Keyboard {
	return fc.page.GetKeyboard()
}

func NewFrontendContext() *Context {
	return &Context{
		page:    nil,
		browser: nil,
	}
}
