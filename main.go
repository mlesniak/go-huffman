// TODO ML Write unit test
// TODO ML Logging
// TODO ML Linter
// TODO ML go fmt
// TODO ML Use abstraction instead of file, e.g. a writer
package main

import (
	"os"
)

func main() {
	//s := "aab"
	s := []byte("aababcabcd")
	m := NewHuffmanTree(s)
	codebook := m.GetCodebook()

	file, _ := os.Create("out.bit")
	defer file.Close()

	WriteCodebook(file, codebook)
	WriteData(file, []byte(s), codebook)
}
