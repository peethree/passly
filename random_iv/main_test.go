package main

import (
	"fmt"
	"testing"
)

func TestGenerateIV(t *testing.T) {
	type testCase struct {
		length   int
		expected int
	}

	tests := []testCase{
		{8, 8},
		{10, 10},
		{16, 16},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{12, 12},
			{14, 14},
		}...)
	}

	for _, test := range tests {
		iv, err := generateIV(test.length)
		if err != nil {
			t.Errorf("Failed to generate IV for length %d: %v", test.length, err)
			continue
		}

		if len(iv) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      length: %d
Expecting:   IV length: %d
Actual:      IV length: %d
Fail`, test.length, test.expected, len(iv))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      length: %d
Expecting:   IV length: %d
Actual:      IV length: %d
Pass
`, test.length, test.expected, len(iv))
		}
	}
}

var withSubmit = true
