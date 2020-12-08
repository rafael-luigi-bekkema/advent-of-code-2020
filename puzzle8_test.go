package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle8(t *testing.T) {
	data := strings.Split(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, "\n")
	expect := 5
	result := puzzle8(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle8() {
	fmt.Println(Puzzle8())

	// Output: 1179
}
