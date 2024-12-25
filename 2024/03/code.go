package main

import (
	"aoc-in-go/utils"
	"regexp"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	res := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		n := utils.Nums(match)
		res += n[0] * n[1]
	}
	return res
}

func part_2(input string) any {
	res := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllString(input, -1)
	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
			continue
		}
		if match == "don't()" {
			do = false
			continue
		}
		if do {
			n := utils.Nums(match)
			res += n[0] * n[1]
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
