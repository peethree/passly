package main

import (
	"math/bits"
)

func hash(input []byte) [4]byte {
	// rotate the bits left 3 bits

	for i := 0; i < len(input); i++ {
		input[i] = bits.RotateLeft8(input[i], 3)
	}

	shifted := make([]byte, len(input))
	// shift the bits left 2 bits
	for i := 0; i < len(input); i++ {
		shifted[i] = input[i] << 2
	}

	// new emptyy final array
	final := [4]byte{}

	// index i%4 to avoid going out of range
	for i := 0; i < len(input); i++ {
		final[i%4] = final[i%4] ^ shifted[i]
	}

	return final
}
