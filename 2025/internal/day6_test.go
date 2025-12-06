package internal

import "testing"

var sixInput = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestDaySixPartTwo(t *testing.T) {
	d := DaySix{ValueRows: 3}
	sol := d.SolvePartTwo(sixInput)

	if sol != 3263827 {
		t.Errorf("SolvePartTwo = %d; want 3263827", sol)
	}
}
