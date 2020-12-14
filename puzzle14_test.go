package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle14(t *testing.T) {
	data := strings.Split(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`, "\n")
	expect := 165
	result := puzzle14(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle14() {
	fmt.Println(Puzzle14())

	// Output: 14839536808842
}

func TestPuzzle14b(t *testing.T) {
	data := strings.Split(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, "\n")
	expect := 208
	result := puzzle14b(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle14b() {
	fmt.Println(Puzzle14b())

	// Output: 4215284199669
}
