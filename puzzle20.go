package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const p20width = 10

type p20tile struct {
	num  int
	data [100]bool
}

func (t *p20tile) cut() [64]bool {
	var result [64]bool
	var j int
	for i, val := range t.data {
		row := i / p20width
		col := i % p20width
		if row == 0 || row == 9 {
			continue
		}
		if col == 0 || col == 9 {
			continue
		}
		result[j] = val
		j++
	}
	return result
}

func rotate(dest []bool, width int) {
	data := append([]bool{}, dest...)
	for i, val := range data {
		row := i / width
		col := i % width
		newrow := col
		newcol := width - 1 - row
		dest[newrow*width+newcol] = val
	}
}

func hflip(dest []bool, width int) {
	data := append([]bool{}, dest...)
	for i, val := range data {
		row := i / width
		col := i % width
		newcol := width - 1 - col
		newrow := row
		dest[newrow*width+newcol] = val
	}
}

func print(data []bool, width int) {
	for i, val := range data {
		if i%width == 0 {
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

func monsters(data []bool, width int) int {
	monster := "                  # " +
		"#    ##    ##    ###" +
		" #  #  #  #  #  #   "
	mWidth := 20
	var count int
	for i := range data {
		row := i / width
		col := i % width
		if col+20 > width {
			continue
		}
		if row+3 > width {
			continue
		}
		found := true
		for j, v := range []byte(monster) {
			if v != '#' {
				continue
			}
			mrow := j / mWidth
			mcol := j % mWidth

			if !data[(mrow+row)*width+(mcol+col)] {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}
	return count
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

type p20result struct {
	board                  []p20pos
	minx, maxx, miny, maxy int
}

func puzzle20solve(data string) p20result {
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
	var result p20result

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
						if newpos.x < result.minx {
							result.minx = newpos.x
						}
						if newpos.x > result.maxx {
							result.maxx = newpos.x
						}
						if newpos.y < result.miny {
							result.miny = newpos.y
						}
						if newpos.y > result.maxy {
							result.maxy = newpos.y
						}
						break Outer
					}
					rotate(dTile.data[:], p20width)
				}
				hflip(dTile.data[:], p20width)
			}
		}
	}

	if len(tiles) > 0 {
		panic(fmt.Sprintf("pieces remaining: %d", len(tiles)))
	}

	result.board = board
	return result
}

func puzzle20(data string) int {
	result := puzzle20solve(data)

	fboard := make(map[[2]int]int, len(result.board))
	for _, tile := range result.board {
		fboard[[2]int{tile.x, tile.y}] = tile.tile.num
	}
	nums := []int{
		fboard[[2]int{result.minx, result.miny}],
		fboard[[2]int{result.minx, result.maxy}],
		fboard[[2]int{result.maxx, result.miny}],
		fboard[[2]int{result.maxx, result.maxy}],
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
	return puzzle20(string(data[:len(data)-1]))
}

func puzzle20b(data string) int {
	result := puzzle20solve(data)
	picture := make([]bool, len(result.board)*64)
	width := int(math.Sqrt(float64(len(result.board)))) * 8

	for _, tile := range result.board {
		col := (result.minx*-1 + tile.x) * 8
		row := (result.miny*-1 + tile.y) * 8
		for i, c := range tile.tile.cut() {
			scol := i % 8
			srow := i / 8
			picture[(row+srow)*width+(col+scol)] = c
		}
	}

	var hashCount int
	for _, c := range picture {
		if c {
			hashCount++
		}
	}

	for i := 0; i < 4; i++ {
		m := monsters(picture[:], width)
		if m > 0 {
			hashCount -= m * 15
			break
		}
		rotate(picture[:], width)
	}
	return hashCount
}

func Puzzle20b() int {
	data, err := ioutil.ReadFile("input/input20")
	if err != nil {
		log.Panicf("could not read file: %s", err)
	}
	return puzzle20b(string(data[:len(data)-1]))
}
