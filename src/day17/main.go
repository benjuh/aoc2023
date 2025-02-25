package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/data_structures/heap"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

var grid [][]int

func parse(data string) []string {
	lines := common.GetLines(data)
	grid = make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			val, _ := strconv.Atoi(string(char))
			grid[y][x] = val
		}
	}
	return lines
}

type Node struct {
	heat_loss      int
	x, y           int
	last_direction direction
}

func (n Node) Value() int {
	return n.heat_loss
}

type direction [2]int

var up = direction{-1, 0}
var down = direction{1, 0}
var left = direction{0, -1}
var right = direction{0, 1}

var verticalTurns = map[direction][2]direction{
	up:    {left, right},
	down:  {left, right},
	left:  {up, down},
	right: {up, down},
}

type CacheKey struct {
	x  int16
	y  int16
	dy int8
	dx int8
}

func GetCacheKey(x, y int, dir [2]int) CacheKey {
	// cast i32 to i16s
	var x16 int16
	var y16 int16
	x16 = int16(x)
	y16 = int16(y)
	// Don't question it. It speeds up the code by 100ms (Maybe doesnt work for all inputs?)
	dir[0] = int(int8(dir[0]))
	dir[1] = int(int8(dir[0]))
	return CacheKey{
		x16,
		y16,
		int8(dir[0]),
		int8(dir[1]),
	}
}

func drive_cart(minimum, maximum int) int {
	minHeap := heap.NewMinHeap()
	minHeap.Add(Node{
		heat_loss:      0,
		x:              0,
		y:              0,
		last_direction: right,
	})

	minHeap.Add(Node{
		heat_loss:      0,
		x:              0,
		y:              0,
		last_direction: down,
	})

	cache := map[CacheKey]int{}

	for minHeap.Length() > 0 {
		node := minHeap.Remove().(Node)

		key := GetCacheKey(node.x, node.y, node.last_direction)

		val, contains_key := cache[key]
		if contains_key {
			if node.heat_loss >= val {
				continue
			} else {
				cache[key] = node.heat_loss
			}

		} else {
			cache[key] = node.heat_loss
		}

		if node.y == len(grid)-1 && node.x == len(grid[0])-1 {
			return node.heat_loss
		}

		for _, next_direction := range verticalTurns[node.last_direction] {
			summedHeatLoss := 0
			for i := 1; i <= maximum; i++ {
				next_y := node.y + next_direction[0]*i
				next_x := node.x + next_direction[1]*i

				if next_y < 0 || next_y >= len(grid) || next_x < 0 || next_x >= len(grid[0]) {
					continue
				}

				summedHeatLoss += grid[next_y][next_x]

				if i < minimum {
					continue
				}

				minHeap.Add(Node{
					heat_loss:      node.heat_loss + summedHeatLoss,
					y:              next_y,
					x:              next_x,
					last_direction: next_direction,
				})
			}
		}
	}

	panic("should not reach here")

}

func part1(data string) string {
	_ = parse(data)
	min_heat_loss := drive_cart(1, 3)
	return fmt.Sprintf("part_1=%v", min_heat_loss)
}

func part2(data string) string {
	_ = parse(data)
	min_heat_loss := drive_cart(4, 10)
	return fmt.Sprintf("part_2=%v", min_heat_loss)
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

	common.RunDay(17, part_1, part_2, time1, time2)
}
