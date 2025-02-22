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

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

var network map[string][2]string
var instructions []Direction

func parse(data string) []string {
	lines := common.GetLines(data)
	network = make(map[string][2]string)
	instructions = make([]Direction, 0)

	finished_instructions := false
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if !finished_instructions {
			for _, instruction := range line {
				if instruction == 'L' {
					instructions = append(instructions, LEFT)

				} else if instruction == 'R' {
					instructions = append(instructions, RIGHT)
				} else {
					break
				}
			}
			finished_instructions = true
		} else {
			name := line[0:3]
			left := line[7:10]
			right := line[12:15]

			network[name] = [2]string{left, right}
		}
	}

	return lines
}

func part1(data string) string {
	_ = parse(data)

	current_location := "AAA"
	steps := 0
	for current_location != "ZZZ" {
		for _, instruction := range instructions {
			current_location = network[current_location][instruction]
			steps++
			if current_location == "ZZZ" {
				break
			}
		}
	}

	return fmt.Sprintf("part_1=%v", steps)
}

func part2(data string) string {
	_ = parse(data)
	results := []int{}

	for node := range network {
		if node[2] != 'A' {
			continue
		}

		steps := 0

		for node[2] != 'Z' {
			next_dir := instructions[steps%len(instructions)]
			node = network[node][next_dir]
			steps++
		}
		results = append(results, steps)
	}

	steps := results[0]
	for i := 1; i < len(results); i++ {
		steps = util.Lcm(steps, results[i])
	}

	return fmt.Sprintf("part_2=%v", steps)
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

	common.RunDay(8, part_1, part_2, time1, time2)
}
