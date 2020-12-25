package main

func findLoops(key int) int {
	value := 1
	loop := 0
	subject := 7
	for {
		value = (subject * value) % 20201227
		loop++
		if value == key {
			break
		}
		if loop >= 100_000_000 {
			panic("too many loops")
		}
	}
	return loop
}
func doLoops(subject, loops int) int {
	value := 1
	for loop := 1; loop <= loops; loop++ {
		value = (subject * value) % 20201227
	}
	return value
}

func puzzle25(cardKey, doorKey int) int {
	cardLoops := findLoops(cardKey)
	doorLoops := findLoops(doorKey)

	encKey1 := doLoops(cardKey, doorLoops)
	encKey2 := doLoops(doorKey, cardLoops)

	if encKey1 != encKey2 {
		panic("keys are not equal")
	}

	return encKey1
}

func (s *Solutions) Puzzle25() int {
	data := toNumbers(readFile("input/input25"))
	return puzzle25(data[0], data[1])
}
