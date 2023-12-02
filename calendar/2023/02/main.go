package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("part one:", one())
	fmt.Println("part two:", two())
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func one() int {
	var sum int
	in := parseInput(input)
	for _, line := range in {
		g, err := parseGame(line)
		if err != nil {
			panic(err)
		}
		if g.isValid() {
			sum += g.id
		}
	}
	return sum
}

func two() int {
	var sumOfPowers int
	in := parseInput(input)
	for _, line := range in {
		g, err := parseGame(line)
		if err != nil {
			panic(err)
		}
		powers := 1
		for colour := range colourConfig() {
			powers *= g.getMaxForColour(colour)
		}
		sumOfPowers += powers
	}
	return sumOfPowers
}

func colourConfig() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

type game struct {
	id     int
	reds   []int
	greens []int
	blues  []int
}

func (g *game) getMaxForColour(colour string) int {
	switch colour {
	case "red":
		return slices.Max(g.reds)
	case "green":
		return slices.Max(g.greens)
	case "blue":
		return slices.Max(g.blues)
	default:
		return -1
	}
}

func (g *game) isValid() bool {
	for colour, maxUnits := range colourConfig() {
		if g.getMaxForColour(colour) > maxUnits {
			return false
		}
	}
	return true
}

func parseGame(line string) (*game, error) {
	g := &game{
		reds:   make([]int, 0),
		greens: make([]int, 0),
		blues:  make([]int, 0),
	}
	parts := strings.Split(line, ":")
	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		return nil, err
	}
	g.id = id
	for _, turn := range strings.Split(parts[1][1:], ";") {
		for _, color := range strings.Split(turn, ",") {
			colorInfo := strings.Split(strings.TrimSpace(color), " ")
			quantity, err := strconv.Atoi(colorInfo[0])
			if err != nil {
				return nil, err
			}
			switch colorInfo[1] {
			case "red":
				g.reds = append(g.reds, quantity)
			case "green":
				g.greens = append(g.greens, quantity)
			case "blue":
				g.blues = append(g.blues, quantity)
			}
		}
	}
	return g, nil
}
