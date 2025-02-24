package main

import (
	_ "embed"
	"fmt"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
)

const (
	IS_TEST = false
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type Pattern []string

var patterns []Pattern

func parse(data string) []string {
	lines := common.GetLines(data)
	patterns = make([]Pattern, 0)

	pattern := make([]string, 0)
	for i, line := range lines {
		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
			continue
		}
		pattern = append(pattern, line)

		if i == len(lines)-1 {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
		}
	}

	return lines
}

func (p Pattern) has_reflection(smudges int) (int, bool) {
Outer:
	for line := range len(p) - 1 {
		changes := smudges
		for delta := 0; ; delta++ {
			up, down := line-delta, line+delta+1
			if up < 0 || down >= len(p) {
				if changes == 0 {
					return line, true
				}
				continue Outer
			}

			diff := util.Levenshtein(p[up], p[down])
			if diff > changes {
				continue Outer
			}
			changes -= diff
		}
	}
	return 0, false
}

func (p Pattern) Transpose() Pattern {
	newP := Pattern{}
	for c := range len(p[0]) {
		var row string
		for r := range len(p) {
			row += string(p[r][c])
		}
		newP = append(newP, row)
	}
	return newP
}

func part1(data string) string {
	_ = parse(data)

	summarized := 0
	for _, pattern := range patterns {
		if vr, found := pattern.Transpose().has_reflection(0); found {
			summarized += vr + 1
		}
		if hr, found := pattern.has_reflection(0); found {
			summarized += 100 * (hr + 1)
		}
	}

	return fmt.Sprintf("part_1=%v", summarized)
}

func part2(data string) string {
	_ = parse(data)
	summarized := 0
	for _, pattern := range patterns {
		if vr, found := pattern.Transpose().has_reflection(1); found {
			summarized += vr + 1
		}
		if hr, found := pattern.has_reflection(1); found {
			summarized += 100 * (hr + 1)
		}
	}
	return fmt.Sprintf("part_2=%v", summarized)
}

func main() {
	var part_1 string
	var part_2 string
	start1 := time.Now()
	part_1 = part1(input)
	end1 := time.Now()

	start2 := time.Now()
	part_2 = part2(input)
	end2 := time.Now()

	time1 := end1.Sub(start1)
	time2 := end2.Sub(start2)

	common.RunDay(13, part_1, part_2, time1, time2)
}
