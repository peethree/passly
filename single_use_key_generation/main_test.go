package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenerateRandomKey(t *testing.T) {
	rand.Seed(0)

	type testCase struct {
		length     int
		shouldFail bool
		expected   string
	}

	tests := []testCase{
		{16, false, "0194fdc2fa2ffcc041d3ff12045b73c8"},                                 // Expected output for 16 bytes
		{32, false, "6e4ff95ff662a5eee82abdf44a2d0b75fb180daf48a79ee0b10d394651850fd4"}, // Expected output for 32 bytes
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{8, false, "a178892ee285ece1"}, // Expected output for 8 bytes
			{64, false, "511455780875d64ee2d3d0d0de6bf8f9b44ce85ff044c6b1f83b8e883bbf857aab99c5b252c7429c32f3a8aeb79ef856f659c18f0dcecc77c75e7a81bfde275f"}, // Expected output for 64 bytes
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		key, err := generateRandomKey(test.length)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      length: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail`, test.length, test.shouldFail, err != nil)
		} else {
			if !test.shouldFail && key != test.expected {
				failCount++
				t.Errorf(`---------------------------------
Inputs:      length: %v
Expecting:   Key: %v
Actual:      Key: %v
Fail`, test.length, test.expected, key)
			} else if test.shouldFail {
				passCount++
				fmt.Printf(`---------------------------------
Inputs:      length: %v
Expecting:   Error: %v
Actual:      Error: %v
Pass`, test.length, test.shouldFail, err != nil)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Inputs:      length: %v
Expecting:   Key: %v
Actual:      Key: %v
Pass`, test.length, test.expected, key)
			}
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
