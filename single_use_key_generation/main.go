package main

// this function should be using crypto/rand instead
import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	// make slice to put read data into
	keySlice := make([]byte, length)

	// read data into slice
	_, err := rand.Read(keySlice)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", keySlice), nil
}
