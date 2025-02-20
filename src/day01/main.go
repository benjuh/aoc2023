package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	common "github.com/BenjuhminStewart/aoc2023/common"
)

//go:embed data/data.txt
var input string

var string_to_int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parse(data string) []string {
	strings := common.GetLines(data)
	return strings
}

func getCalibration(str string) int {
	first := -1
	last := -1
	for _, char := range str {
		val, _ := strconv.Atoi(string(char))
		if val == 0 {
			continue
		}
		if first == -1 {
			first = val
		}
		last = val
	}
	return (first * 10) + last
}

func getAdvancedCalibration(str string) int {
	var found = make([]int, len(str))
	for key, _ := range string_to_int {
		first_index_of := strings.Index(str, key)
		last_index_of := strings.LastIndex(str, key)
		if first_index_of != -1 {
			found[first_index_of] = string_to_int[key]
		}
		if last_index_of != -1 {
			found[last_index_of] = string_to_int[key]
		}
	}

	for i, char := range str {
		val, _ := strconv.Atoi(string(char))
		if val == 0 {
			continue
		}
		found[i] = val
	}

	first := 0
	last := 0
	for _, val := range found {
		if val == 0 {
			continue
		}
		if first == 0 {
			first = val
		}
		last = val
	}
	return (first * 10) + last
}

func part1(data string) string {
	lines := parse(data)
	sum := 0
	for _, line := range lines {
		sum += getCalibration(line)
	}
	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	lines := parse(data)
	sum := 0
	for _, line := range lines {
		sum += getAdvancedCalibration(line)
	}
	return fmt.Sprintf("part_2=%v", sum)
}

func main() {
	part_1 := part1(input)
	part_2 := part2(input)

	common.RunDay(1, part_1, part_2)
}
