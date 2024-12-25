package main

import (
	"aoc-in-go/utils"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func useImports() {
	utils.PrintHello()
}

func main() {
	aoc.Harness(run)
}

func parse(input string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {
		n := utils.Nums(line)
		left = append(left, n[0])
		right = append(right, n[1])
	}

	return left, right
}

func part_1(input string) any {
	left, right := parse(input)

	sort.Ints(left)
	sort.Ints(right)

	res := 0

	for i, l := range left {
		r := right[i]
		res += utils.Abs(r - l)
	}

	return res
}

func part_2(input string) any {
	left, right := parse(input)

	res := 0

	for _, l := range left {
		res += utils.CountInt(right, l) * l
	}

	return res
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
