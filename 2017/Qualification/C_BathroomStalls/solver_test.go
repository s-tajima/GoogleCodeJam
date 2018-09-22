package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in string
		n  uint64
		k  uint64
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
		in  uint64
		max uint64
		min uint64
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

func TestSolve(t *testing.T) {
	tests := []struct {
		n uint64
		k uint64
		expect string
	}{
		{4, 2, "1 0"},
		{5, 2, "1 0"},
		{6, 2, "1 1"},
		{1000, 1000, "0 0"},
		{1000, 1, "500 499"},
	}
	for _, test := range tests {
		sol := solver{}
		sol.n = test.n
		sol.k = test.k

		actual := sol.solve()

		if actual != test.expect {
			t.Errorf("got %s, want %s", actual, test.expect)
		}
	}
}

func TestSelectSeq(t *testing.T) {
	tests := []struct {
		seqs  map[uint64]uint64
		max   uint64
		num   uint64
		rSeqs map[uint64]uint64
	}{
		{map[uint64]uint64{1: 2}, 1, 2, map[uint64]uint64{}},
		{map[uint64]uint64{1: 2, 2: 1}, 2, 1, map[uint64]uint64{1: 2}},
	}
	for _, test := range tests {
		max, num, rSeqs := selectSeq(test.seqs)

		if max != test.max {
			t.Errorf("got %d, want %d", max, test.max)
		}

		if num != test.num {
			t.Errorf("got %d, want %d", num, test.num)
		}

		if rSeqs[1] != test.rSeqs[1] {
			t.Errorf("got %d, want %d", rSeqs, test.rSeqs)
		}
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		seqs map[uint64]uint64
		max  uint64
	}{
		{map[uint64]uint64{1: 2}, 1},
		{map[uint64]uint64{1: 2, 2: 1}, 2},
	}
	for _, test := range tests {
		max := findMax(test.seqs)

		if max != test.max {
			t.Errorf("got %d, want %d", max, test.max)
		}
	}
}
