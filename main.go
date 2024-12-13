package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// convert pw to byte
	bytePw := []byte(password)

	cost := 13

	// generate bcrypt hash
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err := bcrypt.GenerateFromPassword(bytePw, cost)
	if err != nil {
		return "", err
	}

	// convert hash to str
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	// convert to byte
	bytePw := []byte(password)
	byteHash := []byte(hash)

	// compare hash with pw
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err := bcrypt.CompareHashAndPassword(byteHash, bytePw)
	if err != nil {
		return false
	}
	return true
}
