package main

import (
	"fmt"
	"testing"
)

func TestPuzzle2(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	result := puzzle2(input)
	expect := 2
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle2() {
	fmt.Println(Puzzle2())

	// Output: 460
}
