package main

import (
	_ "embed"
	"fmt"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

const (
	IS_TEST = false
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type BoardInfo struct {
	sx     int
	sy     int
	gx     int
	gy     int
	width  int
	height int
	grid   []string
	slopes map[byte][2]int
}

var info BoardInfo

func parse(data string) []string {
	lines := common.GetLines(data)
	width := len(lines[0])
	height := len(lines)
	sy := 0
	sx := strings.Index(lines[0], ".")
	gy := height - 1
	gx := strings.Index(lines[height-1], ".")

	info = BoardInfo{
		sx:     sx,
		sy:     sy,
		gx:     gx,
		gy:     gy,
		width:  width,
		height: height,
		grid:   lines,
		slopes: map[byte][2]int{
			'>': {1, 0},
			'v': {0, 1},
			'<': {-1, 0},
			'^': {0, -1},
		},
	}

	return lines
}

func backtrack_longest(y, x int, visited map[[2]int]bool, steps int) int {
	if y == info.gy && x == info.gx {
		return steps
	}

	if delta, ok := info.slopes[info.grid[y][x]]; ok {

		nextCoord := [2]int{y + delta[1], x + delta[0]}
		if visited[nextCoord] {
			return 0
		}

		visited[[2]int{y, x}] = true

		result := backtrack_longest(y+delta[1], x+delta[0], visited, steps+1)

		visited[[2]int{y, x}] = false
		return result
	}

	best := 0

	for _, delta := range info.slopes {
		nextRow := y + delta[1]
		nextCol := x + delta[0]

		if nextRow < 0 || nextRow >= info.height ||
			nextCol < 0 || nextCol >= info.width {
			continue
		}

		nextCoord := [2]int{nextRow, nextCol}

		if visited[nextCoord] {
			continue
		}

		if info.grid[nextRow][nextCol] != '#' {
			visited[[2]int{y, x}] = true

			result := backtrack_longest(nextRow, nextCol, visited, steps+1)
			best = max(best, result)

			visited[[2]int{y, x}] = false
		}
	}

	return best
}

func part1(data string) string {
	_ = parse(data)
	best_hike := backtrack_longest(info.sy, info.sx, map[[2]int]bool{}, 0)
	return fmt.Sprintf("part_1=%v", best_hike)
}

type Node struct {
	y, x  int
	edges map[*Node]int
}

func part2(data string) string {
	_ = parse(data)
	nodes := map[[2]int]*Node{}

	for y := range info.height {
		for x := range info.width {
			if info.grid[y][x] == '#' {
				continue
			}
			nodes[[2]int{y, x}] = &Node{
				y:     y,
				x:     x,
				edges: map[*Node]int{},
			}
		}
	}

	for coords, node := range nodes {
		for _, delta := range info.slopes {
			nextCoord := [2]int{coords[0] + delta[1], coords[1] + delta[0]}

			if neighbor, ok := nodes[nextCoord]; ok {
				node.edges[neighbor] = 1
				neighbor.edges[node] = 1
			}
		}
	}

	for _, current := range nodes {
		if len(current.edges) == 2 {
			neighbors := []*Node{}
			summed_weight := 0
			for neighbor := range current.edges {
				neighbors = append(neighbors, neighbor)
				summed_weight += neighbor.edges[current]
			}

			delete(neighbors[0].edges, current)
			delete(neighbors[1].edges, current)
			neighbors[0].edges[neighbors[1]] = summed_weight
			neighbors[1].edges[neighbors[0]] = summed_weight

			delete(nodes, [2]int{current.y, current.x})
		}
	}

	best := backtrackThroughGraph(nodes[[2]int{info.sy, info.sx}],
		map[*Node]bool{}, 0)

	return fmt.Sprintf("part_2=%v", best)
}

func backtrackThroughGraph(current *Node, visited map[*Node]bool, distance int) int {
	if current.y == info.gy {
		return distance
	}

	best := 0
	visited[current] = true

	for neighbor, weight := range current.edges {
		if visited[neighbor] {
			continue
		}

		best = max(best, backtrackThroughGraph(neighbor, visited, distance+weight))
	}
	visited[current] = false

	return best
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

	common.RunDay(23, part_1, part_2, time1, time2)
}
