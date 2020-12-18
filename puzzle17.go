package main

type coord struct {
	x, y, z, w int
}

func puzzle17(data []string) int {
	grid := make(map[coord]bool)
	for x, line := range data {
		for y, c := range line {
			grid[coord{x, y, 0, 0}] = c == '#'
		}
	}

	var isActive func(newGrid map[coord]bool, depth int, c coord) bool
	isActive = func(newGrid map[coord]bool, depth int, c coord) bool {
		var count int
		for i := 0; i < 27; i++ {
			if i == 13 {
				continue // 0, 0, 0
			}
			newc := c
			z := i/9 - 1
			x := i%3 - 1
			y := i%9/3 - 1
			newc.x += x
			newc.y += y
			newc.z += z

			if active, ok := grid[newc]; active {
				count++
			} else if !ok && depth == 0 {
				newGrid[newc] = isActive(newGrid, depth+1, newc)
			}
		}
		if grid[c] {
			return count == 2 || count == 3
		} else {
			return count == 3
		}
	}

	for pass := 0; pass < 6; pass++ {
		newGrid := make(map[coord]bool)
		for coord := range grid {
			newGrid[coord] = isActive(newGrid, 0, coord)
		}
		grid = newGrid
	}

	var total int
	for _, active := range grid {
		if active {
			total++
		}
	}
	return total
}

func Puzzle17() int {
	data := readFile("input/input17")
	return puzzle17(data)
}

func puzzle17b(data []string) int {
	grid := make(map[coord]bool)
	for x, line := range data {
		for y, c := range line {
			grid[coord{x, y, 0, 0}] = c == '#'
		}
	}

	var isActive func(newGrid map[coord]bool, depth int, c coord) bool
	isActive = func(newGrid map[coord]bool, depth int, c coord) bool {
		var count int
		for i := 0; i < 81; i++ {
			if i == 40 {
				continue // 0, 0, 0, 0
			}

			newc := c
			z := i/27 - 1
			x := i%3 - 1
			y := i%9/3 - 1
			w := i%27/9 - 1

			newc.x += x
			newc.y += y
			newc.z += z
			newc.w += w

			if active, ok := grid[newc]; active {
				count++
			} else if !ok && depth == 0 {
				newGrid[newc] = isActive(newGrid, depth+1, newc)
			}
		}
		if grid[c] {
			return count == 2 || count == 3
		} else {
			return count == 3
		}
	}

	for pass := 0; pass < 6; pass++ {
		newGrid := make(map[coord]bool)
		for coord := range grid {
			newGrid[coord] = isActive(newGrid, 0, coord)
		}
		grid = newGrid
	}

	var total int
	for _, active := range grid {
		if active {
			total++
		}
	}
	return total
}

func Puzzle17b() int {
	data := readFile("input/input17")
	return puzzle17b(data)
}
