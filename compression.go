package main

import (
	"io"
	"os"
)

func ReadBits(r io.Reader) []int8 {
	buffer := make([]int8, 0)

	byteBuffer := make([]byte, 1024)
	for {
		_, err := r.Read(byteBuffer)
		if err == io.EOF {
			break
		}

		for _, b := range byteBuffer {
			buffer = append(buffer, padLeft(intToBinary(int(b)), 8)...)
		}
	}

	return buffer
}

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

	// Encode codebook length.
	length := padLeft(intToBinary(len(codebook)), 8)
	codeBits = append(codeBits, length...)

	// byte, 3 bit for length, bit for codebook
	for byteValue, code := range codebook {
		// We always have at least one bit.
		encLen := padLeft(intToBinary(len(code)-1), 3)

		iv := int(byteValue)
		byteBinary := padLeft(intToBinary(iv), 8)

		codeBits = append(codeBits, byteBinary...)
		codeBits = append(codeBits, encLen...)
		if len(code) == 1 {
			// We have only one value with a code length of 1, i.e. the byte value with the highest
			// occurence, hence we do not need to encode it explicitly.
			continue
		}
		codeBits = append(codeBits, code...)
	}

	// Append missing bits (padRight with modulo).
	for len(codeBits)%8 != 0 {
		codeBits = append(codeBits, 0)
	}

	WriteBits(file, codeBits)
}

func WriteData(file *os.File, bytes []byte, codebook map[byte][]int8) {
	dataBuffer := make([]int8, 0)

	// Encode length.
	length := padLeft(intToBinary(len(bytes)), 8)
	dataBuffer = append(dataBuffer, length...)

	// Encode actual data.
	for _, byteValue := range bytes {
		dataBuffer = append(dataBuffer, codebook[byteValue]...)
	}

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
