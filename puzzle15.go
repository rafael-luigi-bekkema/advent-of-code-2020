package main

import (
	"strings"
)

func puzzle15(data []int, endTurn int) int {
	h := make([]int, endTurn)
	for i, n := range data {
		h[n] = i + 1
	}

	var prev int
	for turn := len(data) + 1; turn < endTurn; turn++ {
		i := h[prev]
		h[prev] = turn

		if i != 0 {
			prev = turn - i
		} else {
			prev = 0
		}
	}
	return prev
}

func Puzzle15() int {
	data := "0,12,6,13,20,1,17"
	return puzzle15(toNumbers(strings.Split(data, ",")), 2020)
}

func Puzzle15b() int {
	data := "0,12,6,13,20,1,17"
	return puzzle15(toNumbers(strings.Split(data, ",")), 30_000_000)
}
