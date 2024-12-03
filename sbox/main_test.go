package main

import (
	"fmt"
	"testing"
)

func TestSBox(t *testing.T) {
	type testCase struct {
		input    byte
		expected byte
	}

	tests := []testCase{
		{0b0000, 0b00},
		{0b0001, 0b10},
		{0b0110, 0b11},
		{0b1111, 0b00},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{0b1001, 0b11},
			{0b0010, 0b01},
		}...)
	}

	for _, test := range tests {
		result, err := sBox(test.input)
		if err != nil {
			t.Errorf("Failed for input: %04b, error: %v", test.input, err)
			continue
		}

		if result != test.expected {
			t.Errorf(`---------------------------------
Inputs:      input: %04b
Expecting:   output: %02b
Actual:      output: %02b
Fail`, test.input, test.expected, result)
		} else {
			fmt.Printf(`---------------------------------
Inputs:      input: %04b
Expecting:   output: %02b
Actual:      output: %02b
Pass
`, test.input, test.expected, result)
		}
	}
}

var withSubmit = true
