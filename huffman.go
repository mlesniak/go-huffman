package main

import (
	"sort"
)

// HuffmanTree is a codification of a Huffman tree. We arbitrary decide that we interpret the left tree as bit 1 and
// the right tree as bit 0.
type HuffmanTree struct {
	Value byte
	Left  *HuffmanTree
	Right *HuffmanTree
}

func (tree HuffmanTree) String() string {
	s1 := ""
	s2 := ""
	v := ""

	if !tree.isLeaf() {
		s1 = "left: " + tree.Right.String()
		// We assume that we always have a left tree if we have a right tree.
		s2 = ", right: " + tree.Left.String()
	} else {
		v = string(tree.Value)
	}

	return "{" + v + s1 + s2 + "}"
}

func (tree HuffmanTree) isLeaf() bool {
	return tree.Right == nil && tree.Left == nil
}

// GenerateHuffmanTree generates a HuffmanTree based on the data passed in the parameter.
func NewHuffmanTree(s []byte) HuffmanTree {
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
