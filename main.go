package main

import (
	"fmt"
	"os"
	"reflect"
)

type Solutions struct{}

var solutions Solutions

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Give day and part (e.g. 23 1)")
		os.Exit(1)
	}

	var part string
	if os.Args[2] == "2" {
		part = "b"
	} else if os.Args[2] != "1" {
		fmt.Println("Choose part 1 or 2")
		os.Exit(1)
	}
	funcName := fmt.Sprintf("Puzzle%s%s", os.Args[1], part)
	s := reflect.ValueOf(&solutions)
	m := s.MethodByName(funcName)
	if !m.IsValid() {
		fmt.Println("Solution not found")
		os.Exit(1)
	}
	values := m.Call(nil)
	fmt.Printf("Day %s part %s solution: ", os.Args[1], os.Args[2])
	for _, val := range values {
		fmt.Printf("%v", val)
	}
	fmt.Println()
}
