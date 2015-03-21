package main

import (
	"strings"
)

func IsStringWhiteSpaceOrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
