package main

import (
	_ "embed"
	"fmt"

	common "github.com/BenjuhminStewart/aoc2023/common"
)

const (
	IS_TEST = true
)

//go:embed data/data.txt
var input string

//go:embed data/test_data.txt
var testInput string

func parse(data string) []string {
	lines := common.GetLines(data)
	common.PrintLines(lines)

	return lines
}

func part1(data string) string {
	_ = parse(data)
	return fmt.Sprintf("part_1=%v", 0)
}

func part2(data string) string {
	// _ = parse(data)
	return fmt.Sprintf("part_2=%v", 0)
}

func main() {
	var part_1 string
	var part_2 string

	if IS_TEST {
		part_1 = part1(testInput)
		part_2 = part2(testInput)
	} else {
		part_1 = part1(input)
		part_2 = part2(input)
	}

	common.RunDay(1, part_1, part_2)

	// TESTING
	common.PrintExpected(1, 0, IS_TEST)
	common.PrintExpected(2, 0, IS_TEST)
}
