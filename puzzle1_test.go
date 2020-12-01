package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	testinput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	result := puzzle1(testinput)
	expect := 514579
	if result != expect {
		t.Fatalf("incorrect output: expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle1() {
	fmt.Println(Puzzle1())

	// Output: Puzzle1 result: 1019571
}
