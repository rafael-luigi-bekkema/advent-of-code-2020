package main

import (
	"fmt"
	"strconv"
	"strings"
)

func puzzle14(data []string) int {
	var mask string
	applyMask := func(val string) string {
		rval := []rune(val)
		for i, c := range mask {
			if c != 'X' {
				rval[i] = c
			}
		}
		return string(rval)
	}

	mem := make(map[int]string)
	for _, line := range data {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
			continue
		}
		var addr, val int
		fmt.Sscanf(parts[0], "mem[%d]", &addr)
		val, _ = strconv.Atoi(parts[1])
		bval := fmt.Sprintf("%036b", val)
		mbval := applyMask(bval)
		mem[addr] = mbval
	}

	var total int
	for _, val := range mem {
		n, _ := strconv.ParseInt(val, 2, 64)
		total += int(n)
	}
	return total
}

func Puzzle14() int {
	data := readFile("input/input14")
	return puzzle14(data)
}
