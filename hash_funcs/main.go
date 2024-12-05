package main

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	// create new hash.Hash
	new := sha256.New()

	hasher := hasher{
		hash: new,
	}
	// return pointer to new hasher
	return &hasher
}

func (h *hasher) Write(s string) (n int, err error) {
	// cast the string to a []byte
	byteString := []byte(s)
	// return the number of bytes written from h + err
	return h.hash.Write(byteString)
}

func (h *hasher) GetHex() string {
	// pass in nil becasue there's no need to append the hash to existing byteslice
	hashValue := h.hash.Sum(nil)
	// encode hash to lowercase hex string
	return hex.EncodeToString(hashValue)
}
