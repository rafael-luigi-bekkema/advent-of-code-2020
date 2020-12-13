package main

import (
	"math"
	"strconv"
	"strings"
)

func puzzle13(data []string) int {
	depart, _ := strconv.Atoi(data[0])
	var min, minBus int
	for _, bus := range strings.Split(data[1], ",") {
		if bus == "x" {
			continue
		}
		b, _ := strconv.Atoi(bus)

		next := int(math.Ceil(float64(depart)/float64(b))) * b
		diff := next - depart
		if min == 0 || diff < min {
			min = diff
			minBus = b
		}
		// fmt.Printf("%d: %d\n", b, next)
	}

	return min * minBus
}

func Puzzle13() int {
	data := readFile("input/input13")
	return puzzle13(data)
}
