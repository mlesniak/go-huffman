// TODO Write unit test
package main

import "fmt"

func main() {
	s := "Huffman encoding and decoding"
	m := computeFrequency([]byte(s))
	fmt.Println(m)
}

func computeFrequency(s []byte) map[byte]float32 {
	m := make(map[byte]float32)

	// Count absolute number.
	for _, v := range(s) {
		m[v]++
	}
	
	// Compute relative number of occurrences.
	for k, v := range(m) {
		m[k] = v / float32(len(s))
	}

	return m
}