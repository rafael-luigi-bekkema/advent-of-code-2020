package main

import "fmt"

type Tile struct {
	black bool
	coord Coord
}

type Coord struct {
	x, y int
}

func puzzle24solve(data []string) map[Coord]*Tile {
	ref := Tile{coord: Coord{0, 0}}
	tiles := map[Coord]*Tile{
		ref.coord: &ref,
	}

	for _, line := range data {
		current := &ref
		for j := 0; j < len(line); j++ {
			c := current.coord
			switch line[j] {
			case 'e':
				c.x++
			case 's':
				switch line[j+1] {
				case 'e':
					c.y++
				case 'w':
					c.x--
					c.y++
				}
				j++
			case 'w':
				c.x--
			case 'n':
				switch line[j+1] {
				case 'e':
					c.x++
					c.y--
				case 'w':
					c.y--
				}
				j++
			}
			if c == current.coord {
				panic("oops")
			}
			if _, ok := tiles[c]; !ok {
				tiles[c] = &Tile{coord: c}
			}
			current = tiles[c]
		}
		current.black = !current.black
	}
	return tiles
}

func puzzle24(data []string) int {
	tiles := puzzle24solve(data)
	var count int
	for _, tile := range tiles {
		if tile.black {
			count++
		}
	}
	return count
}

func (s *Solutions) Puzzle24() int {
	data := readFile("input/input24")
	return puzzle24(data)
}

func puzzle24b(data []string) int {
	debug := false
	tiles := puzzle24solve(data)
	stiles := make([]*Tile, 0, len(tiles))
	for _, tile := range tiles {
		stiles = append(stiles, tile)
	}
	_ = tiles
	for day := 1; day <= 100; day++ {
		actions := make(map[Coord]bool)
		newTiles := make(map[*Tile]uint)
		for i := 0; i < len(stiles); i++ {
			tile := stiles[i]
			var blacks int
			for i := 0; i < 8; i++ {
				if i == 4 || i == 0 {
					continue
				}
				c := tile.coord
				c.x += i%3 - 1
				c.y += i/3 - 1
				if t, ok := tiles[c]; ok {
					if t.black {
						blacks++
					}
				} else {
					if val := newTiles[tile]; val < 1 {
						tiles[c] = &Tile{coord: c}
						stiles = append(stiles, tiles[c])
						newTiles[tiles[c]] = val + 1
					}
				}
			}
			if tile.black && (blacks == 0 || blacks > 2) {
				actions[tile.coord] = false
			}
			if !tile.black && blacks == 2 {
				actions[tile.coord] = true
			}
		}

		for coord, black := range actions {
			tiles[coord].black = black
		}

		if debug && ((1 <= day && day <= 10) || (day >= 20 && day%10 == 0)) {
			var count int
			for _, tile := range tiles {
				if tile.black {
					count++
				}
			}
			fmt.Printf("Day %d: %d\n", day, count)
		}
	}

	var count int
	for _, tile := range tiles {
		if tile.black {
			count++
		}
	}
	return count
}

func (s *Solutions) Puzzle24b() int {
	data := readFile("input/input24")
	return puzzle24b(data)
}
