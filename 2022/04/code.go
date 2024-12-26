package main

import (
	"aoc-in-go/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		n := utils.Nums(line)
		if n[0] >= n[2] && n[1] <= n[3] || n[2] >= n[0] && n[3] <= n[1] {
			res++
		}
	}
	return res
}

func part_2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		n := utils.Nums(line)
		if n[2] >= n[0] && n[2] <= n[1] || n[3] >= n[0] && n[3] <= n[1] || n[0] >= n[2] && n[0] <= n[3] || n[1] >= n[2] && n[1] <= n[3] {
			res++
		}
	}
	return res
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
