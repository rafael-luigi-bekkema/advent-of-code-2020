package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("could not read file: %s", err)
	}
	strs := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	return strs
}

func toNumbers(data []string) []int {
	nmbrs := make([]int, len(data))
	for i, line := range data {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("not a number: %s", err))
		}
		nmbrs[i] = n
	}
	return nmbrs
}

func min(data []int) int {
	var m int
	for i, n := range data {
		if i == 0 || n < m {
			m = n
		}
	}
	return m
}

func max(data []int) int {
	var m int
	for i, n := range data {
		if i == 0 || n > m {
			m = n
		}
	}
	return m
}
