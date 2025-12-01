package main

import (
	"fmt"

	"github.com/marnickvda/aoc/2025/inputs"
	"github.com/marnickvda/aoc/2025/internal"
)

func main() {
	lines, _ := inputs.ReadInput(1)

	d1 := internal.NewDayOne()
	sol := d1.SolvePartTwo(lines)

	fmt.Println(sol)
}
