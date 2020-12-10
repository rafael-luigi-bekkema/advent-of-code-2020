package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle10(t *testing.T) {
	tt := []struct {
		data   string
		expect int
	}{
		{"16 10 15 5 1 11 7 19 6 12 4", 35},
		{"28 33 18 42 31 14 46 20 48 47 24 23 49 45 19 38 39 11 1 32 25 35 8 17 7 9 4 2 34 10 3", 220},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			data := toNumbers(strings.Split(tc.data, " "))
			result := puzzle10(data)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExamplePuzzle10() {
	fmt.Println(Puzzle10())

	// Output: 2470
}
