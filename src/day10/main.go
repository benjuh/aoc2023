package main

import (
	_ "embed"
	"fmt"
	"maps"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

const (
	IS_TEST = false
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

//go:embed data/test_data.txt
var testInput string

//go:embed data/test_data_2.txt
var testInput2 string

//go:embed data/test_data_2_complex.txt
var testInput2Complex string

type Point struct {
	x int
	y int
}

var Pipe map[Point]string
var Start Point
var grid [][]string

func parse(data string) []string {
	lines := common.GetLines(data)
	grid = make([][]string, len(lines))
	Pipe = make(map[Point]string)
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, char := range line {
			grid[y][x] = string(char)
			if char == '.' {
				continue
			}

			p := Point{x, y}
			if char == 'S' {
				Start = p
			}
			Pipe[p] = string(char)
		}
	}

	return lines
}

func (p Point) steps_between(p2 Point, steps int, steps_from_start map[Point]int, visited map[Point]struct{}) {
	_, is_valid_point := Pipe[p2]
	if !is_valid_point {
		return
	}
	_, seen_in_visited := visited[p2]
	if seen_in_visited {
		if steps_from_start[p2] < steps {
			return
		}
	}
	steps_from_start[p2] = steps
	visited[p2] = struct{}{}

	up := Point{p2.x, p2.y - 1}
	down := Point{p2.x, p2.y + 1}
	left := Point{p2.x - 1, p2.y}
	right := Point{p2.x + 1, p2.y}

	switch Pipe[p2] {
	case "|":
		Start.steps_between(up, steps+1, steps_from_start, visited)
		Start.steps_between(down, steps+1, steps_from_start, visited)
	case "-":
		Start.steps_between(left, steps+1, steps_from_start, visited)
		Start.steps_between(right, steps+1, steps_from_start, visited)
	case "L":
		Start.steps_between(up, steps+1, steps_from_start, visited)
		Start.steps_between(right, steps+1, steps_from_start, visited)
	case "F":
		Start.steps_between(down, steps+1, steps_from_start, visited)
		Start.steps_between(right, steps+1, steps_from_start, visited)
	case "J":
		Start.steps_between(left, steps+1, steps_from_start, visited)
		Start.steps_between(up, steps+1, steps_from_start, visited)
	case "7":
		Start.steps_between(left, steps+1, steps_from_start, visited)
		Start.steps_between(down, steps+1, steps_from_start, visited)
	case "S":
		_, contains_up := Pipe[up]
		if contains_up && (Pipe[up] == "7" || Pipe[up] == "F" || Pipe[up] == "|") {
			Start.steps_between(up, steps+1, steps_from_start, visited)
		}
		_, contains_down := Pipe[down]
		if contains_down && (Pipe[down] == "L" || Pipe[down] == "J" || Pipe[down] == "|") {
			Start.steps_between(down, steps+1, steps_from_start, visited)
		}
		_, contains_left := Pipe[left]
		if contains_left && (Pipe[left] == "-" || Pipe[left] == "F" || Pipe[left] == "L") {
			Start.steps_between(left, steps+1, steps_from_start, visited)
		}
		_, contains_right := Pipe[right]
		if contains_right && (Pipe[right] == "-" || Pipe[right] == "J" || Pipe[right] == "7") {
			Start.steps_between(right, steps+1, steps_from_start, visited)
		}
	}

}

var MainPipeline map[Point]struct{}

func part1(data string) (string, map[Point]struct{}) {
	_ = parse(data)
	steps_from_start := make(map[Point]int)
	visited := make(map[Point]struct{})
	MainPipeline = make(map[Point]struct{})
	Start.steps_between(Start, 0, steps_from_start, visited)

	furthest := 0
	for _, v := range steps_from_start {
		if v > furthest {
			furthest = v
		}
	}

	return fmt.Sprintf("part_1=%v", furthest), visited
}

func print_grid() {
	for _, line := range grid {
		for _, char := range line {
			fmt.Printf("%v", char)
		}
		fmt.Printf("\n")
	}
}

func get_free_points() map[Point]struct{} {
	free_points := map[Point]struct{}{}
	for y, line := range grid {
		for x, char := range line {
			switch x {
			case 0:
				if char == "." {
					free_points[Point{x, y}] = struct{}{}
				}
			case len(line) - 1:
				if char == "." {
					free_points[Point{x, y}] = struct{}{}
				}
			default:
				if y > 0 && y < len(grid)-1 {
					continue
				}
				if char == "." {
					free_points[Point{x, y}] = struct{}{}
				}
			}
		}
	}

	return free_points
}

func traverse(src Point, current Point, visited map[Point]struct{}, free_points map[Point]struct{}) {
	if current.x < 0 || current.x >= len(grid[0]) || current.y < 0 || current.y >= len(grid) {
		return
	}
	_, in_main_pipeline := MainPipeline[current]
	if in_main_pipeline {
		return
	}
	_, is_visited := visited[current]
	if is_visited {
		return
	}

	visited[current] = struct{}{}

	_, is_free_point := free_points[current]
	if is_free_point {
		for k := range maps.Keys(visited) {
			free_points[k] = struct{}{}
		}
		return
	}
	up := Point{current.x, current.y - 1}
	down := Point{current.x, current.y + 1}
	left := Point{current.x - 1, current.y}
	right := Point{current.x + 1, current.y}
	traverse(src, up, visited, free_points)
	traverse(src, down, visited, free_points)
	traverse(src, left, visited, free_points)
	traverse(src, right, visited, free_points)
}

func part2(data string) string {
	_ = parse(data)
	print_grid()
	print("[%v] main pipeline\n", len(MainPipeline))
	free_points := get_free_points()

	for y, line := range grid {
		for x, _ := range line {
			_, is_free_point := free_points[Point{x, y}]
			if is_free_point {
				continue
			}
			visited := map[Point]struct{}{}
			source := Point{x, y}
			up := Point{x, y - 1}
			down := Point{x, y + 1}
			left := Point{x - 1, y}
			right := Point{x + 1, y}
			traverse(source, up, visited, free_points)
			traverse(source, down, visited, free_points)
			traverse(source, left, visited, free_points)
			traverse(source, right, visited, free_points)
		}
	}

	enclosed := 0
	for y, line := range grid {
		for x, char := range line {
			if char != "." {
				continue
			}
			_, is_free_point := free_points[Point{x, y}]
			if !is_free_point {
				enclosed++
			}
		}
	}

	return fmt.Sprintf("part_2=%v", enclosed)
}

func main() {
	var part_1 string
	var part_2 string

	if IS_TEST {
		start1 := time.Now()
		part_1, _ = part1(testInput)
		end1 := time.Now()

		start2 := time.Now()
		part_2 = part2(testInput2Complex)
		end2 := time.Now()

		time1 := end1.Sub(start1)
		time2 := end2.Sub(start2)

		common.RunDay(10, part_1, part_2, time1, time2)
	} else {
		start1 := time.Now()
		part_1, _ = part1(input)
		end1 := time.Now()

		start2 := time.Now()
		part_2 = part2(input)
		end2 := time.Now()

		time1 := end1.Sub(start1)
		time2 := end2.Sub(start2)

		common.RunDay(10, part_1, part_2, time1, time2)
	}

	// TESTING
	common.PrintExpected(1, 4, IS_TEST)
	common.PrintExpected(2, 4, IS_TEST)
}
