package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	debug := flag.Bool("d", false, "Enable debug output")
	flag.Parse()
	solutions.debug = *debug
	os.Exit(m.Run())
}

func TestPuzzle24(t *testing.T) {
	data := strings.Split(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`, "\n")
	t.Run("part 1", func(t *testing.T) {
		expect := 10
		result := puzzle24(data)
		if result != expect {
			t.Fatalf("expected %d, got %d", expect, result)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		expect := 2208
		result := puzzle24b(data)
		if result != expect {
			t.Fatalf("expected %d, got %d", expect, result)
		}
	})
}

func ExampleSolutions_Puzzle24() {
	fmt.Println(solutions.Puzzle24())

	// Output: 244
}

func ExampleSolutions_Puzzle24b() {
	fmt.Println(solutions.Puzzle24b())

	// Output: 3665
}
