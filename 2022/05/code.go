package main

import (
	"aoc-in-go/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

//         [M]     [B]             [N]
// [T]     [H]     [V] [Q]         [H]
// [Q]     [N]     [H] [W] [T]     [Q]
// [V]     [P] [F] [Q] [P] [C]     [R]
// [C]     [D] [T] [N] [N] [L] [S] [J]
// [D] [V] [W] [R] [M] [G] [R] [N] [D]
// [S] [F] [Q] [Q] [F] [F] [F] [Z] [S]
// [N] [M] [F] [D] [R] [C] [W] [T] [M]
//  1   2   3   4   5   6   7   8   9

func main() {
	aoc.Harness(run)
}

func input_piles(input string) [][]rune {
	res := [][]rune{}
	for i := 0; i < 9; i++ {
		res = append(res, []rune{})
	}

	lines := strings.Split(input, "\n")
	for i := len(lines) - 2; i >= 0; i-- {
		r := lines[i]
		for j := 0; j < 9; j++ {
			if len(r) > j*4+1 {
				c := rune(r[j*4+1])
				if c == ' ' {
					continue
				}
				res[j] = append(res[j], c)
			}
		}
	}
	return res
}

func part_1(input string) any {
	res := ""
	blocs := strings.Split(input, "\n\n")
	piles := input_piles(blocs[0])
	for _, r := range strings.Split(blocs[1], "\n") {
		n := utils.Nums(r)
		quantity := n[0]
		from := n[1] - 1
		to := n[2] - 1
		for i := 0; i < quantity; i++ {
			piles[to] = append(piles[to], piles[from][len(piles[from])-1])
			piles[from] = piles[from][:len(piles[from])-1]
		}
	}
	for _, p := range piles {
		if len(p) == 0 {
			continue
		}
		res += string(p[len(p)-1])
	}

	return res
}

func part_2(input string) any {
	res := ""
	blocs := strings.Split(input, "\n\n")
	piles := input_piles(blocs[0])

	for _, r := range strings.Split(blocs[1], "\n") {
		n := utils.Nums(r)
		quantity := n[0]
		from := n[1] - 1
		to := n[2] - 1

		piles[to] = append(piles[to], piles[from][len(piles[from])-quantity:]...)
		piles[from] = piles[from][:len(piles[from])-quantity]

	}
	for _, p := range piles {
		if len(p) == 0 {
			continue
		}
		res += string(p[len(p)-1])
	}

	return res
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}
