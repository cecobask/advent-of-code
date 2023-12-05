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
	seeds, translators := parseInput()
	fmt.Println("part one:", one(seeds, translators))
	fmt.Println("part two:", two(seeds, translators))
}

func parseInput() ([]int, [][][]int) {
	lines := strings.Split(input, "\n")
	seeds := parseNumbers(lines[0][strings.Index(lines[0], ": ")+2:])
	translators := make([][][]int, 0, 8)
	translators = append(translators, make([][]int, 0, 8))
	s := 0
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			s++
			translators = append(translators, make([][]int, 0, 8))
			continue
		}
		translators[s] = append(translators[s], parseNumbers(lines[i]))
	}
	return seeds, translators
}

func parseNumbers(s string) []int {
	values := strings.Split(s, " ")
	numbers := make([]int, len(values))
	for i := range values {
		num, err := strconv.Atoi(values[i])
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}
	return numbers
}

func transform(source int, rules [][]int) int {
	for _, rule := range rules {
		if source >= rule[1] && source < rule[1]+rule[2] {
			return rule[0] + source - rule[1]
		}
	}
	return source
}

func reverse(dest int, rules [][]int) int {
	var source int
	for _, rule := range rules {
		source = dest + rule[1] - rule[0]
		if source >= rule[1] && source < rule[1]+rule[2] {
			return source
		}
	}
	return dest
}

func one(seeds []int, translators [][][]int) int {
	result := 1 << 31
	for _, s := range seeds {
		for _, t := range translators {
			s = transform(s, t)
		}
		if s < result {
			result = s
		}
	}
	return result
}

func two(seeds []int, translators [][][]int) int {
	var s int
	var l int
	for {
		s = l
		for i := len(translators) - 1; i >= 0; i-- {
			s = reverse(s, translators[i])
		}
		for i := 0; i < len(seeds); i += 2 {
			if s > seeds[i] && s < seeds[i]+seeds[i+1] {
				return l
			}
		}
		l++
	}
}
