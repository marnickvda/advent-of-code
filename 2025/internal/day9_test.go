package internal

import (
	"strings"
	"testing"
)

var nineInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestDayNinePartOne(t *testing.T) {
	d := DayNine{}
	sol := d.SolvePartOne(strings.Split(nineInput, "\n"))
	if sol != 50 {
		t.Errorf("solution was incorrect, got: %d, want: %d.", sol, 50)
	}
}

func TestDayNinePartTwo(t *testing.T) {
	d := DayNine{}
	sol := d.SolvePartTwo(strings.Split(nineInput, "\n"))
	if sol != 24 {
		t.Errorf("solution was incorrect, got: %d, want: %d.", sol, 24)
	}
}
