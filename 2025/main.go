package main

import (
	"fmt"

	"github.com/marnickvda/aoc/2025/inputs"
	"github.com/marnickvda/aoc/2025/internal"
)

func main() {
	lines, _ := inputs.ReadInput(2)

	//d1 := internal.NewDayOne()
	//sol := d1.SolvePartTwo(lines)

	d2 := internal.DayTwo{}
	//sol := d2.SolvePartOne(lines)
	sol := d2.SolvePartTwo(lines)

	fmt.Println(sol)
}
