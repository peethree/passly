package main

import (
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	type testCase struct {
		plaintext string
		key       int
		expected  string
	}

	tests := []testCase{
		{"abcdefghi", 1, "bcdefghij"},
		{"hello", 5, "mjqqt"},
		{"correcthorsebatterystaple", 16, "sehhusjxehiurqjjuhoijqfbu"},
		{"onetwothreefourfivesixseveneightnineten", 25, "nmdsvnsgqddentqehudrhwrdudmdhfgsmhmdsdm"},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		encrypted := encrypt(test.plaintext, test.key)
		decrypted := decrypt(encrypted, test.key)
		if decrypted != test.plaintext {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      plaintext: %v, key: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Fail`, test.plaintext, test.key, test.plaintext, decrypted)
		} else if encrypted != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      plaintext: %v, key: %v
Expecting:   encrypted: %v
Actual:      encrypted: %v
Fail`, test.plaintext, test.key, test.expected, encrypted)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      plaintext: %v, key: %v
Expecting:   encrypted: %v
Actual:      encrypted: %v
Pass`, test.plaintext, test.key, test.expected, encrypted)
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
