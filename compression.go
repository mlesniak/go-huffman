package main

import (
	"fmt"
	"io"
	"os"
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

func WriteCodebook(file *os.File, codebook map[byte][]int8) {
	codeBits := make([]int8, 0)
	// byte, 3 bit for length, bit for codebook
	for byteValue, code := range codebook {
		// We always have at least one bit.
		encLen := intToBinary(len(code) - 1)
		for len(encLen) < 3 {
			encLen = append([]int8{0}, encLen...)
		}
		fmt.Println(byteValue, encLen, code)
		i := int(byteValue)

		byteBinary := intToBinary(i)
		for len(byteBinary) < 8 {
			byteBinary = append([]int8{0}, byteBinary...)
		}
		// TODO ML Future optimization: if encLen = 0 0 0 , we don't need the code, since it's unique.
		codeBits = append(codeBits, byteBinary...)
		codeBits = append(codeBits, encLen...)
		codeBits = append(codeBits, code...)
	}
	// Append missing bits.
	for len(codeBits)%8 != 0 {
		codeBits = append(codeBits, 0)
	}
	//fmt.Println(codeBits)
	WriteBits(file, codeBits)
}

func WriteData(file *os.File, bytes []byte, codebook map[byte][]int8) {
	dataBuffer := make([]int8, 0)
	length := padLeft(intToBinary(len(bytes)), 8)
	dataBuffer = append(dataBuffer, length...)
	for _, byteValue := range bytes {
		dataBuffer = append(dataBuffer, codebook[byteValue]...)
	}
	fmt.Println(dataBuffer)
	WriteBits(file, dataBuffer)
}

func intToBinary(value int) []int8 {
	buffer := make([]int8, 0)

	for {
		if value == 0 {
			break
		}
		lsb := value % 2
		buffer = append([]int8{int8(lsb)}, buffer...)
		value = value >> 1
	}

	return buffer
}

func padLeft(bits []int8, size int) []int8 {
	if len(bits) > size {
		panic("len larger than padding size")
	}
	for len(bits) < size {
		bits = append([]int8{0}, bits...)
	}
	return bits
}
