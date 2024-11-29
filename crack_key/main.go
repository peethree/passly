package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	maxIndex := int(math.Pow(2, 24))
	// func intToBytes(num int) []byte {
	for i := 0; i < maxIndex; i++ {
		key := intToBytes(i)
		// func crypt(dat, key []byte) []byte {
		decryptedWithKey := crypt(encrypted, key)
		if string(decryptedWithKey) == decrypted {
			return key, nil
		}
	}
	return []byte{}, nil
}

// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}
