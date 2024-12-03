package main

import (
	"math/rand"
)

func generateIV(length int) ([]byte, error) {
	uninitIV := make([]byte, length)
	// func Read(p []byte) (n int, err error)
	_, err := rand.Read(uninitIV)
	if err != nil {
		return []byte{}, err
	}
	// iv not has pseudo random bytes read into it.
	iv := uninitIV
	return iv, nil
}
