package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	// hash the message sha-256
	hash := sha256.New()
	hash.Write([]byte(message))

	// create signature of hashed message with priv key
	// func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash.Sum(nil))
	if err != nil {
		return "", err
	}

	// return in format: MESSAGE.signature -- signature in lowercase hex
	// func EncodeToString(src []byte) string
	encodedSig := hex.EncodeToString(signature)
	formattedSig := fmt.Sprintf("%s.%s", message, encodedSig)

	return formattedSig, nil
}

func verifyECDSAMessage(token string, publicKey *ecdsa.PublicKey) error {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return errors.New("invalid token sections")
	}
	sig, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}
	hash := sha256.Sum256([]byte(parts[0]))

	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	if !valid {
		return errors.New("invalid signature")
	}
	return nil
}
