package main

import "fmt"

func main() {
	s := "Huffman encoding and decoding"
	m := count(s)
	fmt.Println(m)
}

func count(s string) map[string]float32 {
	m := make(map[string]float32)

	for _, v := range(s) {
		m[string(v)]++
	}

	return m
}