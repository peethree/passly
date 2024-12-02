package main

import (
	"crypto/sha256"
)

func feistel(msg []byte, roundKeys [][]byte) []byte {
	// split msg into 2 equal parts left + right
	half := len(msg) / 2

	lhs := msg[0:half]
	rhs := msg[half:]

	// func hash(first, second []byte, outputLength int) []byte {
	for i := 0; i < len(roundKeys); i++ {
		nextRHS := xor(lhs, hash(rhs, roundKeys[i], len(roundKeys[i])))
		nextLHS := rhs
		// update rhs, lhs for next round's loop
		rhs = nextRHS
		lhs = nextLHS
	}
	// first append right side, then left side to the result
	// variadic / spread operator to add every element of rhs / lhs to result
	return append(rhs, lhs...)
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(first, second []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(first, second...))
	return h.Sum(nil)[:outputLength]
}
