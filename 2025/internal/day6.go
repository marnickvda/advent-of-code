package internal

import (
	"strconv"
	"strings"
)

type DaySix struct {
	ValueRows int
}

type Column struct {
	Type  string
	Total int
}

func (c *Column) addRow(val int) {
	if c.Type == "+" {
		c.Total += val
	} else {
		if c.Total == 0 {
			c.Total = 1
		}
		c.Total *= val
	}
}

func (d DaySix) SolvePartOne(lines []string) int {
	// start at the end so we can start incrementing or multiplying right away
	types := strings.ReplaceAll(lines[len(lines)-1], " ", "")
	totals := make([]*Column, len(types))

	for i := 0; i < len(lines)-1; i++ {
		columnIndex := 0
		numStr := ""

		for _, char := range lines[i] {
			if totals[columnIndex] == nil {
				totals[columnIndex] = &Column{
					Type: string(types[columnIndex]),
				}
			}

			if char == ' ' {
				if numStr != "" {
					val, _ := strconv.Atoi(numStr)
					totals[columnIndex].addRow(val)
					numStr = ""
					columnIndex++
				}
			} else {
				numStr += string(char)
			}
		}

		// for the last item
		val, _ := strconv.Atoi(numStr)
		totals[columnIndex].addRow(val)
	}

	sum := 0
	for _, col := range totals {
		sum += col.Total
	}

	return sum
}

func (d DaySix) SolvePartTwo(lines []string) int {
	values := make([]string, len(lines[0]))
	types := strings.ReplaceAll(lines[len(lines)-1], " ", "")

	// create the numbers per column index
	for i := 0; i < len(lines[0]); i++ {
		for j := range d.ValueRows {
			if lines[j][i] != ' ' {
				values[i] += string(lines[j][i])
			}
		}
	}

	// collect the totals per computation type (*, +) based on columns
	totals := make([]int, len(types))
	vI := 0
	for cI := 0; cI < len(types); cI++ {
		for vI < len(values) && values[vI] != "" {
			val, _ := strconv.Atoi(values[vI])

			if types[cI] == '*' {
				if totals[cI] == 0 {
					totals[cI] = 1
				}
				totals[cI] *= val
			} else {
				totals[cI] += val
			}

			// to next value index
			vI++
		}

		// skip empty column between computation columns
		vI++
	}

	sum := 0
	for _, total := range totals {
		sum += total
	}

	return sum
}
