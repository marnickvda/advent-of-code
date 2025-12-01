package internal

import "testing"

func TestDayOnePartTwo(t *testing.T) {
	d := DayOne{c: 50}

	sol := d.SolvePartTwo([]string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	})

	if sol != 6 {
		t.Error("expected 6, got: ", sol)
	}
}

func TestDayOnePartTwoPt2(t *testing.T) {
	d := DayOne{c: 50}

	sol := d.SolvePartTwo([]string{
		"R1000",
	})

	if sol != 10 {
		t.Error("expected 10, got: ", sol)
	}
}
