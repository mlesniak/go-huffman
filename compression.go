package main

import (
	"io"
)

func WriteBits(w io.Writer, bits []int8) {
	buffer := make([]byte, 0)

	// Store bit stream into buffer. Consume 8 bits as a byte. Padding is done below.
	b := int8(0)
	j := 1
	for i := 0; i < len(bits); i++ {
		b = b | bits[i]

		if j != 8 {
			// Shift on everything but the last byte.
			b = b << 1
		}

		// If we collect 8 bits, store the byte.
		if j == 8 {
			buffer = append(buffer, byte(b))
			b = 0
			j = 1
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

	w.Write(buffer)
}
