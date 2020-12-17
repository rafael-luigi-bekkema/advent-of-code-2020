package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle17(t *testing.T) {
	data := strings.Split(`.#.
..#
###`, "\n")
	expect := 112
	result := puzzle17(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle17() {
	fmt.Println(Puzzle17())

	// Output: 359
}
