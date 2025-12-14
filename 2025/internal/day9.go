package internal

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type DayNine struct {
}

// find the largest area
// from red tiles (input locations)

// (2,2) & (0,5)
//
// ..#.....      ..OOOO..
// ........  ->  ..OOOO..
// .....#..      ..OOOO..
//
// x1 - x2 = math.Abs(2 - 0) + 1 = 3
// y1 - y2 = math.Abs(2 - 5) + 1 = 4
// 								  --- *
//  							   12

type Tile struct {
	x, y int
}

func (t Tile) calculateM2(tile Tile) int {
	return int((math.Abs(float64(t.x-tile.x)) + 1) * (math.Abs(float64(t.y-tile.y)) + 1))
}

// naive approach: brute force all combinations
// better approach: find the top right, top left, bottom left, bottom right tiles.

func (d DayNine) SolvePartOne(lines []string) int {
	tiles := make([]Tile, len(lines))
	for i := range tiles {
		cords := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(cords[0])
		y, _ := strconv.Atoi(cords[1])

		tiles[i] = Tile{x, y}
	}

	areas := make([]struct {
		t1   Tile
		t2   Tile
		area int
	}, 0)

	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := tiles[i].calculateM2(tiles[j])

			areas = append(areas, struct {
				t1   Tile
				t2   Tile
				area int
			}{
				t1:   tiles[i],
				t2:   tiles[j],
				area: area,
			})
		}
	}

	slices.SortFunc(areas, func(a, b struct {
		t1   Tile
		t2   Tile
		area int
	}) int {
		return -1 * cmp.Compare(a.area, b.area) // DESC
	})

	return areas[0].area
}

func (d DayNine) SolvePartTwo(lines []string) int {
	tiles := make([]Tile, len(lines))
	maxX, maxY := 0, 0
	log.Println("reading tiles")
	for i := range tiles {
		cords := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(cords[0])
		y, _ := strconv.Atoi(cords[1])

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		tiles[i] = Tile{x, y}
	}
	log.Printf("maxX: %d, maxY: %d\n", maxX, maxY)

	log.Println("collecting tile fills")
	fills := make(map[Tile]bool)
	for i, tile := range tiles {
		nextTile := tiles[(i+1)%len(tiles)]

		if tile.x == nextTile.x {
			for y := int(math.Min(float64(tile.y), float64(nextTile.y))); y <= int(math.Max(float64(tile.y), float64(nextTile.y))); y++ {
				fills[Tile{tile.x, y}] = true
			}
		} else if tile.y == nextTile.y {
			for x := int(math.Min(float64(tile.x), float64(nextTile.x))); x <= int(math.Max(float64(tile.x), float64(nextTile.x))); x++ {
				fills[Tile{x, tile.y}] = true
			}
		} else {
			panic("this shouldn't happen")
		}
	}

	log.Println("collecting valid tile rectangle fills")
	for _, tile := range tiles {
		markedUp := make([]Tile, 0)
		nextY := tile.y + 1
		for nextY <= maxY {
			markedUp = append(markedUp, Tile{tile.x, nextY})

			// if we hit the 'wall', we add the marked fields as visited
			if _, ok := fills[Tile{tile.x, nextY}]; ok {
				for _, m := range markedUp {
					fills[Tile{m.x, m.y}] = true
				}
				break
			}

			nextY++
		}

		markedDown := make([]Tile, 0)
		nextY = tile.y - 1
		for nextY >= 0 {
			markedDown = append(markedDown, Tile{tile.x, nextY})

			// if we hit the 'wall', we add the marked fields as visited
			if _, ok := fills[Tile{tile.x, nextY}]; ok {
				for _, m := range markedDown {
					fills[Tile{m.x, m.y}] = true
				}
				break
			}

			nextY--
		}
		markedRight := make([]Tile, 0)
		nextX := tile.x + 1
		for nextX <= maxX {
			markedRight = append(markedRight, Tile{nextX, tile.y})

			// if we hit the 'wall', we add the marked fields as visited
			if _, ok := fills[Tile{nextX, tile.y}]; ok {
				for _, m := range markedRight {
					fills[Tile{m.x, m.y}] = true
				}
				break
			}

			nextX++
		}

		markedLeft := make([]Tile, 0)
		nextX = tile.x - 1
		for nextX >= 0 {
			markedLeft = append(markedLeft, Tile{nextX, tile.y})

			// if we hit the 'wall', we add the marked fields as visited
			if _, ok := fills[Tile{nextX, tile.y}]; ok {
				for _, m := range markedLeft {
					fills[Tile{m.x, m.y}] = true
				}
				break
			}

			nextX--
		}
	}

	log.Println("collecting valid areas")
	areas := make([]struct {
		t1   Tile
		t2   Tile
		area int
	}, 0)

	for i := range tiles {
	iter:
		for j := i + 1; j < len(tiles); j++ {
			// (x1, y2) ---- (x2, y2)
			//    |				|
			//    |				|
			// (x1, y1) ---- (x2, y1)
			x1, x2 := int(math.Min(float64(tiles[i].x), float64(tiles[j].x))), int(math.Max(float64(tiles[i].x), float64(tiles[j].x)))
			y1, y2 := int(math.Min(float64(tiles[i].y), float64(tiles[j].y))), int(math.Max(float64(tiles[i].y), float64(tiles[j].y)))

			// check if the box is fully within the borders
			for x := x1; x <= x2; x++ {
				if _, ok := fills[Tile{x, y1}]; !ok {
					continue iter
				}

				if _, ok := fills[Tile{x, y2}]; !ok {
					continue iter
				}
			}

			for y := y1; y <= y2; y++ {
				if _, ok := fills[Tile{x1, y}]; !ok {
					continue iter
				}

				if _, ok := fills[Tile{x2, y}]; !ok {
					continue iter
				}
			}

			// when we got through the checks for the borders, we can append to the valid areas
			area := tiles[i].calculateM2(tiles[j])
			log.Printf("valid area (%d m2) for (x1,y1): %d,%d + (x2,y2): %d,%d\n", area, x1, y1, x2, y2)

			areas = append(areas, struct {
				t1   Tile
				t2   Tile
				area int
			}{
				t1:   tiles[i],
				t2:   tiles[j],
				area: area,
			})
		}
	}

	slices.SortFunc(areas, func(a, b struct {
		t1   Tile
		t2   Tile
		area int
	}) int {
		return -1 * cmp.Compare(a.area, b.area) // DESC
	})

	//d.printDayNine(maxX, maxY, fills)

	return areas[0].area
}

func (d DayNine) printDayNine(maxX, maxY int, fills map[Tile]bool, markedTiles ...Tile) {
	markedMap := make(map[Tile]bool)
	for _, m := range markedTiles {
		markedMap[m] = true
	}

	for y := 0; y < maxY+3; y++ {
		for x := 0; x < maxX+3; x++ {
			if markedMap[Tile{x, y}] {
				fmt.Print("O")
			} else if fills[Tile{x, y}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
