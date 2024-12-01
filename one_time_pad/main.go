package main

// plaintext [1101 0011 0111], key [1001 0110 1110] for example
func crypt(plaintext, key []byte) []byte {

	byteResult := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		byteResult[i] = plaintext[i] ^ key[i]
	}

	return byteResult
}

// byteResult[i] = plaintext[i] ^ key[i] performs a bitwise XOR of the i-th byte of plaintext and key. The result is stored in the i-th position of byteResult

// The big problem with the One Time Pad is that the key needs to be the same length as the message. That means to encrypt a 128 GB hard drive, I'd need a 128 GB key!! That's just not practical.
