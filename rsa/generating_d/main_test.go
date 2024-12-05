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
		{"I hid the cash under the sink", 512, "1594991729..."},
		{"Don't you think they will look there??", 512, "1857109338..."},
		{"They'll look at everything but the kitchen sink", 1024, "1585633466..."},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Where should we stash it next?", 512, "1332321687..."},
			{"No one would think to check under the rug", 1024, "3130951121..."},
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

		d := getD(e, tot)
		firstD := firstNDigits(*d, 10)

		if firstD != test.expected {
			t.Errorf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated tot: %s
generated e: %s
Expecting d: %s
Actual d: %s
Fail`, test.msg, test.keySize, firstP, firstQ, firstN, firstTOT, firstE, test.expected, firstD)
		} else {
			fmt.Printf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated tot: %s
generated e: %s
Expecting d: %s
Actual d: %s
Pass
`, test.msg, test.keySize, firstP, firstQ, firstN, firstTOT, firstE, test.expected, firstD)
		}
	}
}

var withSubmit = true
