package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle16(t *testing.T) {
	data := strings.Split(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`, "\n")
	expect := 71
	result := puzzle16(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle16() {
	fmt.Println(Puzzle16())

	// Output: 18142
}
