package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle18Calc(t *testing.T) {
	tt := []struct {
		expr   string
		expect int
	}{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3) ", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := p18calc(tc.expr)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func TestPuzzle18(t *testing.T) {
	data := strings.Split(`2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, "\n")

	expect := 26335
	result := puzzle18(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle18() {
	fmt.Println(Puzzle18())

	// Output: 9535936849815
}
