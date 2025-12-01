package challenges

import (
	"fmt"
)

type DayFifteen struct{}

func (DayFifteen) SolvePartOne(input []string) (int, error) {
	m := [][]string{}

	x, y := 0, 0

	i := 0
	for i = 0; input[i] != ""; i++ {
		m = append(m, make([]string, len(input[i])))
		for j, c := range input[i] {
			m[i][j] = string(c)

			if c == '@' {
				x = j
				y = i
			}
		}
	}

	type Pos struct {
		x, y int
	}

	dirs := map[rune]Pos{
		'^': {y: -1, x: 0},
		'>': {y: 0, x: 1},
		'v': {y: 1, x: 0},
		'<': {y: 0, x: -1},
	}

	for i = i + 1; i < len(input); i++ {
		for _, c := range input[i] {
			d := dirs[c]

			s := Pos{x: x, y: y}
			swap := []Pos{s}

		loop:
			for s.x+d.x >= 0 && s.x+d.x < len(m[0]) && s.y+d.y >= 0 && s.y+d.y < len(m) {
				// fmt.Printf("Position: %+v, val: %s\t", s, m[s.y][s.x])
				switch m[s.y+d.y][s.x+d.x] {
				case "O":
					s = Pos{x: s.x + d.x, y: s.y + d.y}

					if swap[len(swap)-1] == s {
						panic(fmt.Errorf("WTF? %v\n", swap))
					}

					swap = append(swap, s)

					// fmt.Printf("Position: %+v, val: %s \n", s, m[s.y][s.x])
				case ".":
					swap = append(swap, Pos{x: s.x + d.x, y: s.y + d.y})
					// fmt.Printf("Position: %+v, empty position, swapping %s\n", swap[len(swap)-1], m[swap[len(swap)-1].y][swap[len(swap)-1].x])
					break loop
				case "#":
					swap = nil
					// fmt.Printf("Position: %+vj, BORDER! aborting.\n", s)
					break loop
				}
			}

			if swap != nil {
				for z := len(swap) - 1; z >= 1; z-- {
					a := swap[z]
					b := swap[z-1]

					// fmt.Printf("1. Swapping %v (%s) with %v (%s)\n", a, m[a.y][a.x], b, m[b.y][b.x])
					m[a.y][a.x] = m[b.y][b.x]

					// fmt.Printf("2. Swapping %v (%s) with '.'\n", b, m[b.y][b.x])
					m[b.y][b.x] = "."
				}

				x = swap[1].x
				y = swap[1].y
			}

			// fmt.Printf("DIRECTION %s ('%v') \n", string(c), d)
			// printMatrix(1, m)
		}
	}

	sum := 0
	for y := 1; y < len(m)-1; y++ {
		for x := 1; x < len(m[0])-1; x++ {
			if m[y][x] == "O" {
				sum += 100*y + x
			}
		}
	}

	return sum, nil
}

func (DayFifteen) SolvePartTwo(input []string) (int, error) {
	m := [][]string{}

	// x, y := 0, 0

	i := 0
	for i = 0; input[i] != ""; i++ {
		m = append(m, make([]string, len(input[i])*2))
		for j := 0; j < len(input[i]); j += 1 {
			c := string(input[i][j])

			xj := j * 2

			switch c {
			case "#", ".":
				m[i][xj] = c
				m[i][xj+1] = c
			case "@":
				m[i][xj] = "@"
				m[i][xj+1] = "."
				// x = xj
				// y = i
			case "O":
				m[i][xj] = "["
				m[i][xj+1] = "]"
			}
		}
	}

	printMatrix(1, m)

	return 0, nil
}
