package main

import (
	"fmt"
	"os"
)

func Demo(filename string, bits []int) {
	if len(bits)%8 != 0 {
		panic("not supported yet")
	}

	// Store bit stream into buffer.
	// Consume 8 bits as a byte.
	buffer := make([]byte, 0)
	for i := 0; i < len(bits)/8; i += 8 {
		b := 0
		for j := 0; j < 8; j++ {
			b = b | bits[i*8+j]
			if j != 7 {
				b = b << 1
			}
		}
		buffer = append(buffer, byte(b))
	}

	fmt.Println(len(buffer))

	// Write to file.
	file, _ := os.Create(filename)
	defer file.Close()
	file.Write(buffer)
}
