package main

import (
	"fmt"
	"testing"
)

func TestMacMatches(t *testing.T) {
	type testCase struct {
		message  string
		key      string
		expected bool
		checksum string
	}

	tests := []testCase{
		{"pa$$w0rd", "abdf6b86cb", true, "7b1dede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71"},
		{"buil4WithB1ologee", "6ddf6b86cb", false, "1cddede90198e94c7432358e9bd937b91015cc45c84af5dfbedcd1a3c764ff71"},
		{"br3ak1ngB@d1sB3st", "7adf6b86cb", false, "2c678e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546"},
		{"b3ttterC@llS@ulI$B3tter", "12df6b86cb", true, "eb4d4516bd4141322c3ab160bc2b3010eaf7bd19f821d0c48f480791d32af359"},
		{"wrongMessage", "wrongKey", false, "incorrectchecksum"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"someMessage", "extraKey", false, "0123456789abcdef0123456789abcdef"},
			{"anotherMessage", "superKey", false, "fedcba9876543210fedcba9876543210"},
		}...)
	}

	for _, test := range tests {
		h := macMatches(test.message, test.key, test.checksum)

		if h != test.expected {
			t.Errorf(`---------------------------------
Checking message: '%s' with key: '%s'
Expecting:   %t
Actual:      %t
Fail`, test.message, test.key, test.expected, h)
		} else {
			fmt.Printf(`---------------------------------
Checking message: '%s' with key: '%s'
Expecting:   %t
Actual:      %t
Pass
`, test.message, test.key, test.expected, h)
		}
	}
}
