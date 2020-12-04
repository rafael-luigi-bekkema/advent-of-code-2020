package main

import (
	"fmt"
	"strings"
)

var p4creds = []string{
	"byr", // (Birth Year)
	"iyr", // (Issue Year)
	"eyr", // (Expiration Year)
	"hgt", // (Height)
	"hcl", // (Hair Color)
	"ecl", // (Eye Color)
	"pid", // (Passport ID)
	// "cid", // (Country ID)
}

func p4valid(data map[string]string) bool {
	for _, cred := range p4creds {
		if _, ok := data[cred]; !ok {
			return false
		}
	}
	return true
}

func puzzle4(input []string) int {
	var validCount int
	data := make(map[string]string)
	for _, line := range input {
		if line == "" {
			if p4valid(data) {
				validCount++
			}
			data = make(map[string]string)
		}
		for _, field := range strings.Fields(line) {
			var name, value string
			fmt.Sscanf(field, "%3s:%s", &name, &value)
			data[name] = value
		}
	}
	if p4valid(data) {
		validCount++
	}
	return validCount
}

func Puzzle4() int {
	input := readFile("input/input04")
	return puzzle4(input)
}
