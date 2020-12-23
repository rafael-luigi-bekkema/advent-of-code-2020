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

func (s *Solutions) Puzzle13() int {
	data := readFile("input/input13")
	return puzzle13(data)
}

func puzzle13b(data string) int {
	var buses []int
	for _, b := range strings.Split(data, ",") {
		bus, _ := strconv.Atoi(b)
		buses = append(buses, bus)
	}

	var t int
	matches := map[int]int{
		0: buses[0],
	}
	var addPoss int
	add := buses[0]
	for {
		t += add

		found := true
		for i, bus := range buses {
			if bus == 0 || i == 0 {
				continue
			}
			if (t+i)%bus != 0 {
				found = false
				break
			} else {
				// This is the important part to make it
				// fast enough for the full input.
				// When a bus i > 0 lines up the first time, store t.
				// When it lines up the second time, make the
				// skip we are doing t-prevt.
				if prevt, ok := matches[i]; !ok {
					matches[i] = t
				} else if addPoss < i {
					add = t - prevt
					addPoss = i
					matches[i] = t
				}
			}
		}

		if found {
			break
		}
	}

	return t
}

func (s *Solutions) Puzzle13b() int {
	data := readFile("input/input13")
	return puzzle13b(data[1])
}
