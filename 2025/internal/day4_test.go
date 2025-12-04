package internal

import "testing"

var exampleInput = []string{
	"..@@.@@@@.",
	"@@@.@.@.@@",
	"@@@@@.@.@@",
	"@.@@@@..@.",
	"@@.@@@@.@@",
	".@@@@@@@.@",
	".@.@.@.@@@",
	"@.@@@.@@@@",
	".@@@@@@@@.",
	"@.@.@@@.@.",
}

func TestDayFourPartOne(t *testing.T) {
	d := DayFour{}
	sol := d.SolvePartOne(exampleInput)

	if sol != 13 {
		t.Fatalf("SolvePartOne: expected %d, got %d", 13, sol)
	}
}

func TestDayFourPartTwo(t *testing.T) {
	d := DayFour{}
	sol := d.SolvePartTwo(exampleInput)

	if sol != 43 {
		t.Fatalf("SolvePartTwo: expected %d, got %d", 43, sol)
	}
}
