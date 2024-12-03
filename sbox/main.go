package main

import (
	"errors"
)

func sBox(b byte) (byte, error) {
	sBox := map[byte]byte{
		0b0000: 0b00,
		0b0001: 0b10,
		0b0010: 0b01,
		0b0011: 0b11,
		0b0100: 0b10,
		0b0101: 0b00,
		0b0110: 0b11,
		0b0111: 0b01,
		0b1000: 0b01,
		0b1001: 0b11,
		0b1010: 0b00,
		0b1011: 0b10,
		0b1100: 0b11,
		0b1101: 0b01,
		0b1110: 0b10,
		0b1111: 0b00,
	}

	// get value from map, check if valid result
	if result, ok := sBox[b]; ok {
		return result, nil
	}
	return 0, errors.New("incorrect input")
}
