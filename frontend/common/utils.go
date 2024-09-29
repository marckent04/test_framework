package common

import (
	"cucumber/config"
	"github.com/go-rod/rod"
	"log"
)

func GetElement(page *rod.Page, label string) *rod.Element {
	selectors := config.GetElementSelectors(label)
	return GetElementBySelectors(page, selectors)
}

func GetInputElement(page *rod.Page, label string) (elt *rod.Element) {
	selectors := config.GetInputSelectors(label)
	return GetElementBySelectors(page, selectors)
}

func GetElementBySelectors(page *rod.Page, potentialSelectors []string) *rod.Element {
	ch := make(chan *rod.Element, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		selector := selector
		go func() {
			element, _ := page.Element(selector)
			ch <- element
		}()
	}

	return <-ch
}

func GetActiveSelector(page *rod.Page, potentialSelectors []string) string {
	ch := make(chan string, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		selector := selector
		go func() {
			exists, _, _ := page.Has(selector)
			if exists {
				ch <- selector
			}
		}()
	}

	return <-ch
}

func GetElementCount(page *rod.Page, label string) int {
	potentialSelectors := config.GetElementSelectors(label)
	selector := GetActiveSelector(page, potentialSelectors)
	elements, err := page.Elements(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
}
