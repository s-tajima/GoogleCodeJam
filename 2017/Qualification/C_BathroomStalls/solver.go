package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type solver struct {
	n int // Number of stalls.
	k int // Number of people are about to enter the bathroom.
}

func (sol *solver) parse(str string) error {

	p := strings.Split(str, " ")

	n, err := strconv.Atoi(p[0])
	if err != nil {
		return err
	}

	k, err := strconv.Atoi(p[1])
	if err != nil {
		return err
	}

	sol.n = n
	sol.k = k
	return nil

}

func (sol *solver) solve() string {
	seqs := map[int]int{sol.n:1}

	max := 0
	min := 0

	for i := 0; i < sol.k; i++ {
		var seq int
		seq, seqs = selectSeq(seqs)
		max, min = split(seq)
		seqs = addSeq(max, seqs)
		seqs = addSeq(min, seqs)
	}

	return fmt.Sprintf("%d %d", max, min)
}

func split(seq int) (int, int) {
	rest := seq - 1

	max := int(math.Ceil(float64(rest) / 2))
	min := int(math.Trunc(float64(rest) / 2))

	return max, min
}

func addSeq(seq int, seqs map[int]int) (map[int]int) {
	seqs[seq]++
	return seqs
}

func selectSeq(seqs map[int]int) (int, map[int]int) {
	max := findMax(seqs)
	seqs[max]--

	return max, seqs
}

func findMax(seqs map[int]int) int {
	max := 0
	for k, v := range seqs {
		if k > max && v > 0 {
			max = k
		}
	}
	return max
}
