package lazygo

import (
	"testing"

	"gotest.tools/assert"
)

func TestStrings(t *testing.T) {
	testString := "StRing"
	ToLower(&testString)
	assert.Equal(t, "string", testString)
	ToUpper(&testString)
	assert.Equal(t, "STRING", testString)
}
