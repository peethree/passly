package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func checksumMatches(message string, checksum string) bool {
	// convrt msg to slice of bytes
	hash := sha256.New()
	hash.Write([]byte(message))

	// check if lowercase hexadecimal encoding of the hash matches checksum input
	if checksum == hex.EncodeToString(hash.Sum(nil)) {
		return true
	}

	return false
}
