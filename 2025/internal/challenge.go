package internal

type Challenge interface {
	SolvePartOne(lines []string) any
	SolvePartTwo(lines []string) any
}
