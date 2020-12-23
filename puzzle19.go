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

func (o *p19opts) Check(data *p19data, level int) []int {
	var results []int
	for _, opt := range o.opts {
		matches := true
		ndata := *data
		poss := []int{data.pos}
		for _, rule := range opt {
			var subposs []int
			for _, pos := range poss {
				ndata.pos = pos
				r := o.rules[rule].Check(&ndata, level+1)
				subposs = append(subposs, r...)
			}
			poss = subposs
			if len(subposs) == 0 {
				matches = false
				break
			}
		}
		if matches {
			results = append(results, poss...)
		}
	}
	return results
}

type p19val struct {
	value byte
}

func (v *p19val) Check(data *p19data, level int) []int {
	result := data.next() && data.current == v.value
	if !result {
		return nil
	}
	return []int{data.pos}
}

type p19check interface {
	Check(data *p19data, level int) []int
}

func puzzle19(data []string, b bool) int {
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

	if b {
		m[8] = &p19opts{rules: m, opts: [][]int{{42}, {42, 8}}}
		m[11] = &p19opts{rules: m, opts: [][]int{{42, 31}, {42, 11, 31}}}
	}

	var count int
	for _, line := range data[len(m)+1:] {
		d := p19data{data: line}
		check := m[0].Check(&d, 0)
		if len(check) > 0 && check[0] == len(line) {
			count++
		}
	}
	return count
}

func (s *Solutions) Puzzle19() int {
	data := readFile("input/input19")
	return puzzle19(data, false)
}

func (s *Solutions) Puzzle19b() int {
	data := readFile("input/input19")
	return puzzle19(data, true)
}
