package main

import (
	"bytes"
	"strconv"
	"strings"
)

type solver struct {
	s []byte // Represents the row of pancakes.
	k int    // Flips exactly K consecutive pancakes.
}

func (sol *solver) parse(str string) error {
	s := strings.Split(str, " ")
	k, err := strconv.Atoi(s[1])
	if err != nil {
		return err
	}
	sol.s = []byte(s[0])
	sol.k = k

	return nil
}

func (sol *solver) solve() string {
	y := trimRecursive(sol.s, sol.k, 0)
	if y < 0 {
		return "IMPOSSIBLE"
	}
	return strconv.Itoa(y)
}

func trimRecursive(s []byte, k int, y int) int {
	s = trimHappySide(s)

	if len(s) == 0 {
		return y
	}

	if len(s) < k {
		return -1
	}

	s = flipK(s, k)
	return trimRecursive(s, k, y+1)
}

func trimHappySide(s []byte) []byte {
	return bytes.TrimLeft(s, "+")
}

func flipK(s []byte, k int) []byte {
	ns := []byte{}
	for _, b := range s[0:k] {
		ns = append(ns, flip(b))
	}
	s = append(ns, s[k:]...)

	return s
}

func flip(p byte) byte {
	if p == '+' {
		return '-'
	}
	return '+'
}
