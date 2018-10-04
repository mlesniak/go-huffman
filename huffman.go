package main

// ComputeFrequency computes a relative frequency map of all characters in the array.
func ComputeFrequency(s []byte) map[byte]float32 {
	m := make(map[byte]float32)

	// Count absolute number.
	for _, v := range s {
		m[v]++
	}
	// Compute relative number of occurrences.
	for k, v := range m {
		m[k] = v / float32(len(s))
	}

	return m
}
