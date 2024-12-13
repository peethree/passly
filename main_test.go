package main

import (
	"fmt"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	type testCase struct {
		password1 string
		password2 string
		expected  bool
	}

	tests := []testCase{
		{"thisIsAPassword", "thisIsAPassword", true},
		{"thisIsAPassword", "thisIsAnotherPassword", false},
		{"corr3ct h0rse", "corr3ct h0rse", true},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"thisIsAPassword", "thisIsAPassword", true},
			{"thisIsAPassword", "thisIsAnotherPassword", false},
			{"corr3ct h0rse", "corr3ct h0rse", true},
		}...)
	}

	for _, test := range tests {
		hashed, err := hashPassword(test.password1)
		if err != nil {
			t.Errorf("Error hashing password: %v", err)
			continue
		}

		match := checkPasswordHash(test.password2, hashed)

		if match != test.expected {
			t.Errorf(`---------------------------------
Password 1:  %s
Password 2:  %s
Expecting:   %v
Actual:      %v
Fail`, test.password1, test.password2, test.expected, match)
		} else {
			fmt.Printf(`---------------------------------
Password 1:  %s
Password 2:  %s
Expecting:   %v
Actual:      %v
Pass
`, test.password1, test.password2, test.expected, match)
		}
	}
}

var withSubmit = true
