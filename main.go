// TODO ML Write unit test
// TODO ML Logging
// TODO ML Linter
// TODO ML go fmt
// TODO ML Use abstraction instead of file, e.g. a writer
package main

import (
	"fmt"
	"os"
)

func main() {
	encode()
	bytes := decode()
	fmt.Println(string(bytes))
}

func decode() []byte {
	file, _ := os.Open("out.bit")
	defer file.Close()

	return []byte("<Nothing yet>")
}

func encode() {
	s := []byte("aababcabcd")
	m := NewHuffmanTree(s)
	codebook := m.GetCodebook()
	file, _ := os.Create("out.bit")
	defer file.Close()
	WriteCodebook(file, codebook)
	WriteData(file, []byte(s), codebook)
}
