// TODO Write unit test
package main

func main() {
	//s := "aab"
	//s := "aababcabcd"
	//m := NewHuffmanTree([]byte(s))
	//fmt.Println(m)
	//codebook := m.GetCodebook()
	//fmt.Println(codebook)

	// $ xxd -b out.bit
	// 00000000: 11111111
	bits := []int{1, 1, 1, 1, 1, 1, 1, 1}
	Demo("out.bit", bits)

	// In a file compression.go
	// Find out how to write bits to file
	// Write and think about header
	// Write header to file
	// Write codebook to file
	// Encode contents

}
