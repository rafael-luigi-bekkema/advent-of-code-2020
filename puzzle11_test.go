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

func ExamplePuzzle11() {
	fmt.Println(Puzzle11())

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

func ExamplePuzzle11b() {
	fmt.Println(Puzzle11b())

	// Output: 1974
}
