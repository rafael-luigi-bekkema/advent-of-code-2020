package main

type Tile struct {
	flip  bool
	coord Coord
}

type Coord struct {
	x, y int
}

func puzzle24(data []string) int {
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
		current.flip = !current.flip
	}

	var count int
	for _, tile := range tiles {
		if tile.flip {
			count++
		}
	}
	return count
}

func (s *Solutions) Puzzle24() int {
	data := readFile("input/input24")
	return puzzle24(data)
}
