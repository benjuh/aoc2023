package main

import (
	_ "embed"
	"fmt"

	common "github.com/benjuh/aoc2023/common"
	"time"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

var directions = common.GetDirections()

var sx int
var sy int

func parse(data string) [][]string {
	lines := common.GetLines(data)
	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, char := range line {
			if char == 'S' {
				sx = x
				sy = y
			}
			grid[y][x] = string(char)
		}
	}

	return grid
}

func part1(data string) string {
	steps := 64
	grid := parse(data)
	queue := map[[2]int]struct{}{}
	queue[[2]int{sx, sy}] = struct{}{}

	i := 0
	for i < steps {
		nq := map[[2]int]struct{}{}
		for coord := range queue {
			for _, delta := range directions {
				next_row := coord[1] + delta.Dy
				next_col := coord[0] + delta.Dx

				if next_row < 0 || next_col < 0 || next_row >= len(grid) || next_col >= len(grid[0]) {
					continue
				}

				if grid[next_row][next_col] == "." || grid[next_row][next_col] == "S" {
					nq[[2]int{next_col, next_row}] = struct{}{}
				}
			}
		}
		queue = nq
		i++
	}

	plots := len(queue)

	return fmt.Sprintf("part_1=%v", plots)
}

func part2(data string) string {
	grid := parse(data)
	steps := 26501365

	grid[sy][sx] = "."

	evens := map[[2]int]struct{}{}
	odds := map[[2]int]struct{}{}

	visited := map[[2]int]struct{}{}

	queue := [][2]int{{sx, sy}}

	results := []int{}

	var plots int

	for s := 0; s < steps && len(results) < 3; s++ {
		active := evens
		if s%2 == 1 {
			active = odds
		}

		nq := [][2]int{}
		for _, coord := range queue {
			active[coord] = struct{}{}
			for _, delta := range directions {
				nr := coord[1] + delta.Dy
				nc := coord[0] + delta.Dx

				next := [2]int{nc, nr}

				mod_nr := ((nr % len(grid)) + len(grid)) % len(grid)
				mod_nc := ((nc % len(grid[0])) + len(grid[0])) % len(grid[0])

				if grid[mod_nr][mod_nc] != "." {
					continue
				}

				if _, ok := visited[next]; ok {
					continue
				}
				visited[next] = struct{}{}

				nq = append(nq, next)
			}
		}
		queue = nq
		if s != 0 && s%131 == 65 {
			results = append(results, len(active))
		}

	}

	a := (results[2] + results[0] - 2*results[1]) / 2
	b := results[1] - results[0] - a
	c := results[0]

	n := steps / len(grid)
	plots = a*n*n + b*n + c

	return fmt.Sprintf("part_2=%v", plots)
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

	common.RunDay(21, part_1, part_2, time1, time2)
}
