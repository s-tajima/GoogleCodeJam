package main

import (
	"bytes"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		str string
		s   []byte
		k   int
	}{
		{"+- 1", []byte("+-"), 1},
		{"+-+- 2", []byte("+-+-"), 2},
	}
	for _, test := range tests {
		sol := solver{}
		_ = sol.parse(test.str)
		if !bytes.Equal(sol.s, test.s) {
			t.Errorf("got %v, want %v", sol.s, test.s)
		}
	}
}
