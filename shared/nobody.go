package shared

import "strings"

func EmptyResult(err error) bool {
	return err != nil && strings.Contains(err.Error(), "[scanner]: empty result")
}