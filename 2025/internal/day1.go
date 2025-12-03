package internal

import (
	"fmt"
	"math"
	"strconv"
)

type DayOne struct {
	c int
}

func NewDayOne() *DayOne {
	return &DayOne{
		c: 50,
	}
}

func (d *DayOne) rotate(n int) {
	d.c = (d.c + n) % 100
	if d.c < 0 {
		d.c = 100 + d.c
	}
}

func (d *DayOne) rotateCount(n int) int {
	c := int(math.Abs(float64(n / 100)))

	if n%100 == 0 {
		if d.c == 0 {
			return c + 1
		} else {
			return c
		}
	}

	nn := n % 100
	if d.c != 0 && (d.c+nn < 0 || d.c+nn > 100) {
		c++
	}

	d.rotate(nn)

	if d.c == 0 {
		c += 1
	}

	return c
}

func (d *DayOne) SolvePartOne(lines []string) int {
	zC := 0
	for _, l := range lines {
		n, _ := strconv.ParseInt(l[1:], 10, 0)

		if l[0] == 'R' {
			d.rotate(int(n))
		} else {
			d.rotate(-int(n))
		}

		fmt.Printf("The dial is rotated %s to point at %d\n", l, d.c)
		if d.c == 0 {
			zC++
			fmt.Println("zC++", zC)
		}
	}
	return zC
}

func (d *DayOne) SolvePartTwo(lines []string) int {
	zC := 0
	for _, l := range lines {
		n, _ := strconv.ParseInt(l[1:], 10, 0)
		delta := int(n)
		if l[0] == 'L' {
			delta = -delta
		}

		c := d.rotateCount(delta)
		zC += c
		fmt.Printf("The dial is rotated %s to point at %d", l, d.c)
		if c > 0 {
			fmt.Printf("; adding %d to count", c)
		}
		fmt.Printf("\n")
	}
	return zC
}
