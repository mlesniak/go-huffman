// TODO Write unit test
package main

import "fmt"

func main() {
	//s := "aab"
	s := "aababcabcd"
	m := NewHuffmanTree([]byte(s))
	fmt.Println(m)
	codebook := m.GetCodebook()

	// In a file compression.go
	// Find out how to write bits to file
	// Write and think about header
	// Write header to file
	// Write codebook to file
	// Encode contents

	fmt.Println(codebook)
}
