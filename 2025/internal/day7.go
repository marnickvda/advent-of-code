package internal

type DaySeven struct{}

func (d DaySeven) SolvePartOne(lines []string) int {
	beams := make(map[int]bool)

	splitCount := 0

	for _, line := range lines {
		for j, char := range line {
			if char == 'S' {
				beams[j] = true
			}

			if char == '^' && beams[j] {
				if j-1 >= 0 {
					beams[j-1] = true
				}

				if j+1 < len(line) {
					beams[j+1] = true
				}

				splitCount++
				beams[j] = false
			}
		}
	}

	return splitCount
}

func (d DaySeven) SolvePartTwo(lines []string) int {
	beams := make(map[int]int)

	for _, line := range lines {
		for j, char := range line {
			if char == 'S' {
				beams[j] = 1
			}

			if char == '^' && beams[j] > 0 {
				if j-1 >= 0 {
					beams[j-1] += beams[j]
				}

				if j+1 < len(line) {
					beams[j+1] += beams[j]
				}

				beams[j] = 0
			}
		}
	}

	count := 0
	for _, v := range beams {
		count += v
	}

	return count
}
