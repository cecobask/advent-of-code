package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	cells := make([][]rune, len(in))
	for i := 0; i < len(cells); i++ {
		cells[i] = make([]rune, len(in[0]))
	}
	for i, line := range in {
		for j, c := range line {
			cells[i][j] = c
		}
	}
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if unicode.IsDigit(cells[i][j]) {
				num := getNumber(cells, j, i)
				for j < len(cells[i]) && unicode.IsDigit(cells[i][j]) {
					cells[i][j] = rune(num)
					j++
				}
				j--
			} else if cells[i][j] == '.' {
				cells[i][j] = rune(0)
			} else {
				cells[i][j] = rune(-1)
			}
		}
	}
	sum := rune(0)
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if cells[i][j] > 0 {
				sum += addValid(cells, j, i)
			}
		}
	}
	return int(sum)
}

func two() int {
	in := parseInput(input)
	cells := make([][]rune, len(in))
	for i := 0; i < len(cells); i++ {
		cells[i] = make([]rune, len(in[0]))
	}
	for i, line := range in {
		for j, c := range line {
			cells[i][j] = c
		}
	}
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if unicode.IsDigit(cells[i][j]) {
				num := getNumber(cells, j, i)
				for j < len(cells[i]) && unicode.IsDigit(cells[i][j]) {
					cells[i][j] = rune(num)
					j++
				}
				j--
			} else if cells[i][j] != '*' {
				cells[i][j] = rune(0)
			} else {
				cells[i][j] = rune(-1)
			}
		}
	}
	sum := rune(0)
	stars := make(map[string]map[int]bool)
	ids := make(map[int]rune)
	id := 0
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if cells[i][j] > 0 {
				ids[id] = cells[i][j]
				addValidTwo(cells, j, i, stars, id)
				id++
			}
		}
	}
	for _, list := range stars {
		if len(list) == 2 {
			num := rune(1)
			for val := range list {
				num *= ids[val]
			}
			sum += num
		}
	}
	return int(sum)
}

func addValid(grid [][]rune, x int, y int) rune {
	num := grid[y][x]
	grid[y][x] = rune(0)
	right := rune(0)
	if x+1 < len(grid[y]) && grid[y][x+1] > rune(0) {
		right = addValid(grid, x+1, y)
	}
	if right > 0 {
		return right
	}
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i > 0 && i < len(grid) && j > 0 && j < len(grid[i]) && grid[i][j] < 0 {
				return num
			}
		}
	}
	return 0
}

func addValidTwo(grid [][]rune, x int, y int, m map[string]map[int]bool, id int) rune {
	num := grid[y][x]
	grid[y][x] = rune(0)
	right := rune(0)
	if x+1 < len(grid[y]) && grid[y][x+1] > rune(0) {
		right = addValidTwo(grid, x+1, y, m, id)
	}
	if right > 0 {
		return right
	}
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i > 0 && i < len(grid) && j > 0 && j < len(grid[i]) && grid[i][j] < 0 {
				key := fmt.Sprint(i) + "," + fmt.Sprint(j)
				if m[key] == nil {
					m[key] = make(map[int]bool)
				}
				m[key][id] = true
				return num
			}
		}
	}
	return 0
}

func getNumber(grid [][]rune, x int, y int) int {
	if !unicode.IsDigit(grid[y][x]) {
		return 0
	}
	out := ""
	for x < len(grid[y]) && unicode.IsDigit(grid[y][x]) {
		out += string(grid[y][x])
		x++
	}
	ret, _ := strconv.Atoi(out)
	return ret
}
