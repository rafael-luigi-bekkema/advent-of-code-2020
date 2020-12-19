package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle19(t *testing.T) {
	data := strings.Split(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`, "\n")
	expect := 2
	result := puzzle19(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle19() {
	fmt.Println(Puzzle19())

	// Output: 109
}
