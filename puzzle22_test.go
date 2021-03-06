package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle22(t *testing.T) {
	data := strings.Split(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`, "\n")
	expect := 306
	result := puzzle22(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle22() {
	fmt.Println(solutions.Puzzle22())

	// Output: 33694
}

func TestPuzzle22b(t *testing.T) {
	data := strings.Split(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`, "\n")
	expect := 291
	result := puzzle22b(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle22b() {
	fmt.Println(solutions.Puzzle22b())

	// Output: 31835
}
