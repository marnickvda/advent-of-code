package internal

import (
	"fmt"
)

type DayFour struct{}

type Grid [][]int

type Position struct {
	x, y int
}

var adjDelta = []Position{
	{x: -1, y: -1}, // top-left
	{x: 0, y: -1},  // top-mid
	{x: 1, y: -1},  // top-right
	{x: -1, y: 0},  // left
	{x: 1, y: 0},   // right
	{x: -1, y: 1},  // bot-left
	{x: 0, y: 1},   // bot-mid
	{x: 1, y: 1},   // bot-right
}

func (g Grid) debug(rolls map[Position]*int) {
	for y, row := range g {
		for x, col := range row {
			if _, ok := rolls[Position{x, y}]; ok {
				fmt.Printf("%d", col)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (g Grid) getLines(rolls map[Position]*int) []string {
	lines := make([]string, 0, len(g))
	for y, row := range g {
		line := ""
		for x := range row {
			if _, ok := rolls[Position{x, y}]; ok {
				line += "@"
			} else {
				line += "."
			}
		}

		lines = append(lines, line)
	}
	return lines
}

type Rolls map[Position]*int

func (r *Rolls) accessibleCount() int {
	accessible := 0
	for _, c := range *r {
		if c != nil && *c < 4 {
			accessible++
		}
	}
	return accessible
}

func getState(lines []string) (Grid, Rolls, int) {
	grid := make(Grid, len(lines))
	for y := range grid {
		grid[y] = make([]int, len(lines[y]))
	}

	rolls := make(Rolls)

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == '@' {
				rolls[Position{x, y}] = &grid[y][x]

				for _, dp := range adjDelta {
					dx := dp.x + x
					dy := dp.y + y

					if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid) {
						continue
					}

					grid[dy][dx]++
				}
			} else {
				// reset count for non-rolls
				grid[y][x] = 0
			}
		}
	}

	count := rolls.accessibleCount()
	fmt.Printf("Remove %d rolls of paper:\n", count)
	grid.debug(rolls)

	return grid, rolls, count
}

func (d DayFour) SolvePartOne(lines []string) int {
	_, _, accessibleRolls := getState(lines)

	return accessibleRolls
}

func (d DayFour) SolvePartTwo(lines []string) int {
	accessibleRolls := 0
	currentLines := lines

	for {
		grid, previousState, count := getState(currentLines)

		if count == 0 {
			break
		}

		accessibleRolls += count

		// delete rolls that should be deleted
		for p, c := range previousState {
			if *c < 4 {
				delete(previousState, p)
			}
		}
		currentLines = grid.getLines(previousState)
	}

	return accessibleRolls
}
