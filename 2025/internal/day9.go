package internal

import (
	"cmp"
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
	for i := range tiles {
		cords := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(cords[0])
		y, _ := strconv.Atoi(cords[1])
		tiles[i] = Tile{x, y}
	}

	// Ray casting algorithm
	isInside := func(p Tile) bool {
		inside := false
		j := len(tiles) - 1

		for i := 0; i < len(tiles); i++ {
			xi, yi := tiles[i].x, tiles[i].y
			xj, yj := tiles[j].x, tiles[j].y

			if ((yi > p.y) != (yj > p.y)) &&
				(p.x < (xj-xi)*(p.y-yi)/(yj-yi)+xi) {
				inside = !inside
			}
			j = i
		}
		return inside
	}

	// Check if point is on boundary
	isOnBoundary := func(p Tile) bool {
		for i := 0; i < len(tiles); i++ {
			j := (i + 1) % len(tiles)

			if tiles[i].x == tiles[j].x && tiles[i].x == p.x { // Check vertical edge
				minY := min(tiles[i].y, tiles[j].y)
				maxY := max(tiles[i].y, tiles[j].y)
				if p.y >= minY && p.y <= maxY {
					return true
				}
			} else if tiles[i].y == tiles[j].y && tiles[i].y == p.y { // Check horizontal edge
				minX := min(tiles[i].x, tiles[j].x)
				maxX := max(tiles[i].x, tiles[j].x)
				if p.x >= minX && p.x <= maxX {
					return true
				}
			}
		}
		return false
	}

	// Check if tile is valid (inside or on boundary)
	isValid := func(p Tile) bool {
		return isOnBoundary(p) || isInside(p)
	}

	candidates := make([]struct {
		i, j, area int
	}, 0)

	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := tiles[i].calculateM2(tiles[j])
			candidates = append(candidates, struct{ i, j, area int }{i, j, area})
		}
	}

	slices.SortFunc(candidates, func(a, b struct{ i, j, area int }) int {
		return -1 * cmp.Compare(a.area, b.area) // DESC
	})

	for _, c := range candidates {
		x1 := min(tiles[c.i].x, tiles[c.j].x)
		x2 := max(tiles[c.i].x, tiles[c.j].x)
		y1 := min(tiles[c.i].y, tiles[c.j].y)
		y2 := max(tiles[c.i].y, tiles[c.j].y)

		// check 4 corners first
		if !isValid(Tile{x1, y1}) || !isValid(Tile{x2, y2}) ||
			!isValid(Tile{x1, y2}) || !isValid(Tile{x2, y1}) {
			continue
		}

		valid := true

		// Check top and bottom edges
		for x := x1; x <= x2 && valid; x++ {
			if !isValid(Tile{x, y1}) || !isValid(Tile{x, y2}) {
				valid = false
			}
		}

		// Check left and right edges
		for y := y1; y <= y2 && valid; y++ {
			if !isValid(Tile{x1, y}) || !isValid(Tile{x2, y}) {
				valid = false
			}
		}

		if valid {
			log.Printf("Biggest Valid Area: %d, tiles: (%d,%d) and (%d,%d)",
				c.area, tiles[c.i].x, tiles[c.i].y, tiles[c.j].x, tiles[c.j].y)
			return c.area
		}
	}

	return -1
}
