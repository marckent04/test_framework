package frontend

import (
	"cucumber/config"
	"github.com/go-rod/rod"
)

func getElement(label string) (elt *rod.Element) {
	selectors := config.GetElementSelectors(label)
	ch := make(chan *rod.Element, 1)
	for _, selector := range selectors {
		selector := selector
		go func() {
			ch <- Page.MustElement(selector)
		}()
	}
	elt = <-ch
	close(ch)
	return
}
