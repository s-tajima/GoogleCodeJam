package main

import (
	"fmt"
	"strconv"
	"strings"
)

type solver struct {
	n uint64 // Number of stalls.
	k uint64 // Number of people are about to enter the bathroom.
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

	sol.n = uint64(n)
	sol.k = uint64(k)
	return nil

}

func (sol *solver) solve() string {
	seqs := map[uint64]uint64{sol.n: 1}
	rest := sol.k

	max := uint64(0)
	min := uint64(0)

	for {
		var seq uint64
		var num uint64

		seq, num, seqs = selectSeq(seqs)

		max, min = split(seq)
		seqs[max] += num
		seqs[min] += num

		if rest <= num {
			break
		}
		rest -= num
	}
	return fmt.Sprintf("%d %d", max, min)
}

func split(seq uint64) (uint64, uint64) {
	max := uint64(seq / 2)
	min := uint64((seq - 1) / 2)

	return max, min
}

func selectSeq(seqs map[uint64]uint64) (uint64, uint64, map[uint64]uint64) {
	max := findMax(seqs)
	num := seqs[max]
	delete(seqs, max)
	return max, num, seqs
}

func findMax(seqs map[uint64]uint64) uint64 {
	max := uint64(0)
	for k := range seqs {
		if k > max {
			max = k
		}
	}
	return max
}
