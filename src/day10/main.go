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

var sx int
var sy int

type Direction [2]int

var visited map[Point]int

func parse(data string) []string {
	lines := common.GetLines(data)
	visited = make(map[Point]int)
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				sx = x
				sy = y
				break
			}
		}
	}

	return lines
}

type Point struct {
	x int
	y int
}

func (p Point) Equals(p2 Point) bool {
	return p.x == p2.x && p.y == p2.y
}

var steps_map = map[Point]int{}

func part1(data string) string {
	lines := parse(data)
	x := sx
	y := sy
	start := Point{x, y}

	visited = map[Point]int{start: 0}
	notChecked := []Point{start}

	maxDist := 0
	for len(notChecked) > 0 {
		current := notChecked[0]
		notChecked = notChecked[1:]
		next := nextPoints(lines, current)
		for _, point := range next {
			if _, found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDist = max(maxDist, visited[current]+1)
				notChecked = append(notChecked, point)
			}
		}
	}
	return fmt.Sprintf("part_1=%v", len(visited)/2)
}

func part2(data string) string {
	lines := common.GetLines(data)
	countInside := 0
	for y, row := range lines {
		inside := false
		for x := range len(row) {
			tile := row[x]
			point := Point{x, y}
			if tile == 'S' {
				tile = findStartTile(point, lines)
			}
			if _, part := visited[point]; part {
				if tile == '|' || tile == 'L' || tile == 'J' {
					inside = !inside
				}
			} else if inside {
				countInside++
			}
		}
	}

	return fmt.Sprintf("part_2=%v", countInside)
}

func findStartTile(start Point, grid []string) byte {
	points := nextPoints(grid, start)
	minx, maxx, miny, maxy := min(points[0].x, points[1].x), max(points[0].x, points[1].x), min(points[0].y, points[1].y), max(points[0].y, points[1].y)
	if points[0].x == points[1].x {
		return '|'
	} else if points[0].y == points[1].y {
		return '-'
	} else if minx < start.x && miny < start.y {
		return 'J'
	} else if maxx > start.x && maxy > start.y {
		return 'F'
	} else if maxx > start.x && miny < start.y {
		return 'L'
	} else if minx < start.x && maxy > start.y {
		return '7'
	}
	return '.'
}

func nextPoints(grid []string, p Point) []Point {
	points := []Point{}
	switch grid[p.y][p.x] {
	case '|':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case '-':
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case 'L':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case 'J':
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case '7':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case 'F':
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case '.':
	case 'S':
		down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		if down == '|' || down == 'L' || down == 'J' {
			points = append(points, Point{p.x, p.y + 1})
		}
		if right == '-' || right == '7' || right == 'J' {
			points = append(points, Point{p.x + 1, p.y})
		}
		if up == '|' || up == '7' || up == 'F' {
			points = append(points, Point{p.x, p.y - 1})
		}
		if left == '-' || left == 'L' || left == 'F' {
			points = append(points, Point{p.x - 1, p.y})
		}
	}
	return points
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

	common.RunDay(10, part_1, part_2, time1, time2)
}
