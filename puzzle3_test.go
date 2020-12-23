package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle3(t *testing.T) {
	input := strings.Split(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`, "\n")
	expect := 7
	result := puzzle3(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle3() {
	fmt.Println(solutions.Puzzle3())

	// Output: 218
}

func TestPuzzle3b(t *testing.T) {
	input := strings.Split(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`, "\n")
	expect := 336
	result := puzzle3b(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle3b() {
	fmt.Println(solutions.Puzzle3b())

	// Output: 3847183340
}
