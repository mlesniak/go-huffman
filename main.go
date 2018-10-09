// TODO Write unit test
package main

import "fmt"

func main() {
	//s := "aab"
	s := "aababcabcd"
	m := GenerateHuffmanTree([]byte(s))
	fmt.Println(m)
}
