package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

type rodKeyboard struct {
	keyboard *rod.Keyboard
}

func (k *rodKeyboard) PressEnter() error {
	return k.keyboard.Press(input.Enter)
}

func newRodKeyboard(keyboard *rod.Keyboard) Keyboard {
	return &rodKeyboard{keyboard: keyboard}
}
