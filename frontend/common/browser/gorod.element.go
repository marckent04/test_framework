package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type rodElement struct {
	element *rod.Element
}

func (e *rodElement) Input(text string) error {
	err := e.element.Input(text)
	if err != nil {
		return err
	}
	return nil
}

func (e *rodElement) Click() error {
	return e.element.Click(proto.InputMouseButtonLeft, 1)
}

func (e *rodElement) TextContent() string {
	return e.element.MustText()
}

func (e *rodElement) IsVisible() bool {
	visible, err := e.element.Visible()
	if err != nil {
		return false
	}
	return visible
}

func newRodElement(element *rod.Element) Element {
	return &rodElement{element: element}
}
