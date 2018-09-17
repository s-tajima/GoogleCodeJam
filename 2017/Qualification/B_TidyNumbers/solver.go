package main

import (
	"strconv"
)

type solver struct {
	n int // "She just finished counting all positive integers in ascending order from 1 to N."
}

func (sol *solver) parse(str string) error {
	n, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	sol.n = n
	return nil
}

func (sol *solver) solve() string {
	return strconv.Itoa(sol.n)
}
