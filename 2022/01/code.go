package main

import (
	"aoc-in-go/utils"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	res := 0
	for _, elf := range strings.Split(input, "\n\n") {
		wheight := utils.Sum(utils.Nums(elf))
		if wheight > res {
			res = wheight
		}
	}
	return res
}

func part_2(input string) any {
	res := []int{}
	for _, elf := range strings.Split(input, "\n\n") {
		wheight := utils.Sum(utils.Nums(elf))
		res = append(res, wheight)
		sort.Slice(res, func(i, j int) bool {
			return res[i] > res[j]
		})
	}
	return res[0] + res[1] + res[2]
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
