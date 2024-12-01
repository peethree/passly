package main

import (
	"crypto/aes"
	"crypto/des"
	"errors"
	"fmt"
)

// operates on fixed-length groups of data (symmetric encryption)
func getBlockSize(keyLen, cipherType int) (int, error) {
	key := make([]byte, keyLen)

	if cipherType == typeAES {
		cipher, err := aes.NewCipher(key)
		if err != nil {
			return 0, err
		}
		return cipher.BlockSize(), nil
	}
	if cipherType == typeDES {
		cipher, err := des.NewCipher(key)
		if err != nil {
			return 0, err
		}
		return cipher.BlockSize(), nil
	}

	return 0, errors.New("invalid cipher type")
}

const (
	typeAES = iota
	typeDES
)

func getCipherTypeName(cipherType int) string {
	switch cipherType {
	case typeAES:
		return "AES"
	case typeDES:
		return "DES"
	}
	return "unknown"
}

func test(keyLen, cipherType int) {
	fmt.Printf(
		"Getting block size of %v cipher with key length %v...\n",
		getCipherTypeName(cipherType),
		keyLen,
	)
	blockSize, err := getBlockSize(keyLen, cipherType)
	if err != nil {
		fmt.Println(err)
		fmt.Println("========")
		return
	}
	fmt.Println("Block size:", blockSize)
	fmt.Println("========")
}
