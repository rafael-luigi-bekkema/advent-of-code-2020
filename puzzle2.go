package main

import (
	"fmt"
	"strings"
)

func Puzzle2() int {
	input := readFile("input/input02")
	return puzzle2(input)
}

func puzzle2(input []string) int {
	var valid int
	for _, line := range input {
		var from, to int
		var letter, pass string
		fmt.Sscanf(line, "%d-%d %1s: %s", &from, &to, &letter, &pass)

		c := strings.Count(pass, letter)
		if from <= c && c <= to {
			valid++
		}
	}
	return valid
}
