package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type History []int

var histories []History

func parse(data string) []string {
	lines := common.GetLines(data)
	histories = make([]History, len(lines))

	for i, line := range lines {
		nums := strings.Split(line, " ")
		values := make(History, len(nums))
		for i, num := range nums {
			val, _ := strconv.Atoi(num)
			values[i] = val
		}
		histories[i] = values
	}

	return lines
}

func is_all_zero(diffs []int) bool {
	for _, val := range diffs {
		if val != 0 {
			return false
		}
	}
	return true
}

func get_diffs(arr []int) []int {
	var diffs []int
	j := 0
	for j < len(arr)-1 {
		diffs = append(diffs, arr[j+1]-arr[j])
		j += 1
	}

	return diffs
}

func part1(data string) string {
	_ = parse(data)
	counters := 0
	for _, history := range histories {
		counters += history[len(history)-1]
		diffs := get_diffs(history)
		counters += diffs[len(diffs)-1]

		for !is_all_zero(diffs) {
			diffs = get_diffs(diffs)
			counters += diffs[len(diffs)-1]
		}

	}
	return fmt.Sprintf("part_1=%v", counters)
}

func extrapolate(firsts []int) int {
	new_firsts := make([]int, len(firsts)-1)
	new_firsts[len(new_firsts)-1] = firsts[len(firsts)-2] - firsts[len(firsts)-1]

	for i := len(new_firsts) - 1; i > 0; i-- {
		new_firsts[i-1] = firsts[i-1] - new_firsts[i]

	}
	return new_firsts[0]
}

func part2(data string) string {
	_ = parse(data)
	counters := 0
	for _, history := range histories {
		var first []int
		counter := 0
		first = append(first, history[0])
		diffs := get_diffs(history)
		first = append(first, diffs[0])
		for !is_all_zero(diffs) {
			diffs = get_diffs(diffs)
			first = append(first, diffs[0])
			counter -= diffs[0]
		}

		counters += extrapolate(first)
	}
	return fmt.Sprintf("part_2=%v", counters)
}

func main() {
	start1 := time.Now()
	part_1 := part1(input)
	end1 := time.Now()

	start2 := time.Now()
	part_2 := part2(input)
	end2 := time.Now()

	time1 := end1.Sub(start1)
	time2 := end2.Sub(start2)

	common.RunDay(9, part_1, part_2, time1, time2)
}
