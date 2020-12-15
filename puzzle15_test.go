package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle15(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := puzzle15(toNumbers(strings.Split(tc.input, ",")), 2020)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExamplePuzzle15() {
	fmt.Println(Puzzle15())

	// Output: 620
}

func TestPuzzle15b(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := puzzle15(toNumbers(strings.Split(tc.input, ",")), 30_000_000)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExamplePuzzle15b() {
	fmt.Println(Puzzle15b())

	// Output: 110871
}
