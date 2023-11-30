package main

import (
	_ "embed"
	"fmt"
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
	var result int
	in := parseInput(input)
	for _, line := range in {
		fmt.Println(line)
	}
	return result
}

func two() int {
	var result int
	in := parseInput(input)
	for _, line := range in {
		fmt.Println(line)
	}
	return result
}
