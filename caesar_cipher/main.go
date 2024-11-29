package main

import (
	"strings"
)

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	var result []string

	for _, c := range text {
		result = append(result, getOffsetChar(c, key))
	}

	// join the results into 1 string
	return strings.Join(result, "")
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	for i, v := range alphabet {
		// character if found in alphabet
		if v == c {
			// in case of index greater than 26, or negative index: +26 % 26
			index := (i + offset + 26) % 26
			return string(alphabet[index])
		}
	}
	// char not found
	return ""
}
