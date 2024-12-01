package main

import (
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	type testCase struct {
		plaintext string
		key       string
	}

	tests := []testCase{
		{"Shazam", "Sk7p13"},             // Key length matches plaintext length
		{"I'm lovin it", "mysecurepass"}, // Key length matches plaintext length
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Kaladin", "Radiant"},           // Updated key length matches plaintext
			{"Another test", "shorttestkey"}, // Updated key length matches plaintext
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		fmt.Printf("Encrypting plaintext: '%s' with key: '%s'\n", test.plaintext, test.key)

		ciphertext, err := encrypt([]byte(test.plaintext), []byte(test.key))
		if err != nil {
			t.Errorf("Error during encryption: %v", err)
			failCount++
			continue
		}

		fmt.Printf("Encrypted ciphertext bytes: %v\n", ciphertext)

		decrypted, err := decrypt(ciphertext, []byte(test.key))
		if err != nil {
			t.Errorf("Error during decryption: %v", err)
			failCount++
			continue
		}

		if string(decrypted) != test.plaintext {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      plaintext: %s, key: %s
Expecting:   decrypted: %s
Actual:      decrypted: %s
Fail`, test.plaintext, test.key, test.plaintext, string(decrypted))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      plaintext: %s, key: %s
Expecting:   decrypted: %s
Actual:      decrypted: %s
Pass`, test.plaintext, test.key, test.plaintext, string(decrypted))
		}
	}

	fmt.Printf("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
