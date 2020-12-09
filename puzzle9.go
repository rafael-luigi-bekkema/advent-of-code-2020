package main

import "strconv"

func puzzle9(data []string, preamble int) int {
	nmbrs := make([]int, 0, len(data))
	find := func(inp []int, target int) bool {
		for _, i := range inp {
			for _, j := range inp {
				if i == j {
					continue
				}
				if i+j == target {
					return true
				}
			}
		}
		return false
	}

	for i, line := range data {
		n, _ := strconv.Atoi(line)
		nmbrs = append(nmbrs, n)

		if i < preamble {
			continue
		}

		if !find(nmbrs[i-preamble:], n) {
			return n
		}
	}

	panic("result not found")
}

func Puzzle9() int {
	data := readFile("input/input09")
	return puzzle9(data, 25)
}
