package main

import (
	"errors"
)

// decrypt one character at a time
func crypt(textCh, keyCh <-chan byte, result chan<- byte) {

	defer close(result)

	for {
		textByte, ok := <-textCh
		// once there is nothing more to read from textCh, break loop
		if !ok {
			return
		}
		keyByte, ok := <-keyCh
		if !ok {
			return
		}
		result <- textByte ^ keyByte
	}
}

// encrypt function reads from channels, performs XOR encryption using the crypt function
func encrypt(plaintext, key []byte) ([]byte, error) {
	if len(plaintext) != len(key) {
		return nil, errors.New("plaintext and key must be the same length")
	}

	plaintextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(plaintextCh)
		for _, v := range plaintext {
			plaintextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(plaintextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}

// decrypt function performs XOR decryption using the crypt function
func decrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) != len(key) {
		return nil, errors.New("ciphertext and key must be the same length")
	}

	ciphertextCh := make(chan byte)
	keyCh := make(chan byte)
	result := make(chan byte)

	go func() {
		defer close(ciphertextCh)
		for _, v := range ciphertext {
			ciphertextCh <- v
		}
	}()

	go func() {
		defer close(keyCh)
		for _, v := range key {
			keyCh <- v
		}
	}()

	go crypt(ciphertextCh, keyCh, result)

	res := []byte{}
	for v := range result {
		res = append(res, v)
	}
	return res, nil
}
