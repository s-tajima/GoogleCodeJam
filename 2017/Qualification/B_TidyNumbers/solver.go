package main

import (
	"strconv"
	"strings"
)

type solver struct {
	n []int // "She just finished counting all positive integers in ascending order from 1 to N."
}

func (sol *solver) parse(str string) error {
	for _, s := range strings.Split(str, "") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		sol.n = append(sol.n, i)
	}
	return nil
}

func (sol *solver) solve() string {
	for {
		sol.n = mod(sol.n)
		if check(sol.n) {
			return convert(sol.n)
		}
	}
	return ""
}

func mod(n []int) []int {
	res := []int{}
	for i := range n {
		if len(n) <= i+1 {
			res = append(res, n[i])
			break
		}

		if n[i] > n[i+1] {
			res = append(res, n[i]-1)

			size := len(n) - i - 1
			remains := make([]int, size)
			for i := range remains {
				remains[i] = 9
			}
			res = append(res, remains...)
			break
		}

		res = append(res, n[i])
	}

	return res
}

func check(res []int) bool {
	for i := range res {
		if len(res) <= i+1 {
			return true
		}
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}

func convert(is []int) string {
	str := ""
	for _, v := range is {
		str += strconv.Itoa(v)
	}
	return strings.TrimLeft(str, "0")
}
