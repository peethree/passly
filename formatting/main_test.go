package main

import (
	"fmt"
	"testing"
)

func TestGetHexString(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}

	tests := []testCase{
		{[]byte("Hello"), "48:65:6c:6c:6f"}, // Hex for "Hello"
		{[]byte("World"), "57:6f:72:6c:64"}, // Hex for "World"
	}

	// Additional test cases for withSubmit
	if withSubmit {
		tests = append(tests, []testCase{
			{[]byte("GoLang"), "47:6f:4c:61:6e:67"}, // Hex for "GoLang"
			{[]byte("Passly"), "50:61:73:73:6c:79"}, // Hex for "Passly"
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := getHexString(test.input)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Fail`, test.input, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Pass`, test.input, test.expected, result)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestGetBinaryString(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}

	tests := []testCase{
		{[]byte("Hello"), "01001000:01100101:01101100:01101100:01101111"}, // Binary for "Hello"
		{[]byte("World"), "01010111:01101111:01110010:01101100:01100100"}, // Binary for "World"
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[]byte("GoLang"), "01000111:01101111:01001100:01100001:01101110:01100111"}, // Binary for "GoLang"
			{[]byte("Passly"), "01010000:01100001:01110011:01110011:01101100:01111001"}, // Binary for "Passly"
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		result := getBinaryString(test.input)
		if result != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Fail`, test.input, test.expected, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      %v
Expecting:   %v
Actual:      %v
Pass`, test.input, test.expected, result)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
