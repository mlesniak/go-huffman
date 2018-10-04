package main

import "testing"

func TestFrequencyEmpty(t *testing.T) {
	m := ComputeFrequency([]byte(""))
	if len(m) != 0 {
		t.Error("len is not zero")
	}
}

func TestFrequencySingle(t *testing.T) {
	m := ComputeFrequency([]byte("a"))
	
	if m[byte('a')] != 1.0 {
		t.Error("Single frequency should be 1.0")
	}
}