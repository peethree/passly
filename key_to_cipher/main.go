package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func keyToCipher(key string) (cipher.Block, error) {
	// convert to binary
	binKey := []byte(key)
	return aes.NewCipher(binKey)
}
