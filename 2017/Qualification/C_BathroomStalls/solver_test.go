package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in string
		n  int
		k  int
	}{
		{"4 2", 4, 2},
		{"1000 1", 1000, 1},
	}
	for _, test := range tests {
		sol := solver{}
		_ = sol.parse(test.in)
		if sol.n != test.n {
			t.Errorf("got %d, want %d", sol.n, test.n)
		}

		if sol.k != test.k {
			t.Errorf("got %d, want %d", sol.k, test.k)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		in  int
		max int
		min int
	}{
		{1, 0, 0},
		{2, 1, 0},
		{3, 1, 1},
		{4, 2, 1},
	}
	for _, test := range tests {
		max, min := split(test.in)
		if test.max != max {
			t.Errorf("got %d, want %d", max, test.max)
		}

		if test.min != min {
			t.Errorf("got %d, want %d", min, test.min)
		}
	}
}

func TestSelectSeq(t *testing.T) {
	tests := []struct {
		seqs  map[int]int
		seq   int
		rSeqs map[int]int
	}{
		{map[int]int{1: 2}, 1, map[int]int{1: 1}},
		{map[int]int{1: 2, 2: 1}, 2, map[int]int{1: 2}},
	}
	for _, test := range tests {
		seq, rSeqs := selectSeq(test.seqs)

		if seq != test.seq {
			t.Errorf("got %d, want %d", seq, test.seq)
		}

		if rSeqs[1] != test.rSeqs[1] {
			t.Errorf("got %d, want %d", rSeqs, test.rSeqs)
		}
	}
}
