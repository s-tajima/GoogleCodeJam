package main

import (
	"bytes"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		in string
		s  []byte
		k  int
	}{
		{"+- 1", []byte("+-"), 1},
		{"+-+- 2", []byte("+-+-"), 2},
	}
	for _, test := range tests {
		sol := solver{}
		_ = sol.parse(test.in)
		if !bytes.Equal(sol.s, test.s) {
			t.Errorf("got %v, want %v", sol.s, test.s)
		}
	}
}

func TestTrimRecursive(t *testing.T) {
	tests := []struct {
		s      []byte
		k      int
		expect int
	}{
		{[]byte("---+-++-"), 3, 3},
		{[]byte("+++++"), 4, 0},
		{[]byte("-+-+-"), 4, -1},
	}

	for _, test := range tests {
		actual := trimRecursive(test.s, test.k, 0)
		if actual != test.expect {
			t.Errorf("got %d, want %d", actual, test.expect)
		}
	}
}

func TestTrimHappySide(t *testing.T) {
	tests := []struct {
		in     []byte
		expect []byte
	}{
		{[]byte("+-"), []byte("-")},
		{[]byte("--+"), []byte("--+")},
		{[]byte("+-+-"), []byte("-+-")},
	}

	for _, test := range tests {
		actual := trimHappySide(test.in)
		if !bytes.Equal(actual, test.expect) {
			t.Errorf("got %s, want %s", actual, test.expect)
		}
	}
}

func TestFlipK(t *testing.T) {
	tests := []struct {
		s      []byte
		k      int
		expect []byte
	}{
		{[]byte("+-"), 1, []byte("--")},
		{[]byte("--+"), 2, []byte("+++")},
		{[]byte("+-+-"), 4, []byte("-+-+")},
	}

	for _, test := range tests {
		actual := flipK(test.s, test.k)
		if !bytes.Equal(actual, test.expect) {
			t.Errorf("got %s, want %s", actual, test.expect)
		}
	}
}

func TestFlip(t *testing.T) {
	tests := []struct {
		in     byte
		expect byte
	}{
		{'+', '-'},
		{'-', '+'},
		{'?', '#'},
	}

	for _, test := range tests {
		actual := flip(test.in)
		if actual != test.expect {
			t.Errorf("got %c, want %c", actual, test.expect)
		}
	}
}
