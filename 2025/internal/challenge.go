package internal

type Challenge interface {
	SolvePartOne(lines []string) int
	SolvePartTwo(lines []string) int
}
