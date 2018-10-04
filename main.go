// TODO Write unit test
package main

import "fmt"

func main() {
	s := "Huffman encoding and decoding"
	m := ComputeFrequency([]byte(s))
	fmt.Println(m)
}