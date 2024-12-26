package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	// if len(input) > 1000 {
	// 	return 42
	// }
	for i := range input[:len(input)-4] {
		// fmt.Println(i, string(input[i]), string(input[i+1]), string(input[i+2]), string(input[i+3]))
		seen := ""
		ok := true
		for j := 0; j < 4; j++ {
			c := input[i+j]
			if strings.Contains(seen, string(c)) {
				ok = false
				break
			}
			seen += string(input[i+j])
		}
		if ok {
			return i + 4
		}
	}
	return -1
}

func part_2(input string) any {
	for i := range input[:len(input)-14] {
		seen := ""
		ok := true
		for j := 0; j < 14; j++ {
			c := input[i+j]
			if strings.Contains(seen, string(c)) {
				ok = false
				break
			}
			seen += string(input[i+j])
		}
		if ok {
			return i + 14
		}
	}
	return -1
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
