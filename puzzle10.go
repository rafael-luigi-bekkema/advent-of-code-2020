package main

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
