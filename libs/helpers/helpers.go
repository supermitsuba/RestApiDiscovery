package helpers

import (
	"strings"
)

func IsStringWhiteSpaceOrEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
