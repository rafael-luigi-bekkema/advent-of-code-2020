package main

import (
	"regexp"
	"strings"
)

func puzzle7(data []string) int {
	subject := "shiny gold"

	outerExpr := regexp.MustCompile(`^(\w+ \w+) bags contain (.+)\.$`)
	innerExpr := regexp.MustCompile(`^\d+ (\w+ \w+) bags?$`)

	bags := make(map[string][]string)
	for _, line := range data {
		matches := outerExpr.FindStringSubmatch(line)
		if matches == nil {
			panic("unexpected line: " + line)
		}
		outer := matches[1]
		inners := strings.Split(matches[2], ", ")
		for _, inner := range inners {
			if inner == "no other bags" {
				continue
			}
			smatches := innerExpr.FindStringSubmatch(inner)
			if smatches == nil {
				panic("unexpected subline: " + inner)
			}
			bags[smatches[1]] = append(bags[smatches[1]], outer)
		}
	}

	result := make(map[string]bool)
	var count func(bag string)
	count = func(bag string) {
		for _, outer := range bags[bag] {
			result[outer] = true
			count(outer)
		}
	}
	count(subject)

	total := len(result)
	return total
}

func Puzzle7() int {
	data := readFile("input/input07")
	return puzzle7(data)
}
