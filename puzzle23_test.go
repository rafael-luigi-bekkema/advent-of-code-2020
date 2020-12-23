package main

import (
	"fmt"
	"testing"
)

func TestPuzzle23(t *testing.T) {
	data := `389125467`
	expect := 67384529
	result := puzzle23(data, 100)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestPuzzle23_2(t *testing.T) {
	data := `389125467`
	expect := 92658374
	result := puzzle23(data, 10)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle23() {
	fmt.Println(solutions.Puzzle23())
	// Output: 62934785
}

func TestPuzzle23b(t *testing.T) {
	data := `389125467`
	expect := 149245887792
	result := puzzle23b(data, 10_000_000)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle23b() {
	fmt.Println(solutions.Puzzle23b())
	// Output: 693659135400
}
