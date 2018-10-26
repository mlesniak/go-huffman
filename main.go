// TODO ML Write unit test
// TODO ML Logging
// TODO ML Linter
// TODO ML go fmt
// TODO ML Use abstraction instead of file, e.g. a writer
// TODO ML Consistent encoding, e.g. always the same order in the codebook
// TODO ML Streaming instead of full read into RAM? Or non-feature?
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
	file, _ := os.Open("/tmp/out.bit")
	defer file.Close()

	bits := ReadBits(file)
	codebook := ReadCodebook(bits)
	fmt.Println(codebook)

	// Decode codebook.
	// Use codebook and bit stream.

	return []byte("<Nothing yet>")
}

func encode() {
	s := []byte("aababcabcd")
	m := NewHuffmanTree(s)
	codebook := m.GetCodebook()
	file, _ := os.Create("/tmp/out.bit")
	defer file.Close()
	WriteCodebook(file, codebook)
	WriteData(file, []byte(s), codebook)
}
