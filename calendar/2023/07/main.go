package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Hand struct {
	Cards string
	Bid   int
}

func main() {
	var hands []Hand
	for _, s := range parseInput() {
		hands = append(hands, parseHand(s))
	}
	fmt.Println(calculateWinnings(hands, false))
	fmt.Println(calculateWinnings(hands, true))
}

func parseInput() []string {
	return strings.Split(input, "\n")
}

func parseHand(line string) Hand {
	h := Hand{}
	_, err := fmt.Sscanf(line, "%s %d", &h.Cards, &h.Bid)
	if err != nil {
		panic(err)
	}
	return h
}

func calculateWinnings(hands []Hand, jokers bool) int {
	var result int
	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp(a.Cards, b.Cards, jokers)
	})
	for i, h := range hands {
		result += (i + 1) * h.Bid
	}
	return result
}

func cmp(a, b string, jokers bool) int {
	j, r := "J", "TAJBQCKDAE"
	if jokers {
		j, r = "23456789TQKA", "TAJ0QCKDAE"
	}
	kind := func(cards string) string {
		k := 0
		for _, j := range strings.Split(j, "") {
			n, t := strings.ReplaceAll(cards, "J", j), 0
			for _, s := range n {
				t += strings.Count(n, string(s))
			}
			k = slices.Max([]int{
				k,
				t,
			})
		}
		return map[int]string{
			5:  "0",
			7:  "1",
			9:  "2",
			11: "3",
			13: "4",
			17: "5",
			25: "6",
		}[k]
	}
	return strings.Compare(
		kind(a)+strings.NewReplacer(strings.Split(r, "")...).Replace(a),
		kind(b)+strings.NewReplacer(strings.Split(r, "")...).Replace(b),
	)
}
