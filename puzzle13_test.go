package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle13(t *testing.T) {
	data := strings.Split(`939
7,13,x,x,59,x,31,19`, "\n")
	expect := 295
	result := puzzle13(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle13() {
	fmt.Println(Puzzle13())

	// Output: 259
}
