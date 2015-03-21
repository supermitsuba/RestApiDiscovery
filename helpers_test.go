// handlers_test
package main

import (
	"testing"
)

func TestIsStringWhitespaceOrEmpty_EmptyString(t *testing.T) {
	var result = IsStringWhiteSpaceOrEmpty("")
	if !result {
		t.Errorf("Should have said true.  It is empty")
	}
}

func TestIsStringWhitespaceOrEmpty_WhitespaceString(t *testing.T) {
	var result = IsStringWhiteSpaceOrEmpty("	  ")
	if !result {
		t.Errorf("Should have said true.  It is empty")
	}
}

func TestIsStringWhitespaceOrEmpty_RegularString(t *testing.T) {
	var result = IsStringWhiteSpaceOrEmpty("Hello World")
	if result {
		t.Errorf("Should have said false.  There is a string here.")
	}
}
