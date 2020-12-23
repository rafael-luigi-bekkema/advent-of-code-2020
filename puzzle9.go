package main

func puzzle9(data []int, preamble int) int {
	nmbrs := make([]int, 0, len(data))
	find := func(inp []int, target int) bool {
		for _, i := range inp {
			for _, j := range inp {
				if i == j {
					continue
				}
				if i+j == target {
					return true
				}
			}
		}
		return false
	}

	for i, n := range data {
		nmbrs = append(nmbrs, n)

		if i < preamble {
			continue
		}

		if !find(nmbrs[i-preamble:], n) {
			return n
		}
	}

	panic("result not found")
}

func (s *Solutions) Puzzle9() int {
	data := toNumbers(readFile("input/input09"))
	return puzzle9(data, 25)
}

func puzzle9b(data []int, preamble int) int {
	result := puzzle9(data, preamble)
	for i := range data {
		j := i
		total := data[i]
		for {
			j++
			total += data[j]
			if total == result {
				return min(data[i:j+1]) + max(data[i:j+1])
			}
			if total > result {
				break
			}
		}
	}
	return result
}

func (s *Solutions) Puzzle9b() int {
	data := toNumbers(readFile("input/input09"))
	return puzzle9b(data, 25)
}
