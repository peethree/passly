package main

import (
	"math"
	"math/rand"
)

// nonce = arbitrary numbr used only once in cryptographic op
func nonceStrength(nonce []byte) int {
	// number of bytes in nonce
	length := len(nonce)

	// * 8 to get bits
	bits := float64(length * 8)

	// entropy is 2^(number of bits)
	return int(math.Pow(2, bits))
}

func generateIV(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}
