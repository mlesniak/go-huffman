// TODO ML Write unit test
// TODO ML Linter
// TODO ML go fmt
package main

import (
	"fmt"
)

func main() {
	//s := "aab"
	s := "aababcabcd"
	m := NewHuffmanTree([]byte(s))
	//fmt.Println(m)
	codebook := m.GetCodebook()

	//file, _ := os.Create("out.bit")

	// byte, 3 bit for length, bit for codebook
	for byteValue, code := range codebook {
		encLen := intToBinary(len(code))
		for len(encLen) < 3 {
			encLen = append([]int{0}, encLen...)
		}
		fmt.Println(byteValue, encLen, code)
	}

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

func intToBinary(value int) []int {
	buffer := make([]int, 0)

	fmt.Println(value)
	for {
		if value == 0 {
			break
		}
		buffer = append([]int{value & 1}, buffer...)
		value = value >> 1
	}

	fmt.Println(buffer)
	return buffer
}
