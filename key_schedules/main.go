package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	// represent roundnumber as (a single) byte
	roundNumByte := byte(roundNumber)

	for i := 0; i < len(masterKey); i++ {
		masterKey[i] = masterKey[i] ^ roundNumByte
	}

	return masterKey
}

// A key schedule is an algorithm that a block cipher employs to split an original key into multiple "round keys" or "subkeys". These round keys are deterministically derived from the original key, meaning that the same original key will always produce the same round keys.
