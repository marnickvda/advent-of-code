package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type DayThree struct{}

func (d *DayThree) SolvePartOne(lines []string) int {
	sum := 0

	for _, line := range lines {
		maxOne, maxTwo := 0, 0

		line = strings.TrimSpace(line)

		for i, char := range line {
			val, _ := strconv.Atoi(string(char))

			if val > maxOne && i != len(line)-1 {
				maxOne = val
				maxTwo = 0
			} else if val > maxTwo {
				maxTwo = val
			}
		}

		add, _ := strconv.Atoi(fmt.Sprintf("%d%d", maxOne, maxTwo))
		sum += add
	}

	return sum
}

func (d *DayThree) SolvePartTwo(lines []string) int {
	sum := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)

		maxes := make([]int, 12)

		for i, char := range line {
			val, _ := strconv.Atoi(string(char))

			for j := range maxes {
				if val > maxes[j] && i <= len(line)-(12-j) {
					maxes[j] = val

					// zero the remaining values
					for k := j + 1; k < len(maxes); k++ {
						maxes[k] = 0
					}

					break
				}
			}
		}

		str := ""
		for _, m := range maxes {
			str += strconv.Itoa(m)
		}
		add, _ := strconv.Atoi(str)

		fmt.Printf("In %s, got %d\n", line, add)

		sum += add
	}

	return sum
}
