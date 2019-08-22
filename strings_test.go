package lazygo

import (
	"testing"
)

func TestStrings(t *testing.T) {
	testString := "StRing"
	ToLower(&testString)
	if testString != "string" {
		t.Error("ToLower operation failed")
	}
	ToUpper(&testString)
	if testString != "STRING" {
		t.Error("ToLower operation failed")
	}
}
