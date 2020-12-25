package main

func puzzle25(cardKey, doorKey int) int {
	loops := 0
	for value := 1; value != cardKey; loops++ {
		value = (7 * value) % 20201227
	}

	value := 1
	for loop := 1; loop <= loops; loop++ {
		value = (doorKey * value) % 20201227
	}
	return value
}

func (s *Solutions) Puzzle25() int {
	data := toNumbers(readFile("input/input25"))
	return puzzle25(data[0], data[1])
}
