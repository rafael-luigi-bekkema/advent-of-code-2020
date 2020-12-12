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

func ExamplePuzzle12() {
	fmt.Println(Puzzle12())

	// Output: 1601
}
