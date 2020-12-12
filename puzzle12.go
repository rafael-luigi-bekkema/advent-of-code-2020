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

func puzzle12b(data []string) int {
	turn := func(east, south int, deg int) (int, int) {
		// Convert turn angle to radians
		rad := float64(deg) * math.Pi / 180

		// Calculate current angle and add turn angle
		at := math.Atan2(float64(east), float64(south)*-1) + rad

		// Calculate radius
		r := math.Sqrt(float64(math.Pow(math.Abs(float64(east)), 2) + math.Pow(math.Abs(float64(south)), 2)))

		// Calculate new coordinates
		nx := r * math.Sin(at)
		ny := r * math.Cos(at) * -1

		return int(math.Round(nx)), int(math.Round(ny))
	}

	ferry := struct {
		x, y       int
		wayx, wayy int
	}{0, 0, 10, -1}

	for _, line := range data {
		var ins rune
		var n int
		fmt.Sscanf(line, "%c%d", &ins, &n)

		switch ins {
		case 'N':
			ferry.wayy -= n
		case 'S':
			ferry.wayy += n
		case 'E':
			ferry.wayx += n
		case 'W':
			ferry.wayx -= n
		case 'R':
			ferry.wayx, ferry.wayy = turn(ferry.wayx, ferry.wayy, n)
		case 'L':
			ferry.wayx, ferry.wayy = turn(ferry.wayx, ferry.wayy, 360-n)
		case 'F':
			ferry.x += (ferry.wayx * n)
			ferry.y += (ferry.wayy * n)
		}
	}

	return int(math.Abs(float64(ferry.x)) + math.Abs(float64(ferry.y)))
}

func Puzzle12b() int {
	data := readFile("input/input12")
	return puzzle12b(data)
}
