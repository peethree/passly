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
		{"I hid the cash under the sink", 512, "7346023901..."},
		{"Don't you think they will look there??", 512, "1098234792..."},
		{"They'll look at everything but the kitchen sink", 1024, "5628412345..."},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"Where should we stash it next?", 512, "4857392091..."},
			{"No one would think to check under the rug", 1024, "6923740198..."},
		}...)
	}

	for _, test := range tests {
		msgBytes := []byte(test.msg)

		p, q := generatePrivateNums(test.keySize)
		n := getN(p, q)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstN := firstNDigits(*n, 10)

		phi := getPhi(p, q)
		e := getE(phi)

		firstPhi := firstNDigits(*phi, 10)
		firstE := firstNDigits(*e, 10)

		plaintext := big.NewInt(0)
		plaintext.SetBytes(msgBytes)
		ciphertext := encrypt(plaintext, e, n)

		d := getD(e, phi)
		decrypted := decrypt(ciphertext, d, n)

		firstDecrypted := string(decrypted.Bytes())

		if firstDecrypted != test.msg {
			t.Errorf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated phi: %s
generated e: %s
Expecting: %s
Actual: %s
Fail`, test.msg, test.keySize, firstP, firstQ, firstN, firstPhi, firstE, test.msg, firstDecrypted)
		} else {
			fmt.Printf(`---------------------------------
Encrypting: %s with key size %v
Generated primes:
p: %s
q: %s
n: %s
generated phi: %s
generated e: %s
Expecting: %s
Actual: %s
Pass
`, test.msg, test.keySize, firstP, firstQ, firstN, firstPhi, firstE, test.msg, firstDecrypted)
		}
	}
}

var withSubmit = true
