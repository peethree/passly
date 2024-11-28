package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	var result []string
	for _, v := range b {
		// formats hexadecimal: 20
		result = append(result, fmt.Sprintf("%02x", v))
	}
	return strings.Join(result, ":")
}

func getBinaryString(b []byte) string {
	// formats binary: 00100000
	str := fmt.Sprintf("%08b\n", b)
	// split at space char
	splitStr := strings.Split(str, " ")
	// join strings together delimited by ":"
	result := strings.Join(splitStr, ":")
	// omit "[" and "]" from return str
	return string(result[1 : len(result)-2])
}
