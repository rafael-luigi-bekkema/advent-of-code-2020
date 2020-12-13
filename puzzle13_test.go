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

func TestPuzzle13b(t *testing.T) {
	tt := []struct {
		data   string
		expect int
	}{
		{"7,13,x,x,59,x,31,19", 1068781},
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	// t * 7 == t * 18 + 1

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := puzzle13b(tc.data)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExamplePuzzle13b() {
	fmt.Println(Puzzle13b())

	// Output: 210612924879242
}
