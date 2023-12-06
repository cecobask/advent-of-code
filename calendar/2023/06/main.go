package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part one:", one())
	fmt.Println("part two:", two())
}

func parseInput() []string {
	return strings.Split(input, "\n")
}

func one() int {
	result := 1
	races := parseRaces(parseInput())
	for _, r := range races {
		result *= r.calculateScenarios()
	}
	return result
}

func two() int {
	r := parseRace(parseInput())
	return r.calculateScenarios()
}

func parseRace(lines []string) race {
	duration, err := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	if err != nil {
		panic(err)
	}
	return race{
		duration: duration,
		distance: distance,
	}
}

func parseRaces(lines []string) []race {
	var races []race
	durations := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	strings.Join(durations, "")
	strings.Join(distances, "")
	for i := range durations {
		duration, err := strconv.Atoi(durations[i])
		if err != nil {
			panic(err)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}
		races = append(races, race{
			duration: duration,
			distance: distance,
		})
	}
	return races
}

func (r *race) calculateScenarios() int {
	nums := make([]int, 0)
	for speed := 1; speed < r.duration/2+1; speed++ {
		travel := r.duration - speed
		totalTraveled := travel * speed
		if totalTraveled > r.distance {
			nums = append(nums, totalTraveled)
		}
	}
	return len(nums) * 2
}

type race struct {
	duration int
	distance int
}
