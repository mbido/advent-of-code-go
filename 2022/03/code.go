package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func priority(a rune) int {
	if a < 'a' {
		return int(a-'A') + 27
	}
	return int(a-'a') + 1
}

func part_1(input string) any {
	res := 0
	for _, r := range strings.Split(input, "\n") {
		left := r[:len(r)/2]
		right := r[len(r)/2:]
		for _, c := range left {
			if strings.Count(right, string(c)) > 0 {
				res += priority(c)
				break
			}
		}
	}
	return res
}

func part_2(input string) any {
	res := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		a := lines[i]
		b := lines[i+1]
		c := lines[i+2]
		for _, r := range a {
			if strings.Count(b, string(r)) > 0 && strings.Count(c, string(r)) > 0 {
				res += priority(r)
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
