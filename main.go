package main

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/alecthomas/kingpin.v2"
)

type Solutions struct {
	debug bool
}

var solutions Solutions

func main() {
	day := kingpin.Arg("day", "Day").Required().Int()
	part := kingpin.Arg("part", "Part").Required().Enum("1", "2")
	debug := kingpin.Flag("debug", "Enable debug output").Short('d').Bool()
	kingpin.Parse()
	solutions.debug = *debug

	if len(os.Args) < 3 {
		fmt.Println("Give day and part (e.g. 23 1)")
		os.Exit(1)
	}

	var spart string
	if *part == "2" {
		spart = "b"
	}
	funcName := fmt.Sprintf("Puzzle%d%s", *day, spart)
	s := reflect.ValueOf(&solutions)
	m := s.MethodByName(funcName)
	if !m.IsValid() {
		fmt.Println("Solution not found")
		os.Exit(1)
	}
	values := m.Call(nil)
	fmt.Printf("Day %d part %v solution: ", *day, *part)
	for _, val := range values {
		fmt.Printf("%v", val)
	}
	fmt.Println()
}
