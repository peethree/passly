package main

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
)

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	// get 2 big prime numbers of given keysize
	// func getBigPrime(bits int) (*big.Int, error) {
	p, _ := getBigPrime(keysize)

	q, _ := getBigPrime(keysize)

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	n := new(big.Int)
	return n.Mul(p, q)
}

// firstNDigits returns the first 'numDigits' digits of the big integer n.
func firstNDigits(n big.Int, numDigits int) string {
	if len(n.String()) < numDigits {
		return fmt.Sprintf("%v", n.String())
	}
	return fmt.Sprintf("%v...", n.String()[:numDigits])
}

var randReader = mrand.New(mrand.NewSource(0))

// getBigPrime generates a random prime number of the given size.
func getBigPrime(bits int) (*big.Int, error) {
	if bits < 2 {
		return nil, errors.New("prime size must be at least 2-bit")
	}
	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}
	bytes := make([]byte, (bits+7)/8)
	p := new(big.Int)
	for {
		if _, err := io.ReadFull(randReader, bytes); err != nil {
			return nil, err
		}
		bytes[0] &= uint8(int(1<<b) - 1)
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		bytes[len(bytes)-1] |= 1
		p.SetBytes(bytes)
		if p.ProbablyPrime(20) {
			return p, nil
		}
	}
}
