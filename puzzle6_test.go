package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle6(t *testing.T) {
	data := `abc

a
b
c

ab
ac

a
a
a
a

b`
	expect := 11
	result := puzzle6(strings.Split(data, "\n"))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle6() {
	fmt.Println(Puzzle6())

	// Output: 6735
}