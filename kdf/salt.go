package main

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt(length int) ([]byte, error) {
	// generate random salt
	salt := make([]byte, length)

	_, err := rand.Read(salt)
	if err != nil {
		return []byte{}, err
	}

	return salt, nil
}

func hashPassword(password, salt []byte) []byte {
	// append salt to pw
	pwSalt := append(password, salt...)

	// hash
	hash := sha256.New()
	hash.Write(pwSalt)
	return hash.Sum(nil)
}
