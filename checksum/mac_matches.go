package main

import (
	"crypto/sha256"
	"fmt"
)

func macMatches(message, key, checksum string) bool {
	//concatenate key to message (MAC)
	message += key
	h := sha256.New()
	h.Write([]byte(message))
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}
