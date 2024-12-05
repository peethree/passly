package main

import (
	"fmt"
	"testing"
)

func TestHashFunction(t *testing.T) {
	type testCase struct {
		input    string
		expected [4]byte
	}

	tests := []testCase{
		{"Example message", [4]byte{0x24, 0xC0, 0x40, 0xC4}},
		{"This is a slightly longer example to hash", [4]byte{0x28, 0x88, 0xC8, 0x48}},
		{"This is a much longer example of some text to hash, maybe it's the opening paragraph of a blog post", [4]byte{0xA8, 0x04, 0x44, 0xE8}},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"A very secret password", [4]byte{0x64, 0xA0, 0x64, 0xAC}},
			{"Another very secret password", [4]byte{0x40, 0xCC, 0xEC, 0x2C}},
		}...)
	}

	for _, test := range tests {
		actual := hash([]byte(test.input))

		if actual != test.expected {
			t.Errorf(`---------------------------------
Hashing: '%s'
Expecting:   %X
Actual:      %X
Fail`, test.input, test.expected, actual)
		} else {
			fmt.Printf(`---------------------------------
Hashing: '%s'
Expecting:   %X
Actual:      %X
Pass
`, test.input, test.expected, actual)
		}
	}
}

var withSubmit = true
