package main

import (
	_ "embed"
	"fmt"

	common "github.com/BenjuhminStewart/aoc2023/common"
)

//go:embed data/data.txt
var input string

//go:embed data/test_data.txt
var testInput string

func parse(data string) []string {
	strings := common.GetLines(data)
	return strings
}

func part1() string {
	return fmt.Sprintf("part_1=%v", 0)
}

func part2() string {
	return fmt.Sprintf("part_2=%v", 0)
}

func main() {

	lines := parse(testInput)
	common.PrintLines(lines)

	part_1 := part1()
	part_2 := part2()

	common.RunDay(1, part_1, part_2)
}
