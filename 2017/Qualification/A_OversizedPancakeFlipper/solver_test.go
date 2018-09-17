package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		str string
		s   string
		k   int
	}{
		{"+- 1", "+-", 1},
		{"+-+- 2", "+-+-", 2},
	}
	for _, test := range tests {
		sol := solver{}
		_ = sol.parse(test.str)
		if sol.s != test.s {
			t.Errorf("got %v, want %v", sol.s, test.s)
		}
	}
}
