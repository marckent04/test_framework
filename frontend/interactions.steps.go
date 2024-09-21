package frontend

import (
	"fmt"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"strings"
	"time"
)

func iFillTheInputWith(inputLabel, value string) error {
	input := getInputElement(inputLabel)
	return input.Input(value)
}

func iClickOnButtonOrElement(label string) error {
	button := getElement(label)
	return button.Click(proto.InputMouseButtonLeft, 1)
}

func iClickOnButtonOrElementIfExists(label string) error {
	button := getElement(label)
	if button == nil {
		return nil
	}

	return button.Click(proto.InputMouseButtonLeft, 1)
}

func iMustSeeOnThePage(word string) error {
	if !strings.Contains(Page.MustElement("body").String(), word) {
		return fmt.Errorf("%s not found", word)
	}
	return nil
}

func iClickOnElementWhichContains(text string) error {
	element, err := Page.ElementX(fmt.Sprintf("//*[contains(text(),%s)]", text))
	//element, err := Page.ElementX(fmt.Sprintf("//*[contains(text(),'%s')]", text))
	if err != nil {
		return fmt.Errorf("no element with text containing %s found", text)
	}
	element.MustClick()
	return nil
	//Page.MustElementR("a", text).MustClick()
}

func iPressTheEnterButton() error {
	return Page.Keyboard.Press(input.Enter)
}

func iMustSeeOnPageXElements(expectedCount int, elementName string) error {
	elementCount := getElementCount(elementName)
	if elementCount != expectedCount {
		return fmt.Errorf("%d %s expected but %d %s found", expectedCount, elementName, elementCount, elementName)
	}
	return nil
}

func iMustSeeOnPageAnElementWithText(elementLabel, text string) error {
	cases := map[string]string{
		"link":    "a",
		"button":  "button",
		"element": "*",
	}

	xPath := fmt.Sprintf("//%s[contains(text(),\"%s\")]", cases[elementLabel], text)
	element, err := Page.Timeout(2 * time.Second).ElementX(xPath)

	cErr := fmt.Errorf("no %s is visible with text \"%s\"", elementLabel, text)
	if err != nil {
		return cErr
	}

	visible, _ := element.Visible()
	if !visible {
		return cErr
	}

	return nil
}
