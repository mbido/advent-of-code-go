package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	res := 0
	for _, r := range strings.Split(input, "\n") {
		r := strings.Split(r, " ")
		a := r[0][0] - 'A'
		b := r[1][0] - 'X'
		score := 0
		if (a+1)%3 == b {
			// that's a win
			score = 6
		} else if a == b {
			// tie
			score = 3
		} else if (a+2)%3 == b {
			// that's a loss
			score = 0
		}
		res += score + int(b) + 1
	}
	return res
}

func part_2(input string) any {
	res := 0
	for _, r := range strings.Split(input, "\n") {
		r := strings.Split(r, " ")
		a := r[0][0] - 'A'
		b := r[1][0]
		score := 0
		value := 0
		if b == 'X' {
			// need to loose
			score = 0
			value = (int(a) + 2) % 3
		} else if b == 'Y' {
			// need to tie
			score = 3
			value = int(a)
		} else {
			// need to win
			score = 6
			value = (int(a) + 1) % 3
		}

		res += score + value + 1
	}
	return res
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
