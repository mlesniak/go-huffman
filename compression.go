package main

import (
	"fmt"
	"os"
)

func Demo(filename string, bits []int) {
	// Store bit stream into buffer.
	// Consume 8 bits as a byte.
	buffer := make([]byte, 0)
	b := 0
	j := 0
	for i := 0; i < len(bits); i++ {
		b = b | bits[i]
		if j != 7 {
			b = b << 1
		}
		if j == 7 {
			buffer = append(buffer, byte(b))
			b = 0
			j = 0
		} else {
			j++
		}
	}

	fmt.Println(len(buffer))

	// Write to file.
	file, _ := os.Create(filename)
	defer file.Close()
	file.Write(buffer)
}
