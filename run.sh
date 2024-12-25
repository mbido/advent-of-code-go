#!/bin/bash
set -euf -o pipefail

# functions
function echogrey() {
	echo -e "\033[0;90m$1\033[0m"
}

function template() {
	cat <<EOF
package main

import (
	u "aoc-in-go/utils"

	"github.com/jpillora/puzzler/harness/aoc"
)

func useImports() {
	_ = u.PrintHello
}

func main() {
	aoc.Harness(run)
}

func part_1(input string) any {
	return 42
}

func part_2(input string) any {
	return "not implemented"
}

func run(part2 bool, input string) any {
	if part2 {
		return part_2(input)
	}
	return part_1(input)
}

EOF
}

# two args YEAR and DAY
YEAR="${1:-}"
DAY="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
	echogrey "Created directory $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
	template >"$DIR/code.go"
	echogrey "Created file code.go"
fi
# go run
cd "$DIR" && go run code.go
