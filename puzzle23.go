package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func puzzle23(data string, moves int) int {
	debug := false
	circle := ring.New(len(data))
	var maxval int
	labels := make(map[int]*ring.Ring)
	for _, c := range data {
		val := int(c - '0')
		circle.Value = val
		labels[val] = circle
		circle = circle.Next()
		if val > maxval {
			maxval = val
		}
	}

	for move := 1; move <= moves; move++ {
		if debug {
			fmt.Printf("-- move %d --\n", move)
			fmt.Printf("cups: ")
			circle.Do(func(val interface{}) {
				if val == circle.Value {
					fmt.Printf("(%d) ", val)
				} else {
					fmt.Printf("%d ", val)
				}
			})
			fmt.Println()
		}
		rem := circle.Unlink(3)
		dest := circle.Value.(int)
		for {
			dest--
			if dest < 1 {
				dest = maxval
			}
			if rem.Value.(int) != dest && rem.Next().Value.(int) != dest && rem.Move(2).Value.(int) != dest {
				break
			}
		}
		if debug {
			fmt.Printf("pick up: ")
			rem.Do(func(val interface{}) {
				fmt.Printf("%d ", val)
			})
			fmt.Printf("\ndestination: %d\n", dest)
		}
		labels[dest].Link(rem)
		circle = circle.Next()
	}

	if debug {
		fmt.Printf("-- final --\n")
		fmt.Printf("cups: ")
		circle.Do(func(val interface{}) {
			if val == circle.Value {
				fmt.Printf("(%d) ", val)
			} else {
				fmt.Printf("%d ", val)
			}
		})
		fmt.Println()
	}

	vals := make([]byte, 0, len(data)-1)
	labels[1].Next().Do(func(val interface{}) {
		if val.(int) == 1 {
			return
		}
		vals = append(vals, byte(val.(int))+'0')
	})
	result, _ := strconv.Atoi(string(vals))
	return result
}

func (s *Solutions) Puzzle23() int {
	return puzzle23("198753462", 100)
}

func puzzle23b(data string, moves int) int {
	debug := false
	circle := ring.New(1000_000)
	var maxval int
	labels := make(map[int]*ring.Ring)
	for _, c := range data {
		val := int(c - '0')
		circle.Value = val
		labels[val] = circle
		circle = circle.Next()
		if val > maxval {
			maxval = val
		}
	}
	for i := maxval + 1; i <= 1000_000; i++ {
		circle.Value = i
		labels[i] = circle
		circle = circle.Next()
	}
	maxval = 1000_000

	for move := 1; move <= moves; move++ {
		if debug {
			fmt.Printf("-- move %d --\n", move)
			fmt.Printf("cups: ")
			circle.Do(func(val interface{}) {
				if val == circle.Value {
					fmt.Printf("(%d) ", val)
				} else {
					fmt.Printf("%d ", val)
				}
			})
			fmt.Println()
		}
		rem := circle.Unlink(3)
		dest := circle.Value.(int)
		for {
			dest--
			if dest < 1 {
				dest = maxval
			}
			if rem.Value.(int) != dest && rem.Next().Value.(int) != dest && rem.Move(2).Value.(int) != dest {
				break
			}
		}
		if debug {
			fmt.Printf("pick up: ")
			rem.Do(func(val interface{}) {
				fmt.Printf("%d ", val)
			})
			fmt.Printf("\ndestination: %d\n", dest)
		}
		labels[dest].Link(rem)
		circle = circle.Next()
	}

	if debug {
		fmt.Printf("-- final --\n")
		fmt.Printf("cups: ")
		circle.Do(func(val interface{}) {
			if val == circle.Value {
				fmt.Printf("(%d) ", val)
			} else {
				fmt.Printf("%d ", val)
			}
		})
		fmt.Println()
	}

	result := labels[1].Next().Value.(int) * labels[1].Move(2).Value.(int)
	return result
}

func (s *Solutions) Puzzle23b() int {
	return puzzle23b("198753462", 10_000_000)
}
