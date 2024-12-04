package main

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand" // crypto/rand is better
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	// create new cipher block
	newCipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// create new GCM with cipher block
	// func NewGCM(cipher Block) (AEAD, error)
	gcm, err := cipher.NewGCM(newCipherBlock)
	if err != nil {
		return []byte{}, err
	}

	// decrypt the cyphertext using AEAD interface
	// Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
	decryptedCypherText, err := gcm.Open(ciphertext[:0], nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return decryptedCypherText, nil
}

func encrypt(key, plaintext, nonce []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func generateNonce(length int) []byte {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}
