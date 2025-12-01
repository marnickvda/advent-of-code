package challenges

import (
	"math"
	"regexp"
	"strconv"
)

type DayThirteen struct{}

func (DayThirteen) SolvePartOne(input []string) (int, error) {
	cs := parseClawContraptions(&input)

	m := 0

	for _, c := range cs {
		tokens := bruteforce(c.a, c.b, c.t)

		m += tokens
	}

	return m, nil
}

func (DayThirteen) SolvePartTwo(input []string) (int, error) {
	cs := parseClawContraptions(&input)

	sum := 0

	memo = make(map[[2]int](struct {
		a, b, cost int
	}))

	for _, c := range cs {
		c.t.x += 10000000000000
		c.t.y += 10000000000000

		_, _, cost := findMinimumCost(c.a, c.b, c.t)

		if cost != math.MaxInt {
			sum += cost
		}
	}

	return sum, nil
}

type Button struct {
	dx int
	dy int
}

type Location struct {
	x int
	y int
}

type ClawContraption struct {
	a Button
	b Button
	t Location
}

func parseClawContraptions(input *[]string) []*ClawContraption {
	e := []*ClawContraption{}

	for i := range *input {
		if i%4 == 0 {
			e = append(e, &ClawContraption{})
		}

		if i%4 == 3 {
			continue
		}

		c := e[i/4]

		r := regexp.MustCompile(`([0-9])+`)
		s := r.FindAllString((*input)[i], 2)

		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		switch i % 4 {
		case 0:
			c.a = Button{dx: x, dy: y}
		case 1:
			c.b = Button{dx: x, dy: y}
		case 2:
			c.t = Location{x: x, y: y}
		}

	}

	return e
}

var memo map[[2]int](struct {
	a, b, cost int
})

func findMinimumCost(a, b Button, l Location) (int, int, int) {
	aX, aY, aCost := a.dx, a.dy, 3
	bX, bY, bCost := b.dx, b.dy, 1
	targetX, targetY := l.x, l.y

	key := [2]int{targetX, targetY}
	if result, found := memo[key]; found {
		return result.a, result.b, result.cost
	}

	minCost := math.MaxInt // Initialize with a very large number
	optimalA, optimalB := -1, -1

	// Iterate through possible values of a (button A presses)
	for a := 0; a <= targetX/aX && a <= targetY/aY; a++ {
		// Compute remaining X and Y after pressing A
		remainingX := targetX - a*aX
		remainingY := targetY - a*aY

		// Check if the remaining values can be achieved with button B
		if remainingX >= 0 && remainingY >= 0 && remainingX*bY == remainingY*bX {
			// Calculate b (button B presses)
			b := remainingX / bX // Division is valid because increments are aligned

			// Calculate the total cost
			cost := a*aCost + b*bCost

			// Update minimum cost and optimal values
			if cost < minCost {
				minCost = cost
				optimalA = a
				optimalB = b
			}
		}
	}

	// Store the result in the memoization map
	memo[key] = struct {
		a, b, cost int
	}{optimalA, optimalB, minCost}

	return optimalA, optimalB, minCost

}

func bruteforce(a, b Button, l Location) int {
	cheapest := 401
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i*a.dx+j*b.dx == l.x && i*a.dy+j*b.dy == l.y && i*3+j < cheapest {
				cheapest = i*3 + j
			}
		}
	}

	if cheapest == 401 {
		return 0
	}

	return cheapest
}
