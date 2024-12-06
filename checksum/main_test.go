package main

import (
	"fmt"
	"testing"
)

func TestChecksumMatches(t *testing.T) {
	type testCase struct {
		message  string
		checksum string
		expected bool
	}

	tests := []testCase{
		{"pa$$w0rd", "4b358ed84b7940619235a22328c584c7bc4508d4524e75231d6f450521d16a17", true},
		{"buil4WithB1ologee", "1c489a153271aaf3b234aa154b1a2eef5248eb9ab402e4d3c8b7bc3d81fed1a8", false},
		{"br3ak1ngB@d1sB3st", "5d178e1c6fd5d76415e1632f84e5192fb50ef244d42a02148fedbf991d914546", false},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"invalidTest", "1234567890abcdef", false},
			{"b3ttterC@llS@ulI$B3tter", "8d42f2dc81476123974619969a42b27b8d8a4fa507be99c9623f614ad2d859f7", true},
		}...)
	}

	for _, test := range tests {
		result := checksumMatches(test.message, test.checksum)

		if result != test.expected {
			t.Errorf(`---------------------------------
Checking:   %s
Expecting:  %v
Actual:     %v
Fail`, test.message, test.expected, result)
		} else {
			fmt.Printf(`---------------------------------
Checking:   %s
Expecting:  %v
Actual:     %v
Pass
`, test.message, test.expected, result)
		}
	}
}

var withSubmit = true
