package lazygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	cases := []struct {
		testName  string
		str       string
		expectStr string
	}{
		{
			testName: "to lower",
			str: func() string {
				testString := "StRing"
				ToLower(&testString)
				return testString
			}(),
			expectStr: "string",
		},
		{
			testName: "to upper",
			str: func() string {
				testString := "StRing"
				ToUpper(&testString)
				return testString
			}(),
			expectStr: "STRING",
		},
		{
			testName: "right pad",
			str: func() string {
				testString := "STRING"
				RightPad(&testString, "S", 10)
				return testString
			}(),
			expectStr: "STRINGSSSS",
		},
		{
			testName: "replace",
			str: func() string {
				testString := "ABCABCABC"
				Replace(&testString, "A", "Z")
				return testString
			}(),
			expectStr: "ZBCZBCZBC",
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			assert.Equal(t, c.expectStr, c.str)
		})
	}
}
