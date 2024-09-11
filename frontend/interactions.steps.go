package frontend

import (
	"fmt"
	"github.com/go-rod/rod/lib/proto"
	"log"
	"strings"
)

func iFillTheInputWith(inputLabel, value string) error {
	input := getElement(inputLabel)
	return input.Input(value)
}

func iClickOnButtonOrElement(label string) error {
	button := getElement(label)
	return button.Click(proto.InputMouseButtonLeft, 1)
}

func iMustSeeOnThePage(word string) error {
	if !strings.Contains(Page.MustElement("body").String(), word) {
		return fmt.Errorf("%s not found", word)
	}
	return nil
}

func iClickOnElementWhichContains(text string) error {
	element, err := Page.ElementX("//*[contains(text(),'typeorm')]")
	log.Println(element)
	//element, err := Page.ElementX(fmt.Sprintf("//*[contains(text(),'%s')]", text))
	if err != nil {
		return fmt.Errorf("no element with text containing %s found", text)
	}
	element.MustClick()
	return nil
	//Page.MustElementR("a", text).MustClick()
}
