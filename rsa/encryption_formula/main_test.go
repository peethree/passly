package main

import (
	"fmt"
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		msg      string
		keySize  int
		expected string
	}

	tests := []testCase{
		{"I hid the cash under the sink", 512, "1093960760..."},
		{"Don't you think they will look there??", 512, "5612332892..."},
		{"They'll look at everything but the kitchen sink", 1024, "1551668194..."},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Where should we stash it next?", 512, "1243871900..."},
			{"No one would think to check under the rug", 1024, "1210681187..."},
		}...)
	}

	for _, test := range tests {
		msgBytes := []byte(test.msg)

		p, q := generatePrivateNums(test.keySize)
		n := getN(p, q)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstN := firstNDigits(*n, 10)

		tot := gettot(p, q)
		e := getE(tot)

		firstTOT := firstNDigits(*tot, 10)
		firstE := firstNDigits(*e, 10)

		plaintext := big.NewInt(0)
		plaintext.SetBytes(msgBytes)
		ciphertext := encrypt(plaintext, e, n)

		firstCipher := firstNDigits(*ciphertext, 10)

		if firstCipher != test.expected {
			t.Errorf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated total: %s
generated e: %s
Expecting: %s
Actual: %s
Fail`, test.msg, test.keySize, firstP, firstQ, firstN, firstE, firstTOT, test.expected, firstCipher)
		} else {
			fmt.Printf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated total: %s
generated e: %s
Expecting: %s
Actual: %s
Pass
`, test.msg, test.keySize, firstP, firstQ, firstN, firstE, firstTOT, test.expected, firstCipher)
		}

	}
}

var withSubmit = true