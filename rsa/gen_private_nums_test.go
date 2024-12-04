package main

import (
	"fmt"
	"testing"
)

func TestGeneratePrivateNumsAndGetN(t *testing.T) {
	type testCase struct {
		keySize  int
		expected int // Number of digits in n
	}

	tests := []testCase{
		{512, 309},  // Expected number of digits for 512-bit primes
		{1024, 617}, // Expected number of digits for 1024-bit primes
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{2048, 1233}, // Expected number of digits for 2048-bit primes
		}...)
	}

	for _, test := range tests {
		p, q := generatePrivateNums(test.keySize)
		n := getN(p, q)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstN := firstNDigits(*n, 10)

		if len(n.String()) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
n: %s
Expecting:   n digits: %d
Actual:      n digits: %d
Fail`, test.keySize, firstP, firstQ, firstN, test.expected, len(n.String()))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
n: %s
Expecting:   n digits: %d
Actual:      n digits: %d
Pass
`, test.keySize, firstP, firstQ, firstN, test.expected, len(n.String()))
		}
	}
}
