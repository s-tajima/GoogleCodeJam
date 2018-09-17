package main

import (
	"fmt"
	"strconv"
	"strings"
)

type solver struct {
	s string // Represents the row of pancakes.
	k int    // Flips exactly K consecutive pancakes.
}

func (sol *solver) parse(str string) error {
	s := strings.Split(str, " ")
	k, err := strconv.Atoi(s[1])
	if err != nil {
		return err
	}
	sol.s = s[0]
	sol.k = k

	return nil
}

func (sol *solver) solve() (string, error) {
	fmt.Printf("s: %s k: %d\n", sol.s, sol.k)
	return sol.s, nil
}
