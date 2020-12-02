package main

import (
	"io/ioutil"
	"log"
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
