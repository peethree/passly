package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func hmac(message, key string) string {
	splitIndex := len(key) / 2
	key1 := key[:splitIndex]
	key2 := key[splitIndex:]

	// sha256(keyFirstHalf + sha256(keySecondHalf + message)) as a
	key2PlusMsg := key2 + message
	s := sha256.New()
	s.Write([]byte(key2PlusMsg))

	f := sha256.New()
	// concatenate the first half of the key and the hash of 2nd key half + msg
	key1Plus := append([]byte(key1), s.Sum(nil)...)
	f.Write(key1Plus)

	// string in lowercase hex
	return hex.EncodeToString(f.Sum(nil))
}
