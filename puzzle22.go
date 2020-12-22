package main

import (
	"fmt"
	"sort"
	"strconv"
)

func p22deal(data []string) [2][]int {
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
	return decks
}

func puzzle22(data []string) int {
	decks := p22deal(data)

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

func intsEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, val := range s1 {
		if s2[i] != val {
			return false
		}
	}
	return true
}

func deckEqual(d1, d2 [2][]int) bool {
	return intsEqual(d1[0], d2[0]) && intsEqual(d1[1], d2[1])
}

func copyDecks(data [2][]int) [2][]int {
	cp := [2][]int{
		append([]int{}, data[0]...),
		append([]int{}, data[1]...),
	}
	return cp
}

func puzzle22b(data []string) int {
	var subgame func(decks *[2][]int, game int) int
	debug := true
	subgame = func(decks *[2][]int, game int) int {
		var round int
		var deckHistory [][2][]int
		for {
			round++
			if debug {
				fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
				fmt.Printf("Player 1's deck: %v\n", decks[0])
				fmt.Printf("Player 2's deck: %v\n", decks[1])
				fmt.Printf("Player 1 plays: %d\n", decks[0][0])
				fmt.Printf("Player 2 plays: %d\n", decks[1][0])
			}

			for _, h := range deckHistory {
				if deckEqual(*decks, h) {
					return 0
				}
			}

			cards := []int{
				decks[0][0],
				decks[1][0],
			}
			doSub := true
			for player, card := range cards {
				if len(decks[player])-1 < card {
					doSub = false
					break
				}
			}
			var winner int
			if doSub {
				if debug {
					fmt.Printf("Playing a sub-game to determine the winner...\n")
				}
				winner = subgame(
					&[2][]int{
						append([]int{}, decks[0][1:cards[0]+1]...),
						append([]int{}, decks[1][1:cards[1]+1]...),
					},
					game+1,
				)
			} else {
				if cards[1] > cards[0] {
					winner = 1
				}
			}

			if debug {
				fmt.Printf("Player %d wins round %d of game %d!\n", winner+1, round, game)
			}

			deckHistory = append(deckHistory, copyDecks(*decks))
			for player := range decks {
				decks[player] = decks[player][1:]
				if player == winner {
					if winner == 1 {
						cards[0], cards[1] = cards[1], cards[0]
					}
					decks[player] = append(decks[player], cards...)
					continue
				}
			}

			if len(decks[0]) == 0 || len(decks[1]) == 0 {
				break
			}
		}
		for player, deck := range decks {
			if len(deck) > 0 {
				return player
			}
		}
		panic("how did we get here?")
	}

	decks := p22deal(data)
	winner := subgame(&decks, 1)
	if debug {
		fmt.Printf("\n== Post-game results ==\n")
		fmt.Printf("Player 1's deck: %v\n", decks[0])
		fmt.Printf("Player 2's deck: %v\n", decks[1])
	}
	var score int
	for i, card := range decks[winner] {
		m := len(decks[winner]) - i
		score += m * card
	}
	return score
}

func Puzzle22b() int {
	data := readFile("input/input22")
	return puzzle22b(data)
}
