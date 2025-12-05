package internal

import "testing"

var input = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func TestDayFivePartTwo(t *testing.T) {
	d := DayFive{}
	sol := d.SolvePartTwo(input)

	if sol != 14 {
		t.Errorf("SolvePartTwo = %d; want 14", sol)
	}
}
