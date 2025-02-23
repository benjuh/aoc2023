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

type Point struct {
	row int
	col int
}

var galaxies []Point
var empty_rows map[int]struct{}
var empty_cols map[int]struct{}

func parse(data string) []string {
	lines := common.GetLines(data)
	galaxies = make([]Point, 0)
	empty_rows = make(map[int]struct{})
	empty_cols = make(map[int]struct{})
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Point{row: y, col: x})
			}
		}
	}

	return lines
}

func get_empty(lines []string) {
	for i := range len(lines) {
		is_galaxy_in_row := false
		is_galaxy_in_col := false
		for _, galaxy := range galaxies {
			if galaxy.row == i {
				is_galaxy_in_row = true
			}
			if galaxy.col == i {
				is_galaxy_in_col = true
			}
		}
		if !is_galaxy_in_row {
			empty_rows[i] = struct{}{}
		}
		if !is_galaxy_in_col {
			empty_cols[i] = struct{}{}
		}
	}
}

func get_sum_of_distances(expansions int) int {
	sum := 0
	var min int
	var max int

	for i := range len(galaxies) - 1 {
		for j := i + 1; j < len(galaxies); j++ {

			distance := util.ManhattanDistance(galaxies[i].col, galaxies[i].row, galaxies[j].col, galaxies[j].row)
			min, max = util.Order(galaxies[i].row, galaxies[j].row)

			for dr := min + 1; dr < max; dr++ {
				_, is_empty_row := empty_rows[dr]
				if is_empty_row {
					distance += expansions - 1
				}
			}

			min, max = util.Order(galaxies[i].col, galaxies[j].col)
			for dc := min + 1; dc < max; dc++ {
				_, is_empty_col := empty_cols[dc]
				if is_empty_col {
					distance += expansions - 1
				}
			}

			sum += distance
		}
	}
	return sum
}

func part1(data string) string {
	lines := parse(data)
	get_empty(lines)
	sum := get_sum_of_distances(2)

	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	lines := parse(data)
	get_empty(lines)
	sum := get_sum_of_distances(1000000)
	return fmt.Sprintf("part_2=%v", sum)
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

	common.RunDay(1, part_1, part_2, time1, time2)
}
