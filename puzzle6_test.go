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

func ExampleSolutions_Puzzle6() {
	fmt.Println(solutions.Puzzle6())

	// Output: 6735
}

func TestPuzzle6b(t *testing.T) {
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
	expect := 6
	result := puzzle6b(strings.Split(data, "\n"))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleSolutions_Puzzle6b() {
	fmt.Println(solutions.Puzzle6b())

	// Output: 3221
}
