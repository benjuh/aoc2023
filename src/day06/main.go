package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

//go:embed data/data.txt
var input string

type Race struct {
	time     int
	distance int
}

var races []Race

func parse(data string) []string {
	lines := common.GetLines(data)
	races = make([]Race, 0)
	time := make([]int, 0)
	distance := make([]int, 0)

	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "Time:" {
			for _, field := range fields[1:] {
				num, _ := strconv.Atoi(field)
				time = append(time, num)
			}
		} else if fields[0] == "Distance:" {
			for _, field := range fields[1:] {
				num, _ := strconv.Atoi(field)
				distance = append(distance, num)
			}
		}
	}

	for i := 0; i < len(time); i++ {
		races = append(races, Race{time[i], distance[i]})
	}

	return lines
}
func get_furthest_left(race Race) int {
	left := 1
	right := race.time
	for left <= right {
		mid := (left + right) / 2
		dist := mid * (race.time - mid)
		if dist > race.distance {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func get_furthest_right(race Race) int {
	left := 1
	right := race.time
	for left <= right {
		mid := (left + right) / 2
		dist := mid * (race.time - mid)
		if dist > race.distance {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func part1(data string) string {
	_ = parse(data)

	ways_to_win := 0
	for i, race := range races {
		furthest_left := get_furthest_left(race)
		furthest_right := get_furthest_right(race)

		if i == 0 {
			ways_to_win += furthest_right - furthest_left + 1
		} else {
			ways_to_win *= furthest_right - furthest_left + 1
		}

	}

	return fmt.Sprintf("part_1=%v", ways_to_win)
}

func part2(data string) string {
	_ = parse(data)
	time := ""
	distance := ""
	for _, race := range races {
		time += fmt.Sprintf("%v", race.time)
		distance += fmt.Sprintf("%v", race.distance)
	}
	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	one_race := Race{t, d}
	furthest_left := get_furthest_left(one_race)
	furthest_right := get_furthest_right(one_race)
	ways_to_win := furthest_right - furthest_left + 1
	return fmt.Sprintf("part_2=%v", ways_to_win)
}

func main() {
	start1 := time.Now()
	part_1 := part1(input)
	end1 := time.Now()
	time1 := end1.Sub(start1)

	start2 := time.Now()
	part_2 := part2(input)
	end2 := time.Now()
	time2 := end2.Sub(start2)

	common.RunDay(6, part_1, part_2, time1, time2)
}
