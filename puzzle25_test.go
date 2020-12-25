package main

import (
	"fmt"
	"testing"
)

func TestPuzzle25(t *testing.T) {
	cardKey := 5764801
	doorKey := 17807724
	expect := 14897079
	result := puzzle25(cardKey, doorKey)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle25() {
	fmt.Println(solutions.Puzzle25())

	// Output: 18293391
}
