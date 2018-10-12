// TODO Write unit test
package main

import "fmt"

func main() {
	//s := "aab"
	s := "aababcabcd"
	m := NewHuffmanTree([]byte(s))
	fmt.Println(m)
	codebook := m.GetCodebook()
	fmt.Println(codebook)
}
