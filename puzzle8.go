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

func p8run(data []string) (int, bool) {
	var acc, p int
	hist := make(map[int]bool)
	for {
		if hist[p] {
			return acc, false
		}
		hist[p] = true
		if p == len(data) {
			return acc, true
		}

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
}

func puzzle8b(data []string) int {
	for i := 0; i < len(data); i++ {
		line := data[i]
		var ins string
		var n int
		fmt.Sscanf(line, "%s %d", &ins, &n)
		if ins == "acc" {
			continue
		}

		ndata := append([]string{}, data...)
		var nins string
		if ins == "nop" {
			nins = "jmp"
		} else {
			nins = "nop"
		}
		ndata[i] = fmt.Sprintf("%s %+d", nins, n)

		acc, ok := p8run(ndata)
		if ok {
			return acc
		}
	}
	panic("no answer found")
}

func Puzzle8b() int {
	data := readFile("input/input08")
	return puzzle8b(data)
}
