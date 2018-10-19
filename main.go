// TODO ML Write unit test
// TODO ML Linter
// TODO ML go fmt
package main

import "os"

func main() {
	//s := "aab"
	//s := "aababcabcd"
	//m := NewHuffmanTree([]byte(s))
	//fmt.Println(m)
	//codebook := m.GetCodebook()
	//fmt.Println(codebook)

	// $ xxd -b out.bit
	// 00000000: 11111111
	bits := []int{
		1, 0, 1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 0, 1}
	file, _ := os.Create("out.bit")
	WriteBits(file, bits)

	// In a file compression.go
	// Find out how to write bits to file
	// Write and think about header
	// Write header to file
	// Write codebook to file
	// Encode contents

}
