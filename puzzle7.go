package main

import (
	"regexp"
	"strconv"
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

type bag struct {
	count int
	name  string
}

func puzzle7b(data []string) int {
	subject := "shiny gold"

	outerExpr := regexp.MustCompile(`^(\w+ \w+) bags contain (.+)\.$`)
	innerExpr := regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)

	bags := make(map[string][]bag)
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
			count, _ := strconv.Atoi(smatches[1])
			bags[outer] = append(bags[outer], bag{count, smatches[2]})
		}
	}

	var count func(bag string) int
	count = func(bag string) int {
		var c int
		for _, b := range bags[bag] {
			c += b.count + b.count*count(b.name)
		}
		return c
	}

	result := count(subject)
	return result
}

func Puzzle7b() int {
	data := readFile("input/input07")
	return puzzle7b(data)
}
