package main

import (
	"fmt"
	"io"
	"os"
)

func ReadBits(r io.Reader) []int8 {
	buffer := make([]int8, 0)

	byteBuffer := make([]byte, 1024)
	for {
		count, err := r.Read(byteBuffer)
		if err == io.EOF {
			break
		}

		for _, b := range byteBuffer[:count] {
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

// TODO ML Return bits read
func ReadCodebook(stream []int8) map[byte][]int8 {
	codebook := make(map[byte][]int8)

	fmt.Println(stream)
	length := binaryToInt(stream[:8])
	pos := 8
	for i := 0; i < length; i++ {
		byteBinary := byte(binaryToInt(stream[pos:pos+8]))
		encLen := binaryToInt(stream[pos+8:pos+8+3])
		fmt.Println(byteBinary, encLen)
		if encLen == 0 {
			codebook[byteBinary] = []int8{1}
			pos += 8 + 3
		} else {
			// TODO ML Continue here
			codebook[byteBinary] = stream[]
		}
	}

	return codebook
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

func binaryToInt(bits []int8) int {
	num := 0

	p := 1
	for i := len(bits) - 1; i >= 0; i-- {
		num += p * int(bits[i])
		p = p << 1
	}

	return num
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
