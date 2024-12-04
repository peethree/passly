package main

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	type testCase struct {
		key       []byte
		plaintext []byte
		nonce     []byte
		expected  string
	}

	tests := []testCase{
		{[]byte("d00c5215-60f6-4ac4-9648-532b5dad"), []byte("I wonder what he's thinking about me??"), generateNonce(12), "I wonder what he's thinking about me??"},
		{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("I knew it, Becky has been cheating this whole time!"), generateNonce(12), "I knew it, Becky has been cheating this whole time!"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("AES in GCM mode is strong!"), generateNonce(12), "AES in GCM mode is strong!"},
			{[]byte("db50ecaaa-23ed-43eb-9f8b-6fc5931"), []byte("Testing AES-GCM encryption."), generateNonce(12), "Testing AES-GCM encryption."},
		}...)
	}

	for _, test := range tests {
		ciphertext, err := encrypt(test.key, test.plaintext, test.nonce)
		if err != nil {
			t.Errorf("Encryption failed for key: %v, plaintext: %v, error: %v", string(test.key), string(test.plaintext), err)
			continue
		}

		decryptedText, err := decrypt(test.key, ciphertext, test.nonce)
		if err != nil {
			t.Errorf("Decryption failed for key: %v, ciphertext: %v, error: %v", string(test.key), ciphertext, err)
			continue
		}

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
