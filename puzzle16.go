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

func p16scanFields(data []string) p16fields {
	var fields p16fields
	for _, line := range data {
		if line == "" {
			break
		}

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
	return fields
}

func puzzle16(data []string) int {
	fields := p16scanFields(data)

	var error int
	for _, line := range data[len(fields)+5:] {
		for _, n := range toNumbers(strings.Split(line, ",")) {
			if !fields.in(n) {
				error += n
			}
		}
	}
	return error
}

func puzzle16b(data []string) int {
	fields := p16scanFields(data)

	foptions := make(map[int]map[string]struct{})
	var myTicket []int
	for _, line := range data[len(fields)+2:] {
		if line == "" || line == "nearby tickets:" {
			continue
		}
		nmbrs := toNumbers(strings.Split(line, ","))
		if myTicket == nil {
			myTicket = nmbrs
		}
		for i, n := range nmbrs {
			matches := make(map[string]struct{})
			for _, f := range fields {
				if f.in(n) {
					matches[f.name] = struct{}{}
				}
			}
			if len(matches) == 0 {
				break // invalid ticket
			}
			if len(foptions[i]) == 0 {
				foptions[i] = matches
				continue
			}
			for key := range foptions[i] {
				if _, ok := matches[key]; !ok {
					delete(foptions[i], key)
				}
			}
		}
	}

	for {
		var changed bool
		for i, opts := range foptions {
			if len(opts) == 1 {
				continue
			}
			for name := range opts {
				unique := true
				for si, sopts := range foptions {
					if si == i {
						continue
					}
					if _, ok := sopts[name]; ok {
						unique = false
						break
					}
				}
				if !unique {
					continue
				}
				// if unique, remove other items
				foptions[i] = map[string]struct{}{
					name: {},
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	var total int
	for i, fs := range foptions {
		if len(fs) != 1 {
			panic(fmt.Sprintf("expected only 1, got %d", len(fs)))
		}
		for name := range fs {
			if strings.HasPrefix(name, "departure") {
				if total == 0 {
					total = myTicket[i]
				} else {
					total *= myTicket[i]
				}
			}
		}
	}
	return total
}

func (s *Solutions) Puzzle16() int {
	data := readFile("input/input16")
	return puzzle16(data)
}

func (s *Solutions) Puzzle16b() int {
	data := readFile("input/input16")
	return puzzle16b(data)
}
