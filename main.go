// TODO ML Write unit test
// TODO ML Logging
// TODO ML Linter
// TODO ML go fmt
package main

import (
	"fmt"
)

func main() {
	//s := "aab"
	s := []byte("aababcabcd")
	m := NewHuffmanTree(s)
	fmt.Println(m)
	codebook := m.GetCodebook()
	// Bug: map[97:[1] 98:[0 0] 99:[0 0 0] 100:[0 0 0]]
	fmt.Println(s)
	fmt.Println(codebook)

	//file, _ := os.Create("out.bit")

	//WriteCodebook(file, codebook)

	// Generate bitstream for each character with respect of the codebook.
	//WriteData(file, []byte(s), codebook)

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
