package main

import (
	"sort"
)

// HuffmanTree is a codification of a Huffman tree. Since we want to use the tree for compression, we code the
// bit-based path directly into the names of the variables. Note to myself: it this a good design decision or is this
// over-specialization; the rest of the functions are more generalized?
type HuffmanTree struct {
	Value byte // Pointer to byte for clarity? Although it costs more memory? Or a simple isLeaf() function?
	Bit1  *HuffmanTree
	Bit0  *HuffmanTree
}

// TODO ML better formatting
func (ht HuffmanTree) String() string {
	s1 := ""
	if ht.Bit0 != nil {
		s1 = "left: " + ht.Bit0.String()
	}

	s2 := ""
	if ht.Bit1 != nil {
		s2 = ", right: " + ht.Bit1.String()
	}

	v := string(ht.Value)
	if !isLeaf(ht) {
		v = ""
	}

	return "{" + v + s1 + s2 + "}"
}
func isLeaf(tree HuffmanTree) bool {
	return tree.Bit0 == nil && tree.Bit1 == nil
}

// GenerateHuffmanTree generates a HuffmanTree.
// TODO ML Name it MakeHuffmanTree for consistency? Or even New...?
func GenerateHuffmanTree(s []byte) HuffmanTree {
	frequencies := ComputeFrequency(s)
	byteList := SortFrequencyList(frequencies)

	// Combine single sorted lists to a single binary tree.
	root := makeLeaf(byteList[0])
	for _, leafValue := range byteList[1:] {
		leaf := makeLeaf(leafValue)
		root = combineLeafs(root, leaf)
	}

	return root
}

// makeLeaf generate a single leaf without children.
func makeLeaf(value byte) HuffmanTree {
	return HuffmanTree{Value: value}
}

// combineLeafs combines two trees into a single new one without a value. Note that we intentionally pass parameters
// by value to have 'fresh' trees.
func combineLeafs(left, right HuffmanTree) HuffmanTree {
	return HuffmanTree{0, &left, &right}
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
