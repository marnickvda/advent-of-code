package internal

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type DayTwo struct{}

// The ranges are separated by commas (,); each range gives its first ID and last ID separated by a dash (-).

func invalidSequenceCount(start, end string) (int, []string) {
	sI, _ := strconv.Atoi(start)
	eI, _ := strconv.Atoi(end)

	invalids := []string{}

	invalidCount := 0
	for i := sI; i <= eI; i++ {
		str := strconv.Itoa(i)
		if len(str)%2 == 0 {
			if str[:len(str)/2] == str[len(str)/2:] {
				invalidCount++
				invalids = append(invalids, str)
			}
		}
	}

	return invalidCount, invalids
}

func (d *DayTwo) SolvePartOne(lines []string) any {
	line := lines[0]

	ranges := strings.Split(line, ",")
	c := 0
	for i := range ranges {
		r := strings.Split(ranges[i], "-")
		x, l := invalidSequenceCount(r[0], r[1])

		if x > 0 {
			fmt.Printf("%s has %d invalid IDs: %s\n", ranges[i], x, l)
		}

		for j := range l {
			z, _ := strconv.Atoi(l[j])
			c += z
		}
	}

	return c
}

func (d *DayTwo) SolvePartTwo(lines []string) any {
	l := lines[0]
	rs := strings.Split(l, ",")

	// set of matches
	invalids := make(map[int]bool)

	for i := range rs {
		r := strings.Split(rs[i], "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		invalid := make(map[int]bool)

		fmt.Printf("\nchecking range: %s\n", rs[i])

		for j := start; j <= end; j++ {
			sJ := strconv.Itoa(j)

			// check for all the same number
			chars := strings.Split(sJ, "")
			slices.Sort(chars)
			if len(chars) > 1 && chars[0] == chars[len(chars)-1] {
				fmt.Printf("found match %d for rem 1 ('%s' matched %d times)\n", j, chars[0], len(chars))
				invalid[j] = true
				continue
			}

			// check biggest match to the smallest possible match
			for rem := int(math.Floor(float64(len(sJ)) / 2)); rem >= 2; rem-- {

				// check if we can use this chunk as a splitter
				if len(sJ)%rem == 0 {
					// match each chunk to the next one
					prev := ""
					matched := true
					for s := range slices.Chunk(strings.Split(sJ, ""), rem) {
						if prev == "" {
							prev = strings.Join(s, "")
							continue
						}

						if prev != strings.Join(s, "") {
							matched = false
							break
						}
						prev = strings.Join(s, "")
					}

					if matched {
						fmt.Printf("found match %d for rem %d ('%s' matched %d times)\n", j, rem, strings.Join(slices.Collect(slices.Chunk(strings.Split(sJ, ""), rem))[0], ""), len(sJ)/rem)
						invalid[j] = true
						break
					}
				}
			}
		}

		fmt.Printf("%s has %d invalid IDs: %v\n", rs[i], len(invalid), slices.Collect(maps.Keys(invalid)))
		for k := range invalid {
			invalids[k] = true
		}
	}

	c := 0
	for k := range invalids {
		c += k
	}

	return c
}
