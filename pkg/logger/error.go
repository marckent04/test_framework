package logger

import "github.com/charmbracelet/log"

func Error(msg string, potentialCauses []string, potentialSolutions []string) {
	log.Error(msg, potentialCauses, potentialSolutions)
}

func Fatal(context string, err error) {
	log.Fatal(err, context)
}
