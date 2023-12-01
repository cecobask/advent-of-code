package main

import (
	_ "embed"
	"fmt"
	"regexp"
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
	return sum(in, digits())
}

func two() int {
	in := parseInput(input)
	return sum(in, append(digits(), numberNames()...))
}

func sum(in, nums []string) int {
	var result int
	first := regexp.MustCompile(`(` + strings.Join(nums, "|") + `)`)
	last := regexp.MustCompile(`.*` + first.String())
	for _, s := range in {
		result += 10 * (slices.Index(nums, first.FindStringSubmatch(s)[1])%9 + 1)
		result += slices.Index(nums, last.FindStringSubmatch(s)[1])%9 + 1
	}
	return result
}

func digits() []string {
	return []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
}

func numberNames() []string {
	return []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}
