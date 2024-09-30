package utils

import (
	"strings"
)

type Wildcard = string
type WildcardID = string

const (
	StringId WildcardID = "{string}"
	NumberId            = "{number}"
)

const (
	NumberWildcard Wildcard = `(\d+)`
	StringWildcard          = `([^"]*)`
)

var wildcards = map[WildcardID]Wildcard{
	NumberId: NumberWildcard,
	StringId: StringWildcard,
}

func ConvertWildcards(current string) string {
	for id, wildcard := range wildcards {
		current = strings.ReplaceAll(current, id, wildcard)
	}
	return current
}
