// handlers_test
package helpers_tests

import (
	helpers "RestApiDiscovery/libs/helpers"
	"testing"
)

func TestIsStringWhitespaceOrEmpty_EmptyString_True(t *testing.T) {
	var result = helpers.IsStringWhiteSpaceOrEmpty("")
	if !result {
		t.Errorf("Should have said true.  It is empty")
	}
}

func TestIsStringWhitespaceOrEmpty_WhitespaceString_True(t *testing.T) {
	var result = helpers.IsStringWhiteSpaceOrEmpty("	  ")
	if !result {
		t.Errorf("Should have said true.  It is empty")
	}
}

func TestIsStringWhitespaceOrEmpty_RegularString_False(t *testing.T) {
	var result = helpers.IsStringWhiteSpaceOrEmpty("Hello World")
	if result {
		t.Errorf("Should have said false.  There is a string here.")
	}
}
