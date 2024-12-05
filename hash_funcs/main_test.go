package main

import (
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	type testCase struct {
		passwords []string
		expected  string
	}

	tests := []testCase{
		{[]string{"password1", "password2", "password3"}, "2ccb27b6da"},
		{[]string{"abercromni3", "f1tch", "123456", "abcdefg1234"}, "a03ea2f828"},
		{[]string{"IHeartNanciedrake", "m7B1rthd@y"}, "7bf580ff47"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{[]string{"morepassw0rds", "evenmorepassw0rds"}, "2f595e539a"},
			{[]string{"s3cur3passw0rd"}, "db25c1918f"},
		}...)
	}

	for _, test := range tests {
		h := newHasher()

		for _, password := range test.passwords {
			h.Write(password)
		}

		actual := h.GetHex()

		if actual[:10] != test.expected {
			t.Errorf(`---------------------------------
Hashing vault with passwords: %v
Expecting:   hash starts with: %s
Actual:      hash starts with: %s
Fail`, test.passwords, test.expected, actual[:10])
		} else {
			fmt.Printf(`---------------------------------
Hashing vault with passwords: %v
Expecting:   hash starts with: %s
Actual:      hash starts with: %s
Pass
`, test.passwords, test.expected, actual[:10])
		}
	}
}

var withSubmit = true
