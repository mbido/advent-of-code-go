package main

import (
	"aoc-in-go/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isSafe(n []int) bool {
	if utils.IsSorted(n, func(a, b int) bool { return a < b }) || utils.IsSorted(n, func(a, b int) bool { return a > b }) {
		for i := 0; i < len(n)-1; i++ {
			if utils.Abs(n[i]-n[i+1]) < 1 || utils.Abs(n[i]-n[i+1]) > 3 {
				return false
			}
		}
		return true
	}
	return false
}

func part_1(input string) any {
	res := 0
	for _, c := range strings.Split(input, "\n") {
		n := utils.Nums(c)
		if isSafe(n) {
			res++
		}
	}
	return res
}

func part_2(input string) any {

	res := 0
	for _, c := range strings.Split(input, "\n") {
		n := utils.Nums(c)
		for i := range n {

			n2 := make([]int, len(n))
			copy(n2, n)

			utils.RemoveAtIndex(&n2, i)
			if isSafe(n2) {
				res++
				break
			}
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
