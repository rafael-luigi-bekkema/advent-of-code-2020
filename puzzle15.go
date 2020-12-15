package main

import "strings"

func puzzle15(data []int) int {
	h := make(map[int][]int)
	var last int
	for i, n := range data {
		h[n] = append(h[n], i+1)
		last = n
	}
	for turn := len(data) + 1; ; turn++ {
		var n int
		if len(h[last]) != 1 {
			n = h[last][len(h[last])-1] - h[last][len(h[last])-2]
		}

		h[n] = append(h[n], turn)
		last = n

		if turn == 2020 {
			break
		}
	}
	return last
}

func Puzzle15() int {
	data := "0,12,6,13,20,1,17"
	return puzzle15(toNumbers(strings.Split(data, ",")))
}
