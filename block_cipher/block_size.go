package main

func padWithZeros(block []byte, desiredSize int) []byte {

	length := len(block)
	dif := desiredSize - length

	// while loop that keeps pumping 0s into block until it's the desired size
	for dif > 0 {
		block = append(block, 0)
		dif--
	}

	return block
}
