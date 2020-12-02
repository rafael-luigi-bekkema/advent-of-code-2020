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

func Puzzle2b() int {
	input := readFile("input/input02")
	return puzzle2b(input)
}

func puzzle2b(input []string) int {
	var valid int
	for _, line := range input {
		var pos1, pos2 int
		var pass string
		var letter byte
		fmt.Sscanf(line, "%d-%d %c: %s", &pos1, &pos2, &letter, &pass)

		if (pass[pos1-1] == letter) != (pass[pos2-1] == letter) {
			valid++
		}
	}
	return valid
}
