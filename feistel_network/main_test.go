package main

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"testing"
)

func TestFeistel(t *testing.T) {
	type testCase struct {
		msg      []byte
		key      []byte
		rounds   int
		expected string
	}

	tests := []testCase{
		{[]byte("General Kenobi!!!!"), []byte("thesecret"), 8, "General Kenobi!!!!"},
		{[]byte("Hello there!"), []byte("@n@kiN"), 16, "Hello there!"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[]byte("Goodbye!"), []byte("roundkey"), 8, "Goodbye!"},
		}...)
	}

	for _, test := range tests {
		roundKeys := generateRoundKeys(test.key, test.rounds)
		encrypted := feistel(test.msg, roundKeys)
		decrypted := feistel(encrypted, reverse(roundKeys))

		if string(decrypted) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      msg: %v, key: %v, rounds: %d
Expecting:   decrypted: %s
Actual:      decrypted: %s
Fail`, test.msg, test.key, test.rounds, test.expected, string(decrypted))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      msg: %v, key: %v, rounds: %d
Expecting:   decrypted: %s
Actual:      decrypted: %s
Pass
`, test.msg, test.key, test.rounds, test.expected, string(decrypted))
		}
	}
}

func generateRoundKeys(key []byte, rounds int) [][]byte {
	roundKeys := [][]byte{}
	for i := 0; i < rounds; i++ {
		ui := binary.BigEndian.Uint32(key)
		rotated := bits.RotateLeft32(uint32(ui), i)
		finalRound := make([]byte, len(key))
		binary.LittleEndian.PutUint32(finalRound, uint32(rotated))
		roundKeys = append(roundKeys, finalRound)
	}
	return roundKeys
}

var withSubmit = true
