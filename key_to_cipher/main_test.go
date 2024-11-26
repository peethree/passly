package main

import (
	"fmt"
	"testing"
)

func TestKeyToCipher(t *testing.T) {
	type testCase struct {
		key        string
		shouldFail bool
	}

	tests := []testCase{
		{"thisIsMySecretKeyIHopeNoOneFinds", false}, // Valid key
		{"short", true}, // Too short key
		{"an extremely long key that exceeds the block size", true}, // Too long key
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"thisIsA32ByteKeyForAES256Testing!", true}, // Valid 32-byte key for AES-256
			{"valid16ByteKeyHere", true},                // Valid 16-byte key for AES-128
			{"invalid-key", true},                       // Invalid key, not the correct length
			{"ThisIsA24ByteKeyForAES192Testing", false}, // Valid key for AES-192
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		_, err := keyToCipher(test.key)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      key: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail`, test.key, test.shouldFail, err != nil)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      key: %v
Expecting:   Error: %v
Actual:      Error: %v
Pass`, test.key, test.shouldFail, err != nil)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
