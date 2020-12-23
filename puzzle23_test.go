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

func ExamplePuzzle23() {
	fmt.Println(Puzzle23())
	// Output: 62934785
}
