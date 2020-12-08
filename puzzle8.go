package main

import "fmt"

func puzzle8(data []string) int {
	var acc, p int
	hist := make(map[int]bool)
	for {
		if hist[p] {
			break
		}
		hist[p] = true

		line := data[p]
		var ins string
		var n int
		fmt.Sscanf(line, "%s %d", &ins, &n)

		switch ins {
		case "nop":
			p++
		case "acc":
			acc += n
			p++
		case "jmp":
			p += n
		}
	}
	return acc
}

func Puzzle8() int {
	data := readFile("input/input08")
	return puzzle8(data)
}
