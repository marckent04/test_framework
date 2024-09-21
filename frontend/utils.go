package frontend

import (
	"cucumber/config"
	"github.com/go-rod/rod"
	"log"
)

func getElement(label string) *rod.Element {
	selectors := config.GetElementSelectors(label)
	return getElementBySelectors(selectors)
}

func getInputElement(label string) (elt *rod.Element) {
	selectors := config.GetInputSelectors(label)
	return getElementBySelectors(selectors)
}

func getElementBySelectors(potentialSelectors []string) *rod.Element {
	ch := make(chan *rod.Element, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		selector := selector
		go func() {
			element, _ := Page.Element(selector)
			ch <- element
		}()
	}

	return <-ch
}

func getActiveSelector(potentialSelectors []string) string {
	ch := make(chan string, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		selector := selector
		go func() {
			exists, _, _ := Page.Has(selector)
			if exists {
				ch <- selector
			}
		}()
	}

	return <-ch
}

func getElementCount(label string) int {
	potentialSelectors := config.GetElementSelectors(label)
	selector := getActiveSelector(potentialSelectors)
	elements, err := Page.Elements(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
}
