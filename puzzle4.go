package main

import (
	"fmt"
	"regexp"
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

func p4validate(name, value string) bool {
	switch name {
	case "byr":
		if "1920" <= value && value <= "2002" {
			return true
		}
	case "iyr":
		if "2010" <= value && value <= "2020" {
			return true
		}
	case "eyr":
		if "2020" <= value && value <= "2030" {
			return true
		}
	case "hgt":
		var height int
		var unit string
		fmt.Sscanf(value, "%d%s", &height, &unit)
		if unit == "cm" && 150 <= height && height <= 193 {
			return true
		}
		if unit == "in" && 59 <= height && height <= 76 {
			return true
		}
	case "hcl":
		match, _ := regexp.MatchString("^#[0-9a-f]{6}$", value)
		if match {
			return true
		}
	case "ecl":
		colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, clr := range colors {
			if clr == value {
				return true
			}
		}
	case "pid":
		match, _ := regexp.MatchString("^[0-9]{9}$", value)
		if match {
			return true
		}
	case "cid":
		return true
	}
	return false
}

func p4valid(data map[string]string, validation bool) bool {
	for _, cred := range p4creds {
		if _, ok := data[cred]; !ok {
			return false
		}
		if validation && !p4validate(cred, data[cred]) {
			return false
		}
	}
	return true
}

func puzzle4(input []string, validation bool) int {
	var validCount int
	data := make(map[string]string)
	for _, line := range input {
		if line == "" {
			if p4valid(data, validation) {
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
	if p4valid(data, validation) {
		validCount++
	}
	return validCount
}

func (s *Solutions) Puzzle4() int {
	input := readFile("input/input04")
	return puzzle4(input, false)
}

func (s *Solutions) Puzzle4b() int {
	input := readFile("input/input04")
	return puzzle4(input, true)
}
