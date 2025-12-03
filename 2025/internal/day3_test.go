package internal

import "testing"

var lines = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestDayThreePartOne(t *testing.T) {
	d := DayThree{}
	sol := d.SolvePartOne(lines)

	if sol != 357 {
		t.Fatalf("expected: %s, got: %d", "357", sol)
	}
}

func TestDayThreePartTwo(t *testing.T) {
	d := DayThree{}
	sol := d.SolvePartTwo(lines)

	if sol != 3121910778619 {
		t.Fatalf("expected: %s, got: %d", "3121910778619", sol)
	}
}

func TestDayThreePartTwoPt2(t *testing.T) {
	d := DayThree{}
	sol := d.SolvePartTwo([]string{"23934332564372337559263534255825234943"})
	if sol != 996825234943 {
		t.Fatalf("expected: %s, got: %d", "996825234943", sol)
	}
}
