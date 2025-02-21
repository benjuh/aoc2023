package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	common "github.com/benjuh/aoc2023/common"
)

//go:embed data/data.txt
var input string

type Range struct {
	dest   int
	source int
	len    int
}

type SeedRange struct {
	start int
	len   int
}

func (r Range) is_src_in_range(x int) bool {
	return x < r.source+r.len && x >= r.source
}

func (r Range) is_dest_in_range(x int) bool {
	return x < r.dest+r.len && x >= r.dest
}

func (r Range) get_dest(x int) int {
	if !r.is_src_in_range(x) {
		return x
	}
	var diff int
	if x < r.source {
		diff = r.source - x
	} else {
		diff = x - r.source
	}

	return r.dest + int(diff)
}

func (r Range) get_src(x int) int {
	if !r.is_dest_in_range(x) {
		return x
	}
	var diff int
	if x < r.dest {
		diff = r.dest - x
	} else {
		diff = x - r.dest
	}

	return r.source + int(diff)
}

var seeds []int
var seed_range []SeedRange
var seed_to_soil []Range
var soil_to_fertilizer []Range
var fertilizer_to_water []Range
var water_to_light []Range
var light_to_temperature []Range
var temperature_to_humidity []Range
var humidity_to_location []Range

func parse(data string) []string {
	lines := common.GetLines(data)
	current_section := 0
	for _, line := range lines {
		if len(line) == 0 {
			current_section++
			continue
		}
		if current_section != 0 && strings.Contains(line, ":") {
			continue
		}

		var src int
		var dest int
		var length int
		if current_section != 0 {
			split := strings.Split(line, " ")
			dest, _ = strconv.Atoi(split[0])
			src, _ = strconv.Atoi(split[1])
			length, _ = strconv.Atoi(split[2])
		}

		switch current_section {
		case 0:
			seeds_str := strings.Split(line, " ")[1:]
			for _, seed := range seeds_str {
				seed_val, _ := strconv.Atoi(seed)
				seeds = append(seeds, seed_val)
			}
			i := 0
			for i < len(seeds_str) {
				start, _ := strconv.Atoi(seeds_str[i])
				amount, _ := strconv.Atoi(seeds_str[i+1])
				seed_range = append(seed_range, SeedRange{start, amount})
				i += 2
			}

		case 1:
			seed_to_soil = append(seed_to_soil, Range{dest, src, length})
		case 2:
			soil_to_fertilizer = append(soil_to_fertilizer, Range{dest, src, length})
		case 3:
			fertilizer_to_water = append(fertilizer_to_water, Range{dest, src, length})
		case 4:
			water_to_light = append(water_to_light, Range{dest, src, length})
		case 5:
			light_to_temperature = append(light_to_temperature, Range{dest, src, length})
		case 6:
			temperature_to_humidity = append(temperature_to_humidity, Range{dest, src, length})
		case 7:
			humidity_to_location = append(humidity_to_location, Range{dest, src, length})
		}
	}
	return lines
}

func part1(data string) string {
	_ = parse(data)
	lowest_location := math.MaxInt

	for _, seed := range seeds {

		ss_mapping := seed
		for _, r := range seed_to_soil {
			new_mapping := r.get_dest(ss_mapping)
			if new_mapping != ss_mapping {
				ss_mapping = new_mapping
				break
			}
		}

		sf_mapping := ss_mapping
		for _, r := range soil_to_fertilizer {
			new_mapping := r.get_dest(sf_mapping)
			if new_mapping != sf_mapping {
				sf_mapping = new_mapping
				break
			}
		}

		fw_mapping := sf_mapping
		for _, r := range fertilizer_to_water {
			new_mapping := r.get_dest(fw_mapping)
			if new_mapping != fw_mapping {
				fw_mapping = new_mapping
				break
			}
		}

		wl_mapping := fw_mapping
		for _, r := range water_to_light {
			new_mapping := r.get_dest(wl_mapping)
			if new_mapping != wl_mapping {
				wl_mapping = new_mapping
				break
			}
		}

		lt_mapping := wl_mapping
		for _, r := range light_to_temperature {
			new_mapping := r.get_dest(lt_mapping)
			if new_mapping != lt_mapping {
				lt_mapping = new_mapping
				break
			}
		}

		th_mapping := lt_mapping
		for _, r := range temperature_to_humidity {
			new_mapping := r.get_dest(th_mapping)
			if new_mapping != th_mapping {
				th_mapping = new_mapping
				break
			}
		}

		hl_mapping := th_mapping
		for _, r := range humidity_to_location {
			new_mapping := r.get_dest(hl_mapping)
			if new_mapping != hl_mapping {
				hl_mapping = new_mapping
				break
			}
		}

		if hl_mapping < lowest_location {
			lowest_location = hl_mapping
		}

	}
	return fmt.Sprintf("part_1=%v", lowest_location)
}

func solve_backwards(lowest int) bool {
	seed_found := false

	location := lowest
	for _, r := range humidity_to_location {
		new_mapping := r.get_src(location)
		if new_mapping != location {
			location = new_mapping
			break
		}
	}

	humidity := location
	for _, r := range temperature_to_humidity {
		new_mapping := r.get_src(humidity)
		if new_mapping != humidity {
			humidity = new_mapping
			break
		}
	}

	temperature := humidity
	for _, r := range light_to_temperature {
		new_mapping := r.get_src(temperature)
		if new_mapping != temperature {
			temperature = new_mapping
			break
		}
	}

	light := temperature
	for _, r := range water_to_light {
		new_mapping := r.get_src(light)
		if new_mapping != light {
			light = new_mapping
			break
		}
	}

	water := light
	for _, r := range fertilizer_to_water {
		new_mapping := r.get_src(water)
		if new_mapping != water {
			water = new_mapping
			break
		}
	}

	fertilizer := water
	for _, r := range soil_to_fertilizer {
		new_mapping := r.get_src(fertilizer)
		if new_mapping != fertilizer {
			fertilizer = new_mapping
			break
		}
	}

	soil := fertilizer
	for _, r := range seed_to_soil {
		new_mapping := r.get_src(soil)
		if new_mapping != soil {
			soil = new_mapping
			break
		}
	}

	for _, r := range seed_range {
		if soil >= r.start && soil < r.start+r.len {
			lowest = location
			seed_found = true
			break
		}
	}

	return seed_found
}

func part2(data string) string {
	// _ = parse(data)

	i := 0
	last_found := math.MaxInt

	for true {
		found := solve_backwards(i)
		if i == last_found {
			break
		}
		if found {
			last_found = i
			i -= 1000
		}
		if last_found == math.MaxInt {
			i += 1000
			continue
		}
		i += 1
	}

	return fmt.Sprintf("part_2=%v", last_found)
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

	common.RunDay(5, part_1, part_2, time1, time2)
}
