package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	// split the hex string
	stringParts := strings.Split(s, ":")

	var result []string

	// append all the parts that were split to result
	result = append(result, stringParts...)
	fmt.Printf("%s", result)
	// [48 65 6c 6c 6f][57 6f 72 6c 64][50 61 73 73 77 6f 72 64][ZZ YY XX]-

	joinedResult := strings.Join(result, "")
	fmt.Printf("%s", joinedResult)
	// [48656c6c6f]

	bitPart, err := hex.DecodeString(joinedResult)
	if err != nil {
		return []byte{}, err
	}
	return bitPart, nil
}
