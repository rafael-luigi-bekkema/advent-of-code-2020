package main

import "sort"

func puzzle10(data []int) int {
	adapters := make(map[int]struct{}, len(data))
	for _, j := range data {
		adapters[j] = struct{}{}
	}

	result := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}

	j := 0
Outer:
	for {
		for i := 1; i <= 3; i++ {
			if _, ok := adapters[i+j]; !ok {
				continue
			}
			result[i]++
			j += i
			continue Outer
		}
		break
	}
	return result[1] * result[3]
}

func Puzzle10() int {
	data := toNumbers(readFile("input/input10"))
	return puzzle10(data)
}

func puzzle10b(data []int) int {
	// Add 0 and max+3 to the list
	data = append([]int{0}, data...)
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)

	// Create a map of jolts for fast lookup
	jmap := make(map[int]struct{}, len(data))
	for _, n := range data {
		jmap[n] = struct{}{}
	}

	//
	var count func(jmap map[int]struct{}, j, maxj int) int
	count = func(jmap map[int]struct{}, j, maxj int) int {
		var total int
		for i := 1; i <= 3; i++ {
			if maxj == j+i {
				total++
			}
			if _, ok := jmap[j+i]; !ok {
				continue
			}
			total += count(jmap, j+i, maxj)
		}
		return total
	}

	var begin int
	var perms []int
	for i := 0; i < len(data); i++ {
		n := data[i]
		if i == 0 || data[i-1] == n-3 {
			if begin != 0 {
				perms = append(perms, count(toIntMap(data[begin:i]), data[begin-1], n))
				begin = 0
			}
			continue
		}
		if begin == 0 {
			begin = i
		}
	}

	var result int
	for i, perm := range perms {
		if i == 0 {
			result = perm
			continue
		}
		result *= perm
	}

	return result
}

func Puzzle10b() int {
	data := toNumbers(readFile("input/input10"))
	return puzzle10b(data)
}
