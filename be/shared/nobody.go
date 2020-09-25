package shared

import (
	"regexp"
	"strings"
)

func EmptyResult(err error) bool {
	return err != nil && strings.Contains(err.Error(), "[scanner]: empty result")
}

var wordExp = regexp.MustCompile("^[A-Za-z]+$")

func IsEnglishWord(s string) bool {
	return wordExp.MatchString(s)
}
