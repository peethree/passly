package main

import (
	"crypto/rand"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	// tot = ϕ(n) = (p - 1) * (q - 1)
	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	one := big.NewInt(1)
	tot := new(big.Int)

	pMinusOne := a.Sub(p, one)
	qMinusOne := b.Sub(q, one)

	tot = c.Mul(pMinusOne, qMinusOne)

	return tot
}

func getE(tot *big.Int) *big.Int {

	two := big.NewInt(2)
	// tot - 2
	upperBound := new(big.Int).Sub(tot, two)

	for {
		// func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
		e, _ := rand.Int(rand.Reader, upperBound)

		// ensure e is greater than 1
		e.Add(e, two)

		// func gcd(x, y *big.Int) *big.Int {
		// calculates greatest common divisor between e n totient
		greatestCommonDivisor := gcd(e, tot)

		// e and tot have a greatest common divisor of 1
		if greatestCommonDivisor.Cmp(big.NewInt(1)) == 0 {
			return e
		}
	}
}

func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)
	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
	}
	return xCopy
}
