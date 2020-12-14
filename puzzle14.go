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

func puzzle14b(data []string) int {
	var mask string
	getAddrs := func(addr int) []int {
		baddr := fmt.Sprintf("%036b", addr)
		rddr := []rune(baddr)
		var xses []int
		// Construct placeholder address and collect indeces of Xes
		for i, c := range mask {
			if c == '0' {
				continue
			}
			rddr[i] = c
			if c == 'X' {
				xses = append(xses, i)
			}
		}

		n := len(xses)
		// Calculate maximum value of bitstring of size n
		maxn, err := strconv.ParseInt(strings.Repeat("1", n), 2, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid binary string: %s", err))
		}
		var val int
		var addrs []int

		// Go over every possible value of the Xes
		for i := 0; i <= int(maxn); i++ {
			rddr := append([]rune{}, rddr...)
			bs := []rune(fmt.Sprintf(fmt.Sprintf("%%0%db\n", n), val))
			for j, xi := range xses {
				rddr[xi] = bs[j]
			}
			addr, _ := strconv.ParseInt(string(rddr), 2, 64)
			addrs = append(addrs, int(addr))
			val += 1
		}
		return addrs
	}

	mem := make(map[int]int)
	for _, line := range data {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
			continue
		}
		var addr, val int
		val, _ = strconv.Atoi(parts[1])
		fmt.Sscanf(parts[0], "mem[%d]", &addr)
		for _, raddr := range getAddrs(addr) {
			mem[raddr] = val
		}
	}

	var total int
	for _, val := range mem {
		total += val
	}
	return total
}

func Puzzle14b() int {
	data := readFile("input/input14")
	return puzzle14b(data)
}
