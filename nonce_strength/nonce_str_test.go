package main

import (
	"fmt"
	"testing"
)

func TestNonceStrength(t *testing.T) {
	type testCase struct {
		nonceLength int
		expected    int
	}

	tests := []testCase{
		{1, 256},           // 1 byte (8 bits) => 2^8 = 256
		{2, 65536},         // 2 bytes (16 bits) => 2^16 = 65536
		{3, 16777216},      // 3 bytes (24 bits) => 2^24 = 16777216
		{4, 4294967296},    // 4 bytes (32 bits) => 2^32 = 4294967296
		{5, 1099511627776}, // 5 bytes (40 bits) => 2^40 = 1099511627776
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{6, 281474976710656},   // 6 bytes (48 bits) => 2^48
			{7, 72057594037927936}, // 7 bytes (56 bits) => 2^56
		}...)
	}

	for _, test := range tests {
		nonce, _ := generateIV(test.nonceLength)
		actualStrength := nonceStrength(nonce)

		if actualStrength != test.expected {
			t.Errorf(`---------------------------------
Inputs:      nonce length: %d
Expecting:   nonce strength: %d
Actual:      nonce strength: %d
Fail`, test.nonceLength, test.expected, actualStrength)
		} else {
			fmt.Printf(`---------------------------------
Inputs:      nonce length: %d
Expecting:   nonce strength: %d
Actual:      nonce strength: %d
Pass
`, test.nonceLength, test.expected, actualStrength)
		}
	}
}

var withSubmit = true
