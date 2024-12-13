package main

import (
	"fmt"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	type testCase struct {
		password1 string
		password2 string
		saltLen   int
		expect    bool
	}

	tests := []testCase{
		{"samepass", "samepass", 16, true},
		{"passone", "passtwo", 24, false},
		{"correct horse battery staple", "correct horse battery staple", 32, true},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"bigtimepass", "notthesame", 16, false},
			{"kaladin", "kaladin", 24, true},
			{"stormlight archive", "stormlight archive", 32, true},
		}...)
	}

	for _, test := range tests {
		salt, err := generateSalt(test.saltLen)
		if err != nil {
			t.Errorf("Error generating salt: %v", err)
			continue
		}

		hashed1 := hashPassword([]byte(test.password1), salt)
		hashed2 := hashPassword([]byte(test.password2), salt)

		match := string(hashed1) == string(hashed2)
		if match != test.expect {
			t.Errorf(`---------------------------------
Password 1:  %s
Password 2:  %s
Salt length: %d
Expecting:   %v
Actual:      %v
Fail`, test.password1, test.password2, test.saltLen, test.expect, match)
		} else {
			fmt.Printf(`---------------------------------
Password 1:  %s
Password 2:  %s
Salt length: %d
Expecting:   %v
Actual:      %v
Pass
`, test.password1, test.password2, test.saltLen, test.expect, match)
		}
	}
}

var withSubmit = true
