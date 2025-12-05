package internal

import (
	"sort"
	"strconv"
	"strings"
)

type DayFive struct{}

type Range struct {
	start, end int
}

func (r Range) isFresh(id int) bool {
	return id >= r.start && id <= r.end
}

func (d DayFive) ReadInput(lines []string) ([]Range, []int) {
	ranges := make([]Range, 0)
	ingredients := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		values := strings.Split(line, "-")
		start, _ := strconv.Atoi(values[0])
		if len(values) == 2 {
			end, _ := strconv.Atoi(values[1])
			ranges = append(ranges, Range{start: start, end: end})
		}

		if len(values) == 1 {
			ingredients = append(ingredients, start)
		}
	}

	return ranges, ingredients
}

func (d DayFive) SolvePartOne(lines []string) int {
	ranges, ingredients := d.ReadInput(lines)

	freshCount := 0
	for _, p := range ingredients {
		for _, r := range ranges {
			if r.isFresh(p) {
				freshCount++
				break
			}
		}
	}

	return freshCount
}

func (d DayFive) SolvePartTwo(lines []string) int {
	ranges, _ := d.ReadInput(lines)

	// sort the ranges by starting point
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// iterate through the ranges to collect 'ingredient id chunks'
	chunks := make([]Range, 0)
	chunkStart, chunkEnd := -1, -1
	for i := 0; i < len(ranges); i++ {
		// set the chunk
		if chunkStart == -1 {
			chunkStart = ranges[i].start
			chunkEnd = ranges[i].end
		}

		// if next range start is smaller than the end, merge
		if i+1 < len(ranges) && ranges[i+1].start <= chunkEnd {
			if chunkEnd < ranges[i+1].end {
				chunkEnd = ranges[i+1].end
			}
		} else {
			// end of chunk, go to the next
			chunks = append(chunks, Range{start: chunkStart, end: chunkEnd})
			chunkStart = -1
			chunkEnd = -1
		}
	}

	count := 0
	for _, p := range chunks {
		count += 1 + p.end - p.start
	}

	return count
}
