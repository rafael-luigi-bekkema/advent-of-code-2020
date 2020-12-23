package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle5(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
		{"BBFFBBFRLL\nFFFBBBFRRR\nBFFFBBFRRR", 820},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			data := strings.Split(tc.input, "\n")
			result := puzzle5(data)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleSolutions_Puzzle5() {
	fmt.Println(solutions.Puzzle5())

	// Output: 835
}

func ExampleSolutions_Puzzle5b() {
	fmt.Println(solutions.Puzzle5b())

	// Output: 649
}
