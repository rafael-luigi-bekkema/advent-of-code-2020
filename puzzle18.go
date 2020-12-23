package main

import (
	"fmt"
	"strings"
)

type p18op struct {
	value int
	op    byte
}

type tokenizer struct {
	data       string
	pos        int
	tok        byte
	precedence map[byte]int
}

func (t *tokenizer) nextToken() bool {
	if t.pos >= len(t.data) {
		return false
	}

	t.tok = t.data[t.pos]
	t.pos++
	return true
}

func (t *tokenizer) resolve(calcs []p18op) int {
	for prec := 0; prec <= 1; prec++ {
		for i := 0; i < len(calcs); i++ {
			op := calcs[i]
			if i > 0 && t.precedence[op.op] == prec {
				prevop := calcs[i-1]
				if op.op == '+' {
					calcs[i-1].value = prevop.value + op.value
				} else {
					calcs[i-1].value = prevop.value * op.value
				}
				calcs = append(calcs[:i], calcs[i+1:]...)
				i--
				continue
			}
		}
	}
	if len(calcs) != 1 {
		panic(fmt.Sprintf("expected 1 value, got %d", len(calcs)))
	}
	return calcs[0].value
}

func (t *tokenizer) parse() int {
	var op byte
	var calcs []p18op
	for t.nextToken() {
		var n1 int
		if t.tok == '(' {
			n1 = t.parse()
		} else {
			n1 = int(t.tok) - 48
		}
		calcs = append(calcs, p18op{n1, op})
		if !t.nextToken() || t.tok == ')' {
			break
		}
		op = t.tok
	}
	return t.resolve(calcs)
}

func p18calc(line string) int {
	t := tokenizer{data: strings.ReplaceAll(line, " ", "")}
	return t.parse()
}

func puzzle18(data []string) int {
	var sum int
	for _, line := range data {
		sum += p18calc(line)
	}
	return sum
}

func (s *Solutions) Puzzle18() int {
	data := readFile("input/input18")
	return puzzle18(data)
}

func p18calcb(expr string) int {
	t := tokenizer{data: strings.ReplaceAll(expr, " ", ""), precedence: map[byte]int{
		'+': 0,
		'*': 1,
	}}
	return t.parse()
}

func puzzle18b(data []string) int {
	var sum int
	for _, line := range data {
		sum += p18calcb(line)
	}
	return sum
}

func (s *Solutions) Puzzle18b() int {
	data := readFile("input/input18")
	return puzzle18b(data)
}
