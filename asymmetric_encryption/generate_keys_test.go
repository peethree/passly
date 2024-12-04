package main

import (
	"fmt"
	"testing"
)

func TestGenKeys(t *testing.T) {
	type testCase struct {
		expected bool
	}

	tests := []testCase{
		{true},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{true},
		}...)
	}

	for _, test := range tests {
		pubKey, privKey, err := genKeys()
		if err != nil {
			t.Errorf("Failed to generate key pair, error: %v", err)
			continue
		}

		arePaired := keysArePaired(pubKey, privKey)
		if arePaired != test.expected {
			t.Errorf(`---------------------------------
Inputs:      public key: %v, private key: %v
Expecting:   paired: %v
Actual:      paired: %v
Fail`, pubKey, privKey, test.expected, arePaired)
		} else {
			fmt.Printf(`---------------------------------
Inputs:      public key: %v, private key: %v
Expecting:   paired: %v
Actual:      paired: %v
Pass
`, pubKey, privKey, test.expected, arePaired)
		}
	}
}

var withSubmit = true
