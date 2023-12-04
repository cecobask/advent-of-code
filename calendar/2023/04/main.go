package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	in := parseInput(input)
	var result int
	for _, line := range in {
		c := parseCard(line)
		if wins := c.countWins(); wins > 0 {
			result += 1 << (wins - 1)
		}
	}
	return result
}

func two() int {
	in := parseInput(input)
	result := len(in)
	count := make([]int, result)
	for i, line := range in {
		c := parseCard(line)
		count[i] += 1
		for j := i + 1; j <= c.countWins()+i; j++ {
			count[j] += count[i]
			result += count[i]
		}
	}
	return result
}

type card struct {
	winningNums []string
	ownedNums   []string
}

func parseCard(line string) card {
	c := card{}
	nums := strings.Split(strings.Split(line, ": ")[1], " | ")
	c.winningNums = append(c.winningNums, strings.Fields(nums[0])...)
	c.ownedNums = append(c.ownedNums, strings.Fields(nums[1])...)
	return c
}

func (c card) countWins() int {
	var count int
	for _, ownedNum := range c.ownedNums {
		if slices.Contains(c.winningNums, ownedNum) {
			count += 1
		}
	}
	return count
}
