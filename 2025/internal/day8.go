package internal

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type DayEight struct {
	MaxConnections int
}

const DIMS = 3

type Connection struct {
	p        [DIMS]int
	q        [DIMS]int
	distance float64
}

// https://en.wikipedia.org/wiki/Euclidean_distance#Higher_dimensions
func euclideanDistance(p [DIMS]int, q [DIMS]int) float64 {
	sum := 0.0
	for i := 0; i < DIMS; i++ {
		sum += math.Pow(float64(p[i]-q[i]), 2)
	}

	return math.Sqrt(sum)
}

func posToStr(pos [DIMS]int) string {
	str := make([]string, 0, DIMS)
	for i := 0; i < DIMS; i++ {
		str = append(str, strconv.Itoa(pos[i]))
	}
	return strings.Join(str, ",")
}

func initDayEight(lines []string) (map[string]int, [][]string, []Connection) {
	boxes := make([][DIMS]int, len(lines))
	for i, line := range lines {
		pos := strings.Split(line, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])

		boxes[i] = [DIMS]int{x, y, z}
	}

	connected := make(map[string]int) // maps to index of circuits slice
	circuits := make([][]string, 0)   // keep track of group of circuits
	distances := make([]Connection, 0)
	for i := 0; i < len(boxes); i++ {
		circuits = append(circuits, []string{posToStr(boxes[i])})
		connected[posToStr(boxes[i])] = len(circuits) - 1

		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, Connection{
				p:        boxes[i],
				q:        boxes[j],
				distance: euclideanDistance(boxes[i], boxes[j]),
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	return connected, circuits, distances
}

func (d DayEight) SolvePartOne(lines []string) int {
	connected, circuits, distances := initDayEight(lines)

	i := 0
	for i < d.MaxConnections {
		d := distances[i]
		i++

		pi := connected[posToStr(d.p)]
		qi := connected[posToStr(d.q)]

		if pi == qi {
			continue
		}

		// merge circuits into p
		circuits[pi] = append(circuits[pi], circuits[qi]...)

		// map circuit q values to circuit of p
		for _, v := range circuits[qi] {
			connected[v] = pi
		}

		// delete circuit q
		circuits[qi] = nil
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	// three largest circuits, multiplied
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func (d DayEight) SolvePartTwo(lines []string) int {
	connected, circuits, distances := initDayEight(lines)

	i := 0
	for i < len(distances) {
		dist := distances[i]
		i++
		ci := -1
		pi := connected[posToStr(dist.p)]
		qi := connected[posToStr(dist.q)]

		// if connected[p] == connected[q], they are part of the same circuit
		if pi == qi {
			continue
		}

		// merge circuits into p
		circuits[pi] = append(circuits[pi], circuits[qi]...)

		// map circuit q values to circuit of p
		for _, v := range circuits[qi] {
			connected[v] = pi
		}

		ci = pi

		// delete circuit q
		circuits[qi] = nil

		if len(circuits[ci]) == len(connected) {
			return dist.p[0] * dist.q[0]
		}
	}

	return -1
}
