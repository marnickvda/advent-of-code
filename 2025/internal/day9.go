package internal

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type DayNine struct{}

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

	type Coords struct {
		x, y int
	}

	fills := make(map[Coords]bool)
	for i, tile := range tiles {
		nextTile := tiles[(i+1)%len(tiles)]

		if tile.x == nextTile.x {
			for y := math.Min(float64(tile.y), float64(nextTile.y)); y < math.Max(float64(tile.y), float64(nextTile.y)); y++ {
				fills[Coords{tile.x, int(y)}] = true
			}
		} else if tile.y == nextTile.y {
			for x := math.Min(float64(tile.x), float64(nextTile.x)); x < math.Max(float64(tile.x), float64(nextTile.x)); x++ {
				fills[Coords{int(x), tile.y}] = true
			}
		} else {
			panic("this shouldn't happen")
		}
	}

	for i, tile := range tiles {
		nextTile := tiles[(i+1)%len(tiles)]

		if tile.x == nextTile.x {
		} else if tile.y == nextTile.y {
		} else {
			panic("this shouldn't happen")
		}
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

			minX, maxX := math.Min(float64(tiles[i].x), float64(tiles[j].x)), math.Max(float64(tiles[i].x), float64(tiles[j].x))
			minY, maxY := math.Min(float64(tiles[i].y), float64(tiles[j].y)), math.Max(float64(tiles[i].y), float64(tiles[j].y))
			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					fills[Coords{int(x), int(y)}] = true
				}
			}
		}
	}

	slices.SortFunc(areas, func(a, b struct {
		t1   Tile
		t2   Tile
		area int
	}) int {
		return cmp.Compare(a.area, b.area)
	})

	for _, area := range areas {

	}

	return maxArea
}

func (d DayNine) SolvePartTwo(lines []string) int {
	return -1
}
