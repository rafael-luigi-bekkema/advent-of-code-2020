package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Puzzle1() string {
	data, _ := ioutil.ReadFile("input/input01")
	strs := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	input := make([]int, len(strs))
	for i, s := range strs {
		input[i], _ = strconv.Atoi(s)
	}
	return fmt.Sprintf("Puzzle1 result: %d", puzzle1(input))
}

func puzzle1(input []int) int {
	sum := 2020
	for _, i := range input {
		for _, j := range input {
			if i+j == sum {
				return i * j
			}
		}
	}
	panic("puzzle1 result not found")
}

func Puzzle1b() string {
	data, _ := ioutil.ReadFile("input/input01")
	strs := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	input := make([]int, len(strs))
	for i, s := range strs {
		input[i], _ = strconv.Atoi(s)
	}
	return fmt.Sprintf("Puzzle1b result: %d", puzzle1b(input))
}

func puzzle1b(input []int) int {
	sum := 2020
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if i+j+k == sum {
					return i * j * k
				}
			}
		}
	}
	panic("puzzle1 result not found")
}
