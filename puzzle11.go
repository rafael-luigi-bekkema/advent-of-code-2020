package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func puzzle11(data string) int {
	floor := '.'
	empty := 'L'
	occupied := '#'

	width := strings.Index(data, "\n")
	seats := []rune(strings.ReplaceAll(data, "\n", ""))

	adjCount := func(i int) int {
		var result int
		row := i / width
		col := i % width
		for crow := row - 1; crow <= row+1; crow++ {
			for ccol := col - 1; ccol <= col+1; ccol++ {
				if crow < 0 || crow >= len(seats)/width {
					continue
				}
				if ccol < 0 || ccol >= width {
					continue
				}
				j := crow*width + ccol
				if i == j {
					continue
				}
				if seats[j] == occupied {
					result++
				}
			}
		}
		return result
	}

	printMap := func(m []rune) {
		for i := 0; i < len(m)/width; i++ {
			fmt.Printf("%s\n", string(m[i*width:i*width+width]))
		}
		fmt.Println()
	}

	for {
		newS := make([]rune, len(seats))
		for i, s := range seats {
			if s == floor {
				newS[i] = floor
				continue
			}
			occCount := adjCount(i)
			if s == empty && occCount == 0 {
				newS[i] = occupied
				continue
			}
			if s == occupied && occCount >= 4 {
				newS[i] = empty
				continue
			}
			newS[i] = s
		}

		oldS := seats
		seats = newS
		if string(oldS) == string(newS) {
			break
		}
	}

	_ = printMap

	return strings.Count(string(seats), string(occupied))
}

func (s *Solutions) Puzzle11() int {
	data, err := ioutil.ReadFile("input/input11")
	if err != nil {
		log.Panicf("could not read file: %s", err)
	}
	return puzzle11(string(data))
}

func puzzle11b(data string) int {
	floor := '.'
	empty := 'L'
	occupied := '#'

	width := strings.Index(data, "\n")
	seats := []rune(strings.ReplaceAll(data, "\n", ""))

	find := func(i, x, y int) bool {
		row := i / width
		col := i % width
		for {
			row += y
			col += x
			if col < 0 || col >= width || row < 0 || row >= len(seats)/width {
				return false
			}
			s := seats[row*width+col]
			if s == floor {
				continue
			}
			return s == occupied
		}
	}

	adjCount := func(i int) int {
		var result int
		for j := 0; j < 9; j++ {
			if j == 4 {
				continue
			}
			if find(i, j%3-1, j/3-1) {
				result++
			}
		}
		return result
	}

	printMap := func(m []rune) {
		for i := 0; i < len(m)/width; i++ {
			fmt.Printf("%s\n", string(m[i*width:i*width+width]))
		}
		fmt.Println()
	}

	for {
		newS := make([]rune, len(seats))
		for i, s := range seats {
			if s == floor {
				newS[i] = floor
				continue
			}
			occCount := adjCount(i)
			if s == empty && occCount == 0 {
				newS[i] = occupied
				continue
			}
			if s == occupied && occCount >= 5 {
				newS[i] = empty
				continue
			}
			newS[i] = s
		}

		oldS := seats
		seats = newS
		if string(oldS) == string(newS) {
			break
		}
	}

	_ = printMap

	return strings.Count(string(seats), string(occupied))
}

func (s *Solutions) Puzzle11b() int {
	data, err := ioutil.ReadFile("input/input11")
	if err != nil {
		log.Panicf("could not read file: %s", err)
	}
	return puzzle11b(string(data))
}
