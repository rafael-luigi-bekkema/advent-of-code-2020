package main

func Puzzle3() int {
	lines := readFile("input/input03")
	return puzzle3(lines)
}

func puzzle3(input []string) int {
	right := 3
	down := 1

	return calcSlope(input, right, down)
}

func calcSlope(input []string, right, down int) int {
	width := len(input[0])
	var posx, posy, count int
	for {
		if posy > len(input)-1 {
			break
		}

		line := input[posy]
		c := line[posx%width]
		if c == '#' {
			count++
		}

		posx += right
		posy += down
	}
	return count
}

func Puzzle3b() int {
	lines := readFile("input/input03")
	return puzzle3b(lines)
}

func puzzle3b(input []string) int {
	slopes := []struct{ right, down int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	var result int
	for _, slope := range slopes {
		val := calcSlope(input, slope.right, slope.down)
		if result == 0 {
			result = val
		} else {
			result *= val
		}
	}
	return result
}
