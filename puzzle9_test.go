package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle9(t *testing.T) {
	data := toNumbers(strings.Split(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, "\n"))
	expect := 127
	result := puzzle9(data, 5)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle9() {
	fmt.Println(Puzzle9())

	// Output: 32321523
}

func TestPuzzle9b(t *testing.T) {
	data := toNumbers(strings.Split(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, "\n"))
	expect := 62
	result := puzzle9b(data, 5)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleTest9b() {
	fmt.Println(Puzzle9b())

	// Output: 4794981
}
