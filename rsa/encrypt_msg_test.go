package main

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type testCase struct {
		message  string
		expected string
	}

	tests := []testCase{
		{"Hey Darling, don't come over tonight, I'm out with my people", "Hey Darling, don't come over tonight, I'm out with my people"},
		{"Yes, ten million in cash. No, every penny better be accounted for", "Yes, ten million in cash. No, every penny better be accounted for"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Do you know what would happen if I suddenly decided to stop going into work? A business big enough that it could be listed on the NASDAQ goes belly up. Disappears! It ceases to exist without me. No, you clearly don't know who you're talking to, so let me clue you in. I am not in danger, Skyler. I am the danger. A guy opens his door and gets shot and you think that of me? No. I am the one who knocks!", ""},
		}...)
	}

	pub, priv, err := genKeys()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	for _, test := range tests {
		ciphertext, err := encrypt(pub, []byte(test.message))
		if err != nil {
			if test.expected == "" {
				fmt.Printf("Expected failure for long message: %v\n", err)
				continue
			}
			t.Errorf("Encryption failed for message: %v, error: %v", test.message, err)
			continue
		}

		plaintext, err := decrypt(priv, ciphertext)
		if err != nil {
			t.Errorf("Decryption failed for ciphertext: %v, error: %v", ciphertext, err)
			continue
		}

		if string(plaintext) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      message: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Fail`, test.message, test.expected, string(plaintext))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      message: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Pass
`, test.message, test.expected, string(plaintext))
		}
	}
}

var withSubmit = true
