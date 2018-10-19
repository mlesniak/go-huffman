package main

import (
	"os"
)

func WriteBits(filename string, bits []int) {
	// Store bit stream into buffer.
	// Consume 8 bits as a byte.
	buffer := make([]byte, 0)
	b := 0
	j := 0
	for i := 0; i < len(bits); i++ {
		b = b | bits[i]
		if j != 7 {
			// Shift on everything but the last byte.
			b = b << 1
		} else if j == 7 {
			buffer = append(buffer, byte(b))
			b = 0
			j = 0
		} else {
			j++
		}
	}

	// Pad remaining byte.
	if len(bits)%8 != 0 {
		shiftLeft := 8 - j - 1
		b = b << uint(shiftLeft)
		buffer = append(buffer, byte(b))
	}

	// Write to file.
	file, _ := os.Create(filename)
	defer file.Close()
	file.Write(buffer)
}
