package main

import (
	"sort"
	"strconv"
)

func puzzle22(data []string) int {
	var decks [2][]int
	var player int
	for _, line := range data {
		if len(line) > 2 {
			continue
		}
		if line == "" {
			player++
			continue
		}
		c, _ := strconv.Atoi(line)
		decks[player] = append(decks[player], c)
	}

	for {
		cards := []int{
			decks[0][0],
			decks[1][0],
		}
		winner := 0
		if cards[1] > cards[0] {
			winner = 1
		}
		for player := range decks {
			decks[player] = decks[player][1:]
			if player == winner {
				sort.Slice(cards, func(i, j int) bool {
					return cards[i] > cards[j]
				})
				decks[player] = append(decks[player], cards...)
				continue
			}
		}
		if len(decks[0]) == 0 || len(decks[1]) == 0 {
			break
		}
	}
	var winner int
	for player, deck := range decks {
		if len(deck) > 0 {
			winner = player
			break
		}
	}

	var score int
	for i, card := range decks[winner] {
		m := len(decks[winner]) - i
		score += m * card
	}
	return score
}

func Puzzle22() int {
	data := readFile("input/input22")
	return puzzle22(data)
}
