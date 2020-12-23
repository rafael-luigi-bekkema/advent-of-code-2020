package main

import (
	"fmt"
	"testing"
)

func TestPuzzle11(t *testing.T) {
	data := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	expect := 37
	result := puzzle11(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle11() {
	fmt.Println(solutions.Puzzle11())

	// Output: 2164
}

func TestPuzzle11b(t *testing.T) {
	data := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	expect := 26
	result := puzzle11b(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle11b() {
	fmt.Println(solutions.Puzzle11b())

	// Output: 1974
}
