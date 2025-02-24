package main

import (
	_ "embed"
	"fmt"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

const ROLLABLE = 'O'

func parse(data string) [][]string {
	var m [][]string
	lines := common.GetLines(data)

	for _, line := range lines {
		var map_line []string
		for _, char := range line {
			map_line = append(map_line, string(char))
		}
		m = append(m, map_line)
	}

	return m
}

func load(m [][]string) int {
	load := 0
	for y, line := range m {
		for _, char := range line {
			if char == string(ROLLABLE) {
				load += len(m) - y
			}
		}
	}

	return load
}

func part1(data string) string {
	m := parse(data)
	north(m)
	load := load(m)
	return fmt.Sprintf("part_1=%v", load)
}

func north(grid [][]string) {
	for r, row := range grid {
		for c, val := range row {
			if val == "O" {
				for nextRow := r - 1; nextRow >= 0; nextRow-- {
					// can only fall north if nextRow is an empty space
					if grid[nextRow][c] == "." {
						grid[nextRow][c] = "O"
						grid[nextRow+1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func south(grid [][]string) {
	for r := len(grid) - 1; r >= 0; r-- {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val == "O" {
				for nextRow := r + 1; nextRow < len(grid); nextRow++ {
					// can only fall north if nextRow is an empty space
					if grid[nextRow][c] == "." {
						grid[nextRow][c] = "O"
						grid[nextRow-1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func east(grid [][]string) {
	for c := len(grid[0]) - 1; c >= 0; c-- {
		for r := range grid {
			val := grid[r][c]

			if val == "O" {
				for nextCol := c + 1; nextCol < len(grid[0]); nextCol++ {
					// can only fall north if nextCol is an empty space
					if grid[r][nextCol] == "." {
						grid[r][nextCol] = "O"
						grid[r][nextCol-1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func west(grid [][]string) {
	for c := range len(grid[0]) {
		for r := range grid {
			val := grid[r][c]

			if val == "O" {
				for nextCol := c - 1; nextCol >= 0; nextCol-- {
					// can only fall north if nextCol is an empty space
					if grid[r][nextCol] == "." {
						grid[r][nextCol] = "O"
						grid[r][nextCol+1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func Stringify(m [][]string) string {
	ans := ""
	for _, row := range m {
		ans += strings.Join(row, "")
	}
	return ans
}

func part2(data string) string {
	m := parse(data)
	seenStates := map[string]int{}

	cycles := 1000000000
	for c := 0; c < cycles; c++ {
		key := Stringify(m)
		if lastIndex, ok := seenStates[key]; ok {
			cyclePeriod := c - lastIndex
			c = cycles - cyclePeriod
		}
		seenStates[key] = c
		north(m)
		west(m)
		south(m)
		east(m)
	}

	load := load(m)
	return fmt.Sprintf("part_2=%v", load)
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

	common.RunDay(14, part_1, part_2, time1, time2)
}
