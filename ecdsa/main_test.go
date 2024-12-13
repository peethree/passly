package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestECDSAMessage(t *testing.T) {
	type testCase struct {
		message  string
		expected string
	}

	tests := []testCase{
		{"userid:2f9c584e-5d25-4516-a0ed-ddfa6e152006", "valid"},
		{"userid:0e803af6-292f-4432-a285-84a7591e25a8", "valid"},
		{"userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f", "valid"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"userid:f77e36d6-0edc-44ef-964e-af4a5b1ebd5f", "valid"},
		}...)
	}

	for _, test := range tests {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error generating key: %v", err)
			continue
		}

		token, err := createECDSAMessage(test.message, privateKey)
		if err != nil {
			t.Errorf("Error creating ECDSA message: %v", err)
			continue
		}

		err = verifyECDSAMessage(token, &privateKey.PublicKey)

		if err != nil {
			t.Errorf(`---------------------------------
Message:      %s
Expecting:    valid
Actual:       invalid
Fail`, test.message)
		} else {
			fmt.Printf(`---------------------------------
Message:      %s
Expecting:    valid
Actual:       valid
Pass
`, test.message)
		}
	}
}

var withSubmit = true
