// TODO ML Write unit test
// TODO ML Logging
// TODO ML Linter
// TODO ML go fmt
package main

import (
	"fmt"
	"os"
)

func main() {
	//s := "aab"
	s := "aababcabcd"
	m := NewHuffmanTree([]byte(s))
	//fmt.Println(m)
	codebook := m.GetCodebook()

	file, _ := os.Create("out.bit")

	codeBits := make([]int8, 0)
	// byte, 3 bit for length, bit for codebook
	for byteValue, code := range codebook {
		// We always have at least one bit.
		encLen := intToBinary(int8(len(code) - 1))
		for len(encLen) < 3 {
			encLen = append([]int8{0}, encLen...)
		}
		fmt.Println(byteValue, encLen, code)
		i := int8(byteValue)

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

	// $ xxd -b out.bit
	// 00000000: 11111111
	//bits := []int{
	//	1, 0, 1, 0, 0, 1, 1, 0,
	//	1, 0, 0, 0, 0, 1}
	//WriteBits(file, bits)

	// In a file compression.go
	// Find out how to write bits to file
	// Write and think about header
	// Write header to file
	// Write codebook to file
	// Encode contents

}

func intToBinary(value int8) []int8 {
	buffer := make([]int8, 0)

	for {
		if value == 0 {
			break
		}
		buffer = append([]int8{value & 1}, buffer...)
		value = value >> 1
	}

	return buffer
}
