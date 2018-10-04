package main

import "sort"

// SortFrequencyList generates a list of bytes sorted by relative frequency, starting with the smallest ones.
func SortFrequencyList(m map[byte]float32) []byte {
	// Convert to array of struct values.
	type kv struct {
		Key byte
		Value float32
	}
	var kvs []kv
	for k,v := range m {
		kvs = append(kvs, kv{k,v})
	}

	// Sort array.
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value < kvs[j].Value
	})

	// Convert to byte array.
	var result []byte
	for _, v := range kvs {
		result = append(result, v.Key)
	}
	return result
}

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
