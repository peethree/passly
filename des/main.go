package main

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	// create new cipher block
	// func NewCipher(key []byte) (cipher.Block, error)
	newCipherBlock, err := des.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	blockSize := newCipherBlock.BlockSize()

	// pad plaintext with 0s
	paddedText := padMsg(plaintext, blockSize)

	// generate random iv (iv is a bunch of uninitialized 0s in this case), append to beginning of ciphertext (same length as blocksize)
	ciphertext := make([]byte, blockSize+len(paddedText))
	iv := ciphertext[:blockSize]

	// create new encrypter
	mode := cipher.NewCBCEncrypter(newCipherBlock, iv)
	mode.CryptBlocks(ciphertext[blockSize:], paddedText)
	// encrypt the blocks and return the entire ciphertext
	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	// func padWithZeros(block []byte, desiredSize int) []byte {
	// find the last block that needs padding
	if len(plaintext)%blockSize != 0 {
		index := len(plaintext) - (len(plaintext) % blockSize)
		finalBlock := plaintext[index:]

		// get just the full blocks portion from the plaintext
		fullBlocks := plaintext[:index]
		finalBlockPadded := padWithZeros(finalBlock, blockSize)

		// add the pieces back to together (... to add all elements)
		return append(fullBlocks, finalBlockPadded...)
	}
	return plaintext
}

func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}
