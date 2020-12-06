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
