package main

import (
	"fmt"
	"strings"
)

type p16field struct {
	name   string
	ranges []p16range
}

func (f *p16field) in(val int) bool {
	for _, r := range f.ranges {
		if r.in(val) {
			return true
		}
	}
	return false
}

type p16range struct {
	min, max int
}

func (r *p16range) in(val int) bool {
	return r.min <= val && val <= r.max
}

type p16fields []p16field

func (fs p16fields) in(val int) bool {
	for _, f := range fs {
		if f.in(val) {
			return true
		}
	}
	return false
}

func puzzle16(data []string) int {
	var fields p16fields
	var l int
	// Scan field ranges
	for {
		line := data[l]
		if line == "" {
			break
		}
		l++

		parts := strings.Split(line, ": ")
		sranges := strings.Split(parts[1], " or ")
		field := p16field{name: parts[0]}
		for _, sr := range sranges {
			var r p16range
			fmt.Sscanf(sr, "%d-%d", &r.min, &r.max)
			field.ranges = append(field.ranges, r)
		}
		fields = append(fields, field)
	}
	var error int
	for _, line := range data[l+5:] {
		for _, n := range toNumbers(strings.Split(line, ",")) {
			if !fields.in(n) {
				error += n
			}
		}
	}
	return error
}

func Puzzle16() int {
	data := readFile("input/input16")
	return puzzle16(data)
}
