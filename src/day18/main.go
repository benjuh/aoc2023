package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
)

type DigPlan struct {
	dir      string
	distance int
	color    string
}

var plans []DigPlan

var print = fmt.Printf

//go:embed data/data.txt
var input string

func parse(data string) []string {
	lines := common.GetLines(data)
	plans = make([]DigPlan, 0)
	for _, line := range lines {
		var dir string
		var distance int
		var color string
		sections := strings.Split(line, " ")
		dir = sections[0]
		distance, _ = strconv.Atoi(sections[1])
		color = sections[2][2 : len(sections[2])-1]
		plans = append(plans, DigPlan{
			dir:      dir,
			distance: distance,
			color:    color,
		})
	}

	return lines
}

var Delta = map[string]common.Point{
	"U": {X: 0, Y: -1},
	"R": {X: 1, Y: 0},
	"D": {X: 0, Y: 1},
	"L": {X: -1, Y: 0},
}

func part1(data string) string {
	_ = parse(data)

	curr := common.Point{X: 0, Y: 0}
	var edges int
	var vertices = []common.Point{curr}

	for _, plan := range plans {
		curr = common.Point{X: curr.X + Delta[plan.dir].X*plan.distance, Y: curr.Y + Delta[plan.dir].Y*plan.distance}
		vertices = append(vertices, curr)
		edges += plan.distance
	}

	area := util.ShoelaceFormula(vertices)
	area = area + edges/2 + 1
	return fmt.Sprintf("part_1=%d", area)
}

func extract_op(hex string) (string, int) {
	new_hex, _ := strconv.ParseInt(hex[:5], 16, 64)
	dir_int := hex[5] - '0'
	var dir string
	switch dir_int {
	case 0:
		dir = "R"
	case 1:
		dir = "D"
	case 2:
		dir = "L"
	case 3:
		dir = "U"
	}
	return dir, int(new_hex)
}

func part2(data string) string {
	_ = parse(data)

	curr := common.Point{X: 0, Y: 0}
	var edges int
	var vertices = []common.Point{curr}

	for _, plan := range plans {
		dir, distance := extract_op(plan.color)
		curr = common.Point{X: curr.X + Delta[dir].X*distance, Y: curr.Y + Delta[dir].Y*distance}
		vertices = append(vertices, curr)
		edges += distance
	}

	area := util.ShoelaceFormula(vertices)
	area = area + edges/2 + 1
	return fmt.Sprintf("part_2=%v", area)
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

	common.RunDay(18, part_1, part_2, time1, time2)
}
