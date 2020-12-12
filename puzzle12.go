package main

import (
	"fmt"
	"math"
)

func puzzle12(data []string) int {
	east := 0
	south := 1
	west := 2
	north := 3

	ferry := struct {
		facing int
		x, y   int
	}{east, 0, 0}

	for _, line := range data {
		var ins rune
		var n int
		fmt.Sscanf(line, "%c%d", &ins, &n)

		switch ins {
		case 'N':
			ferry.y -= n
		case 'S':
			ferry.y += n
		case 'E':
			ferry.x += n
		case 'W':
			ferry.x -= n
		case 'L':
			ferry.facing = (ferry.facing + ((360 - n) / 90)) % 4
		case 'R':
			ferry.facing = (ferry.facing + (n / 90)) % 4
		case 'F':
			switch ferry.facing {
			case north:
				ferry.y -= n
			case south:
				ferry.y += n
			case east:
				ferry.x += n
			case west:
				ferry.x -= n
			}
		}
	}

	return int(math.Abs(float64(ferry.x)) + math.Abs(float64(ferry.y)))
}

func Puzzle12() int {
	data := readFile("input/input12")
	return puzzle12(data)
}
