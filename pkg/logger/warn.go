package logger

import "github.com/charmbracelet/log"

func Warn(msg, actionExpected string) {
	log.Warn(msg, actionExpected)
}
