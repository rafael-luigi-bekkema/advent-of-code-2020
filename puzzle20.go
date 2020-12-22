package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const p20width = 10

type p20tile struct {
	num  int
	data [100]bool
}

func (t *p20tile) rotate() {
	data := t.data
	for i, val := range data {
		row := i / p20width
		col := i % p20width
		newrow := col
		newcol := p20width - 1 - row
		t.data[newrow*p20width+newcol] = val
	}
}

func (t *p20tile) hflip() {
	data := t.data
	for i, val := range data {
		row := i / p20width
		col := i % p20width
		newcol := p20width - 1 - col
		newrow := row
		t.data[newrow*p20width+newcol] = val
	}
}

func (t *p20tile) print() {
	for i, val := range t.data {
		if i%p20width == 0 {
			fmt.Println()
		}
		var c byte
		if val {
			c = '#'
		} else {
			c = '.'
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func (t *p20tile) match(t2 *p20tile) (x int, y int) {
	var l1, l2 [10]bool
	copy(l1[:], t.data[0:10])
	copy(l2[:], t2.data[90:100])
	if l1 == l2 {
		return 0, -1
	}
	copy(l1[:], t.data[90:100])
	copy(l2[:], t2.data[0:10])
	if l1 == l2 {
		return 0, 1
	}

	for row := 0; row < 10; row++ {
		l1[row] = t.data[row*p20width]
		l2[row] = t2.data[row*p20width+9]
	}
	if l1 == l2 {
		return -1, 0
	}
	for row := 0; row < 10; row++ {
		l1[row] = t.data[row*p20width+9]
		l2[row] = t2.data[row*p20width]
	}
	if l1 == l2 {
		return 1, 0
	}
	return 0, 0
}

type p20pos struct {
	x, y int
	tile *p20tile
}

func puzzle20(data string) int {
	tileDs := strings.Split(data, "\n\n")
	tiles := make([]p20tile, len(tileDs))
	for i, tileD := range tileDs {
		lines := strings.Split(tileD, "\n")
		tile := &tiles[i]
		fmt.Sscanf(lines[0], "Tile %d:", &tile.num)
		for row, line := range lines[1:] {
			for col, c := range line {
				tile.data[row*p20width+col] = c == '#'
			}
		}
	}

	first := tiles[0]
	board := []p20pos{
		{x: 0, y: 0, tile: &first},
	}
	onBoard := map[int]bool{
		first.num: true,
	}
	tiles = tiles[1:]
	var minx, maxx, miny, maxy int

	for i := 0; i < len(board); i++ {
		tile := board[i]
		count := 0
		for j := 0; j < len(tiles); j++ {
			dTile := tiles[j]
			if onBoard[dTile.num] {
				continue
			}
		Outer:
			for fh := 0; fh < 2; fh++ {
				for r := 0; r < 4; r++ {
					x, y := tile.tile.match(&dTile)
					if x != 0 || y != 0 {
						var newpos p20pos
						newpos.x = tile.x + x
						newpos.y = tile.y + y
						newpos.tile = &dTile
						onBoard[dTile.num] = true
						board = append(board, newpos)
						tiles = append(tiles[:j], tiles[j+1:]...)
						j--
						count++
						if newpos.x < minx {
							minx = newpos.x
						}
						if newpos.x > maxx {
							maxx = newpos.x
						}
						if newpos.y < miny {
							miny = newpos.y
						}
						if newpos.y > maxy {
							maxy = newpos.y
						}
						break Outer
					}
					dTile.rotate()
				}
				dTile.hflip()
			}
		}
	}

	fboard := make(map[[2]int]int, len(board))
	for _, tile := range board {
		fboard[[2]int{tile.x, tile.y}] = tile.tile.num
	}
	nums := []int{
		fboard[[2]int{minx, miny}],
		fboard[[2]int{minx, maxy}],
		fboard[[2]int{maxx, miny}],
		fboard[[2]int{maxx, maxy}],
	}

	var total int
	for i, num := range nums {
		if i == 0 {
			total = num
		} else {
			total *= num
		}
	}
	return total
}

func Puzzle20() int {
	data, err := ioutil.ReadFile("input/input20")
	if err != nil {
		log.Panicf("could not read file: %s", err)
	}
	return puzzle20(string(data))
}
