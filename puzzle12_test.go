package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle12(t *testing.T) {
	data := strings.Split(`F10
N3
F7
R90
F11`, "\n")
	expect := 25
	result := puzzle12(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle12() {
	fmt.Println(solutions.Puzzle12())

	// Output: 1601
}

func TestPuzzle12b(t *testing.T) {
	data := strings.Split(`F10
N3
F7
R90
F11`, "\n")
	expect := 286
	result := puzzle12b(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle12b() {
	fmt.Println(solutions.Puzzle12b())

	// Output: 13340
}
