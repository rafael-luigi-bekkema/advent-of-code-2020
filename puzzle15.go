package main

import (
	"fmt"
	"strings"
)

type p15item struct {
	prev, last int
}

func puzzle15(data []int, endTurn int) int {
	h := make(map[int]p15item)
	var last int
	for i, n := range data {
		h[n] = p15item{h[n].last, i + 1}
		last = n
	}
	// fmt.Println(data)
	for turn := len(data) + 1; ; turn++ {
		var n int
		// fmt.Printf("turn %d: ", turn)
		if h[last].prev != 0 {
			n = h[last].last - h[last].prev
			// fmt.Printf("%d old (%d-%d)\n", n, h[last][len(h[last])-1], h[last][len(h[last])-2])
		} else {
			// fmt.Printf("%d new\n", n)
		}

		fmt.Printf("%d: %d\n", turn, n)

		// if turn%1000_000 == 0 {
		// 	fmt.Printf("%.1f %%\n", float64(turn)/float64(endTurn)*100)
		// }

		h[n] = p15item{h[n].last, turn}
		last = n

		if turn == endTurn {
			break
		}
	}
	return last
}

func Puzzle15() int {
	data := "0,12,6,13,20,1,17"
	return puzzle15(toNumbers(strings.Split(data, ",")), 2020)
}

func Puzzle15b() int {
	data := "0,12,6,13,20,1,17"
	return puzzle15(toNumbers(strings.Split(data, ",")), 30_000_000)
}
