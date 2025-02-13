package browser

import (
	"etoolse/internal/browser/common"
	"etoolse/internal/browser/rod"
	"etoolse/internal/config/testsconfig"
	"etoolse/pkg/logger"
	"fmt"
	"sync"
	"time"
)

func CreateInstance(headlessMode bool, timeout, slowMotion time.Duration, incognitoMode bool) common.Browser {
	return rod.New(headlessMode, timeout, slowMotion, incognitoMode)
}

func GetElement(page common.Page, label string) (common.Element, error) {
	selectors, err := testsconfig.GetHTMLElementSelectors(label)
	if err != nil {
		return nil, err
	}
	return GetElementBySelectors(page, selectors), nil
}

func GetElementBySelectors(page common.Page, potentialSelectors []string) common.Element {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var foundElement common.Element

	ch := make(chan common.Element, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		wg.Add(1)

		go func() {
			defer wg.Done()

			element, err := page.GetOneBySelector(selector)
			if err != nil {
				logger.Info(fmt.Sprintf("element not found with selector %s", selector))
				return
			}
			mu.Lock()
			if foundElement == nil {
				foundElement = element
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

	return foundElement
}

func GetActiveSelector(page common.Page, potentialSelectors []string) string {
	ch := make(chan string, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		go func() {
			exists := page.HasSelector(selector)
			if exists {
				ch <- selector
			}
		}()
	}

	return <-ch
}

func GetElementCount(page common.Page, label string) int {
	potentialSelectors, _ := testsconfig.GetHTMLElementSelectors(label)
	selector := GetActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		msg := fmt.Sprintf("Error getting elements with selector %s", selector)
		logger.Error(msg, []string{
			"Incorrect selector defined in the configuration file",
		}, []string{"Check the selector in the configuration file"})
	}

	return len(elements)
}
