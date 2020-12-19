package main

import (
	"strconv"
	"strings"
)

type p19data struct {
	data    string
	pos     int
	current byte
}

func (d *p19data) next() bool {
	if d.pos >= len(d.data) {
		return false
	}
	d.current = d.data[d.pos]
	d.pos++
	return true
}

type p19opts struct {
	opts  [][]int
	rules map[int]p19check
}

func (o *p19opts) Check(data *p19data) bool {
	for _, opt := range o.opts {
		matches := true
		ndata := *data
		for _, rule := range opt {
			if !o.rules[rule].Check(&ndata) {
				matches = false
				break
			}
		}
		if matches {
			data.pos = ndata.pos
			return true
		}
	}
	return false
}

type p19val struct {
	value byte
}

func (v *p19val) Check(data *p19data) bool {
	result := data.next() && data.current == v.value
	return result
}

type p19check interface {
	Check(data *p19data) bool
}

func puzzle19(data []string) int {
	m := make(map[int]p19check)
	for _, line := range data {
		if line == "" {
			break
		}

		parts := strings.Split(line, ": ")
		i, _ := strconv.Atoi(parts[0])
		if parts[1][0] == '"' {
			m[i] = &p19val{value: parts[1][1]}
			continue
		}
		opts := p19opts{rules: m}
		for _, part := range strings.Split(parts[1], " | ") {
			opts.opts = append(opts.opts, toNumbers(strings.Split(part, " ")))
		}
		m[i] = &opts

	}

	var count int
	for _, line := range data[len(m)+1:] {
		d := p19data{data: line}
		if m[0].Check(&d) && d.pos == len(line) {
			count++
		}
	}
	return count
}

func Puzzle19() int {
	data := readFile("input/input19")
	return puzzle19(data)
}
