package main

import (
	"fmt"
	"math"
)

func puzzle5(data []string) int {
	var maxSeatID int
	for _, seat := range data {
		from := 0.0
		to := 127.0
		for i := 0; i < 7; i++ {
			switch c := seat[i]; c {
			case 'F':
				to -= math.Ceil((to - from) / 2)
			case 'B':
				from += math.Ceil((to - from) / 2)
			default:
				panic(fmt.Sprintf("unexpected row char %c", c))
			}
		}
		if from != to {
			panic(fmt.Sprintf("more than 1 row left: %f-%f", from, to))
		}
		row := int(from)

		from = 0
		to = 7
		for i := 7; i < 10; i++ {
			switch c := seat[i]; c {
			case 'L':
				to -= math.Ceil((to - from) / 2)
			case 'R':
				from += math.Ceil((to - from) / 2)
			default:
				panic(fmt.Sprintf("unexpected col char %c", c))
			}
		}
		if from != to {
			panic(fmt.Sprintf("more than 1 column left: %f-%f", from, to))
		}

		col := int(from)

		seatID := row*8 + col
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

func Puzzle5() int {
	data := readFile("input/input05")
	return puzzle5(data)
}

func puzzle5b(data []string) int {
	seatChart := make([]bool, 128*8)
	for _, seat := range data {
		from := 0.0
		to := 127.0
		for i := 0; i < 7; i++ {
			switch c := seat[i]; c {
			case 'F':
				to -= math.Ceil((to - from) / 2)
			case 'B':
				from += math.Ceil((to - from) / 2)
			default:
				panic(fmt.Sprintf("unexpected row char %c", c))
			}
		}
		if from != to {
			panic(fmt.Sprintf("more than 1 row left: %f-%f", from, to))
		}
		row := int(from)

		from = 0
		to = 7
		for i := 7; i < 10; i++ {
			switch c := seat[i]; c {
			case 'L':
				to -= math.Ceil((to - from) / 2)
			case 'R':
				from += math.Ceil((to - from) / 2)
			default:
				panic(fmt.Sprintf("unexpected col char %c", c))
			}
		}
		if from != to {
			panic(fmt.Sprintf("more than 1 column left: %f-%f", from, to))
		}

		col := int(from)

		seatID := row*8 + col
		seatChart[seatID] = true
	}
	var results []int
	for seatID, occ := range seatChart {
		if !occ {
			if seatID > 0 && seatChart[seatID-1] && seatChart[seatID+1] {
				results = append(results, seatID)
			}
		}
	}
	if len(results) != 1 {
		panic(fmt.Sprintf("expected 1 results, got %d", len(results)))
	}

	return results[0]
}

func Puzzle5b() int {
	data := readFile("input/input05")
	return puzzle5b(data)
}
