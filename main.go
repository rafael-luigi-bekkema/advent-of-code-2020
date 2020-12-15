package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Which day and part (e.g. 'go run . 5 1')?")
		os.Exit(1)
	}

	funcs := []func() int{
		Puzzle1,
		Puzzle1b,
		Puzzle2,
		Puzzle2b,
		Puzzle3,
		Puzzle3b,
		Puzzle4,
		Puzzle4b,
		Puzzle5,
		Puzzle5b,
		Puzzle6,
		Puzzle6b,
		Puzzle7,
		Puzzle7b,
		Puzzle8,
		Puzzle8b,
		Puzzle9,
		Puzzle9b,
		Puzzle10,
		Puzzle10b,
		Puzzle11,
		Puzzle11b,
		Puzzle12,
		Puzzle12b,
		Puzzle13,
		Puzzle13b,
		Puzzle14,
		Puzzle14b,
		Puzzle15,
		Puzzle15b,
	}

	day, _ := strconv.Atoi(os.Args[1])
	part, _ := strconv.Atoi(os.Args[2])
	i := (day-1)*2 + (part - 1)
	if i < 0 || i >= len(funcs) {
		fmt.Println("Which day and part (e.g. 'go run . 5 1')?")
		os.Exit(1)
	}

	fmt.Printf("Puzzle %d part %d: %d\n", day, part, funcs[i]())
}
