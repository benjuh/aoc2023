package main

import (
	_ "embed"
	"fmt"
	"strconv"

	common "github.com/benjuh/aoc2023/common"
)

//go:embed data/data.txt
var input string

var schematic [][]rune
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func parse(data string) []string {
	lines := common.GetLines(data)
	width := len(lines[0])
	height := len(lines)
	schematic = make([][]rune, height)
	for i, line := range lines {
		row := make([]rune, width)
		for i, c := range line {
			row[i] = c
		}
		schematic[i] = row
	}
	return lines
}

type Point struct {
	i, j int
}

type Range struct {
	i, start, end int
	final_val     int
}

func get_range_map() map[Range]bool {

	range_map := map[Range]bool{}
	visited := map[Point]bool{}

	for i, row := range schematic {
		for j, _ := range row {
			if visited[Point{i, j}] {
				continue
			}
			start := j
			end := j
			_, err := strconv.Atoi(string(schematic[i][j]))
			for err == nil && end < len(schematic[i]) {
				end += 1
				if end >= len(schematic[i]) {
					break
				}
				_, err = strconv.Atoi(string(schematic[i][start : end+1]))
			}
			final_val, new_err := strconv.Atoi(string(schematic[i][start:end]))
			if new_err != nil {
				continue
			}
			range_map[Range{i, start, end - 1, final_val}] = true
			for start < end {
				visited[Point{i, start}] = true
				start += 1
			}
		}
	}
	return range_map
}

func part1(data string) string {
	_ = parse(data)

	range_map := get_range_map()

	sum := 0

	for r, _ := range range_map {
		row := r.i
		start := r.start
		end := r.end

		found_symbol_adjacent := false
		for start <= end {
			for _, d := range directions {
				if start < 0 || start >= len(schematic[row]) {
					continue
				}
				new_i := row + d[0]
				new_j := start + d[1]
				if new_i < 0 || new_i >= len(schematic) || new_j < 0 || new_j >= len(schematic[new_i]) {
					continue
				}
				_, err := strconv.Atoi(string(schematic[new_i][new_j]))
				if err != nil && schematic[new_i][new_j] != '.' {
					found_symbol_adjacent = true
					break
				}
			}
			if found_symbol_adjacent {
				break
			}
			start += 1
		}
		if found_symbol_adjacent {
			sum += r.final_val
		}
	}

	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	_ = parse(data)

	range_map := get_range_map()
	potential_gears := map[Point][]int{}

	sum := 0

	for r, _ := range range_map {
		row := r.i
		start := r.start
		end := r.end

		found_symbol_adjacent := false
		for start <= end {
			for _, d := range directions {
				if start < 0 || start >= len(schematic[row]) {
					continue
				}
				new_i := row + d[0]
				new_j := start + d[1]
				if new_i < 0 || new_i >= len(schematic) || new_j < 0 || new_j >= len(schematic[new_i]) {
					continue
				}
				_, err := strconv.Atoi(string(schematic[new_i][new_j]))
				if err != nil && schematic[new_i][new_j] == '*' {
					found_symbol_adjacent = true
					potential_gears[Point{new_i, new_j}] = append(potential_gears[Point{new_i, new_j}], r.final_val)
					break
				}
			}
			if found_symbol_adjacent {
				break
			}
			start += 1
		}
	}
	for _, arr := range potential_gears {
		if len(arr) == 2 {
			sum += arr[0] * arr[1]
		}
	}
	return fmt.Sprintf("part_2=%v", sum)
}

func main() {
	part_1 := part1(input)
	part_2 := part2(input)

	common.RunDay(3, part_1, part_2)
}
