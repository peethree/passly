package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {

	// curve
	curve := elliptic.P256()

	// generate key
	// func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)
	// rather than making a random byte slice, pass in rand.Reader directly
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// type PrivateKey struct {
	// 	PublicKey
	// 	D *big.Int
	// }
	publicKey := privateKey.PublicKey

	return &publicKey, privateKey, nil
}

// keysArePaired verifies if the public and private keys are paired using ECDSA.
func keysArePaired(pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey) bool {
	msg := "a test message"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		return false
	}

	return ecdsa.VerifyASN1(pubKey, hash[:], sig)
}
