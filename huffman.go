package main

// TODO ML Generate huffman tree

import (
	// "fmt"
	"sort"
)

// HuffmanTree is a codification of a Huffman tree.
type HuffmanTree struct {
	Value byte
	Bit1  *HuffmanTree
	Bit0  *HuffmanTree
}

func (ht HuffmanTree) String() string {
	s1 := "";
	s2 := "";
	if (ht.Bit0 != nil) {
		s1 = "left: " + ht.Bit0.String()
	}
	if (ht.Bit1 != nil) {
		s2 = ", right: " + ht.Bit1.String()
	}
	return "{" + string(ht.Value) + s1 + s2 + "}"
}

func makeLeaf(value byte) HuffmanTree {
	return HuffmanTree{Value: value}
}

func combineLeafes(l1, l2 HuffmanTree) HuffmanTree {
	return HuffmanTree{0, &l1, &l2}
}

// GenerateHuffmanTree generates a HuffmanTree.
func GenerateHuffmanTree(s []byte) HuffmanTree {
	frequencies := ComputeFrequency(s)
	list := SortFrequencyList(frequencies)
	
	// Generate leafes.
	var leafes []HuffmanTree
	for _, value := range list {
		leafes = append(leafes, makeLeaf(value))
	}
	
	// Combine single leaves to a tree.
	tree := leafes[0]
	for _, leafValue := range leafes[1:] {
		tree = combineLeafes(tree, leafValue)
	}
	
	return tree
}

// SortFrequencyList generates a list of bytes sorted by relative frequency, starting with the smallest ones.
func SortFrequencyList(m map[byte]float32) []byte {
	// Convert to array of struct values.
	type kv struct {
		Key   byte
		Value float32
	}
	var kvs []kv
	for k, v := range m {
		kvs = append(kvs, kv{k, v})
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
