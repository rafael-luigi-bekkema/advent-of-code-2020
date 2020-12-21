package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestPuzzle21(t *testing.T) {
	data := strings.Split(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`, "\n")
	expect := 5
	result := puzzle21(data)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExamplePuzzle21() {
	fmt.Println(Puzzle21())

	// Output: 2374
}

func TestPuzzle21b(t *testing.T) {
	data := strings.Split(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`, "\n")
	expect := "mxmxvkd,sqjhc,fvjkl"
	result := puzzle21b(data)
	if result != expect {
		t.Fatalf("expected %s, got %s", expect, result)
	}
}

func ExamplePuzzle21b() {
	fmt.Println(Puzzle21b())

	// Output: fbtqkzc,jbbsjh,cpttmnv,ccrbr,tdmqcl,vnjxjg,nlph,mzqjxq
}
