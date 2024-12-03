package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	type testCase struct {
		key       []byte
		plaintext []byte
		expected  string
	}

	tests := []testCase{
		{[]byte("12344321"), []byte("Today I met my crush, what a hunk"), "Today I met my crush, what a hunk"},
		{[]byte("p@$$w0rd"), []byte("I hope my boyfriend never finds out about this"), "I hope my boyfriend never finds out about this"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[]byte("secretky"), []byte("The best secret ever!"), "The best secret ever!"},
			{[]byte("longpass"), []byte("Just testing DES encryption with padding"), "Just testing DES encryption with padding"},
		}...)
	}

	for _, test := range tests {
		ciphertext, err := encrypt(test.key, test.plaintext)
		if err != nil {
			t.Errorf("Encryption failed for key: %v, plaintext: %v, error: %v", string(test.key), string(test.plaintext), err)
			continue
		}

		decryptedText, err := decrypt(test.key, ciphertext)
		if err != nil {
			t.Errorf("Decryption failed for key: %v, ciphertext: %v, error: %v", string(test.key), ciphertext, err)
			continue
		}
		decryptedText = bytes.Trim(decryptedText, "\x00")

		if string(decryptedText) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      key: %v, plaintext: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Fail`, string(test.key), string(test.plaintext), test.expected, string(decryptedText))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      key: %v, plaintext: %v
Expecting:   decrypted: %v
Actual:      decrypted: %v
Pass
`, string(test.key), string(test.plaintext), test.expected, string(decryptedText))
		}
	}
}

var withSubmit = true
