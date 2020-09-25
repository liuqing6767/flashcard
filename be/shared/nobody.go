package shared

import (
	"strings"
	"unicode"
)

func EmptyResult(err error) bool {
	return err != nil && strings.Contains(err.Error(), "[scanner]: empty result")
}

func IsEnglishWord(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}

	return true
}
