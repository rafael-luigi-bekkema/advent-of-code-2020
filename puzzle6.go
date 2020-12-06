package main

func puzzle6(data []string) int {
	var total int

	group := make(map[rune]bool)
	for _, line := range data {
		if line == "" {
			total += len(group)
			group = make(map[rune]bool)
		}
		for _, char := range line {
			group[char] = true
		}
	}
	total += len(group)

	return total
}

func Puzzle6() int {
	data := readFile("input/input06")
	return puzzle6(data)
}

func puzzle6b(data []string) int {
	var total int

	group := make(map[rune]int)
	var groupSize int
	for i := 0; i < len(data); i++ {
		line := data[i]
		for _, char := range line {
			group[char]++
		}
		groupSize++

		if i == len(data)-1 || data[i+1] == "" {
			for _, count := range group {
				if count == groupSize {
					total++
				}
			}
			groupSize = 0
			group = make(map[rune]int)
			i++
		}
	}

	return total
}

func Puzzle6b() int {
	data := readFile("input/input06")
	return puzzle6b(data)
}
