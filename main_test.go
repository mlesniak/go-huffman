package main

import "testing"

func TestFrequencyEmpty(t *testing.T) {
	m := ComputeFrequency([]byte(""))
	if len(m) != 0 {
		t.Error("len is not zero")
	}
}