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

func TestFrequencyMultiple(t *testing.T) {
	m := ComputeFrequency([]byte("aba"))
	
	// If we generalize this, we should use epsilon-based comparison.
	if m[byte('a')] != 0.6666667 {
		t.Error("'a' frequency should be 0.66 but is", m[byte('a')])
	}

	if m[byte('b')] != 0.33333334 {
		t.Error("'b' frequency should be 0.33 but is", m[byte('b')])
	}
}