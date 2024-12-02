package main

import (
	"fmt"
	"testing"
)

func TestDeriveRoundKey(t *testing.T) {
	type testCase struct {
		masterKey   [4]byte
		roundNumber int
		expected    [4]byte
	}

	tests := []testCase{
		{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 1, [4]byte{0xAB, 0xFE, 0x10, 0xBD}},
		{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 2, [4]byte{0xE9, 0xCF, 0x11, 0xFE}},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 5, [4]byte{0xAF, 0xFA, 0x14, 0xB9}},
			{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 7, [4]byte{0xEC, 0xCA, 0x14, 0xFB}},
		}...)
	}

	for _, test := range tests {
		result := deriveRoundKey(test.masterKey, test.roundNumber)
		if result != test.expected {
			t.Errorf(`---------------------------------
Inputs:      masterKey: %X, roundNumber: %d
Expecting:   roundKey: %X
Actual:      roundKey: %X
Fail`, test.masterKey, test.roundNumber, test.expected, result)
		} else {
			fmt.Printf(`---------------------------------
Inputs:      masterKey: %X, roundNumber: %d
Expecting:   roundKey: %X
Actual:      roundKey: %X
Pass
`, test.masterKey, test.roundNumber, test.expected, result)
		}
	}
}

var withSubmit = true
