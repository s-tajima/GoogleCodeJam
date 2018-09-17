package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		in string
		n  int
	}{
		{"132", 132},
	}
	for _, test := range tests {
		sol := solver{}
		_ = sol.parse(test.in)
		if sol.n != test.n {
			t.Errorf("got %d, want %d", sol.n, test.n)
		}
	}
}
