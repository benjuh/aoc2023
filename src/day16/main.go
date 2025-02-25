package main

import (
	_ "embed"
	"fmt"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

const (
	IS_TEST = false
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type Point struct {
	x int
	y int
}

type State struct {
	x  int
	y  int
	dx int
	dy int
}

func (s State) fields() (int, int, int, int) {
	return s.x, s.y, s.dx, s.dy
}

var grid []string
var width int
var height int

func parse(data string) []string {
	lines := common.GetLines(data)
	grid = lines
	width = len(lines[0])
	height = len(lines)
	return lines
}

func bfs(start State, ignored map[State]struct{}) int {
	_, contains_start := ignored[start]
	if contains_start {
		return 0
	}

	current := make([]State, 0)
	current = append(current, start)
	visited := make(map[State]struct{})
	warm := make(map[Point]struct{})
	for len(current) > 0 {
		next := make([]State, 0)
		for _, state := range current {
			x, y, dx, dy := state.fields()
			x += dx
			y += dy
			if x < 0 || x >= width || y < 0 || y >= height {
				ignored_state := State{x: x, y: y, dx: -dx, dy: -dy}
				ignored[ignored_state] = struct{}{}
				continue
			}

			new_state := State{x, y, dx, dy}
			if _, contains := visited[new_state]; contains {
				continue
			}

			visited[new_state] = struct{}{}
			warm[Point{x, y}] = struct{}{}
			t := grid[y][x]

			if t == '.' || (t == '|' && dx == 0) || (t == '-' && dy == 0) {
				next = append(next, new_state)
			} else if t == '|' && dx != 0 {
				up := State{x, y, 0, 1}
				down := State{x, y, 0, -1}
				next = append(next, up, down)
			} else if t == '-' && dy != 0 {
				left := State{x, y, -1, 0}
				right := State{x, y, 1, 0}
				next = append(next, left, right)
			} else if t == '/' {
				next = append(next, State{x, y, -dy, -dx})
			} else if t == '\\' {
				next = append(next, State{x, y, dy, dx})
			}
		}
		current = next
	}

	return len(warm)
}

func part1(data string) string {
	_ = parse(data)

	ignored := make(map[State]struct{})
	start := State{x: -1, y: 0, dx: 1, dy: 0}

	energized := bfs(start, ignored)
	return fmt.Sprintf("part_1=%v", energized)
}

func part2(data string) string {
	_ = parse(data)

	ignored := make(map[State]struct{})
	energized := 0
	for x := range width {
		start_1 := State{x: x, y: -1, dx: 0, dy: 1}
		energized = max(energized, bfs(start_1, ignored))
		start_2 := State{x: x, y: height, dx: 0, dy: -1}
		energized = max(energized, bfs(start_2, ignored))
	}
	for y := range height {
		start_1 := State{x: -1, y: y, dx: 1, dy: 0}
		start_2 := State{x: width, y: y, dx: -1, dy: 0}
		energized = max(energized, bfs(start_1, ignored))
		energized = max(energized, bfs(start_2, ignored))
	}
	return fmt.Sprintf("part_2=%v", energized)
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

	common.RunDay(16, part_1, part_2, time1, time2)
}
