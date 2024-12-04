package main

import (
	"fmt"
	"testing"
)

func TestGetTotAndGetE(t *testing.T) {
	type testCase struct {
		keySize  int
		expected int
	}

	tests := []testCase{
		{512, 309},  // Expected number of digits for tot with 512-bit primes
		{1024, 617}, // Expected number of digits for tot with 1024-bit primes
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{2048, 1233}, // Expected number of digits for tot with 2048-bit primes
		}...)
	}

	for _, test := range tests {
		p, q := generatePrivateNums(test.keySize)
		tot := getTot(p, q)
		e := getE(tot)

		firstP := firstNDigits(*p, 10)
		firstQ := firstNDigits(*q, 10)
		firstTot := firstNDigits(*tot, 10)
		firstE := firstNDigits(*e, 10)

		if len(tot.String()) != test.expected {
			t.Errorf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
ϕ(n): %s
Expecting:   tot digits: %d
Actual:      tot digits: %d
Fail`, test.keySize, firstP, firstQ, firstTot, test.expected, len(tot.String()))
		} else {
			fmt.Printf(`---------------------------------
Inputs:      key size: %d
Generated primes:
p: %s
q: %s
ϕ(n): %s
Expecting:   tot digits: %d
Actual:      tot digits: %d
Pass
`, test.keySize, firstP, firstQ, firstTot, test.expected, len(tot.String()))
		}

		fmt.Printf(`Generated e: %s
========
`, firstE)
	}
}
