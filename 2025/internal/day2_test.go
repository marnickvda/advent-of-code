package internal

import "testing"

func TestDayTwoPartOne(t *testing.T) {
	d := DayTwo{}
	sol := d.SolvePartOne([]string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"})

	if sol != 1227775554 {
		t.Fatalf("expected: %s, got: %s", "1227775554", sol)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	d := DayTwo{}
	sol := d.SolvePartTwo([]string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"})

	if sol != 4174379265 {
		t.Fatalf("expected: %s, got: %s", "4174379265", sol)
	}
}
