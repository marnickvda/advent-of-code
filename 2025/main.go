package main

import (
	"flag"
	"log"
	"time"

	"github.com/marnickvda/aoc/2025/inputs"
	"github.com/marnickvda/aoc/2025/internal"
)

var flags struct {
	day  int
	part int
}

var challenges = map[int]internal.Challenge{
	1: &internal.DayOne{},
	2: &internal.DayTwo{},
	3: &internal.DayThree{},
	4: &internal.DayFour{},
	5: &internal.DayFive{},
	6: &internal.DaySix{
		ValueRows: 4,
	},
	7: &internal.DaySeven{},
	8: &internal.DayEight{
		MaxConnections: 1000,
	},
}

func main() {
	flag.IntVar(&flags.day, "day", 0, "Challenge day to execute")
	flag.IntVar(&flags.part, "part", 1, "Part of the challenge to execute")

	flag.Parse()

	if flags.day <= 0 {
		log.Fatalln("-day flag must be positive")
	}

	if flags.part != 1 && flags.part != 2 {
		log.Fatalln("-part flag must be 1 or 2")
	}

	input, err := inputs.ReadInput(flags.day)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d lines of input\n", len(input))

	c := challenges[flags.day]
	if c == nil {
		log.Fatal("Challenge does not exist, did you forget to add it?")
	}

	var ans int
	var elapsed time.Duration

	if flags.part == 1 {
		start := time.Now()
		ans = c.SolvePartOne(input)
		elapsed = time.Since(start)
	}

	if flags.part == 2 {
		start := time.Now()
		ans = c.SolvePartTwo(input)
		elapsed = time.Since(start)
	}

	log.Printf("Solution to day %d part %d is '%d' (in %s)", flags.day, flags.part, ans, elapsed)
}
