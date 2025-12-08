package internal

import (
	"strings"
	"testing"
)

var eightInput = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestDayEightPartOne(t *testing.T) {
	d := DayEight{
		MaxConnections: 10,
	}
	sol := d.SolvePartOne(strings.Split(eightInput, "\n"))
	if sol != 40 {
		t.Errorf("solution was incorrect, got: %d, want: %d.", sol, 40)
	}
}

func TestDayEightPartTwo(t *testing.T) {
	d := DayEight{}
	sol := d.SolvePartTwo(strings.Split(eightInput, "\n"))
	if sol != 25272 {
		t.Errorf("solution was incorrect, got: %d, want: %d.", sol, 25272)
	}
}
