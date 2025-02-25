package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type Rating [4]int

type Condition int

const (
	LESS_THAN Condition = iota + 1
	GREATER_THAN
)

type Rule struct {
	index      int
	condition  Condition
	comparison int
	goes_to    string
}

type Rules struct {
	rules      []Rule
	default_to string
}

var workflows map[string]Rules
var ratings []Rating

var index_of = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

func parse_rules(rules []string) Rules {
	parsed_rules := make([]Rule, 0)

	var default_to string
	for i, rule := range rules {
		if i == len(rules)-1 {
			default_to = rule
		}
		before, after, index_less_than := strings.Cut(rule, "<")
		if index_less_than {
			index := strings.Index(after, ":")
			x, _ := index_of[before]
			y, _ := strconv.Atoi(after[:index])
			rule := Rule{x, LESS_THAN, y, after[index+1:]}
			parsed_rules = append(parsed_rules, rule)
		}

		before, after, index_greater_than := strings.Cut(rule, ">")
		if index_greater_than {
			index := strings.Index(after, ":")
			x, _ := index_of[before]
			y, _ := strconv.Atoi(after[:index])
			rule := Rule{x, GREATER_THAN, y, after[index+1:]}
			parsed_rules = append(parsed_rules, rule)
		}

	}
	return Rules{parsed_rules, default_to}
}

func parse(data string) []string {
	lines := common.GetLines(data)
	workflows = make(map[string]Rules)
	ratings = make([]Rating, 0)
	i := 0
	for _, line := range lines {
		if len(line) == 0 {
			i += 1
			break
		}
		startIndex := strings.Index(line, "{")
		endIndex := strings.Index(line, "}")
		label := line[0:startIndex]

		inner := line[startIndex+1 : endIndex]
		rules := parse_rules(strings.Split(inner, ","))

		workflows[label] = rules
		i += 1
	}

	for i < len(lines) {

		line := lines[i]
		usable := line[1 : len(line)-1]

		ratings_strs := strings.Split(usable, ",")
		ratings_x, _ := strconv.Atoi(ratings_strs[0][2:])
		ratings_m, _ := strconv.Atoi(ratings_strs[1][2:])
		ratings_a, _ := strconv.Atoi(ratings_strs[2][2:])
		ratings_s, _ := strconv.Atoi(ratings_strs[3][2:])

		arr := [4]int{ratings_x, ratings_m, ratings_a, ratings_s}
		ratings = append(ratings, arr)
		i += 1
	}

	return lines
}

func sum_ratings(ratings [4]int) int {
	return ratings[0] + ratings[1] + ratings[2] + ratings[3]
}

func follow_workflow(ratings [4]int, label string) int {
	if label == "A" {
		return sum_ratings(ratings)
	}
	if label == "R" {
		return 0
	}
	rules := workflows[label]
	var next_label string
	for _, rule := range rules.rules {
		switch rule.condition {
		case LESS_THAN:
			if ratings[rule.index] < rule.comparison {
				next_label = rule.goes_to
				return follow_workflow(ratings, next_label)
			}
		case GREATER_THAN:
			if ratings[rule.index] > rule.comparison {
				next_label = rule.goes_to
				return follow_workflow(ratings, next_label)
			}
		}
	}
	return follow_workflow(ratings, rules.default_to)
}

func part1(data string) string {
	_ = parse(data)
	ratings_sum := 0
	for _, rating := range ratings {
		ratings_sum += follow_workflow(rating, "in")
	}
	return fmt.Sprintf("part_1=%v", ratings_sum)
}

type Range struct {
	start int
	end   int
}

func count_combinations(key string, ranges []Range) int {
	if key == "A" {
		return ranges_product(ranges)
	}
	if key == "R" {
		return 0
	}

	result := 0
	workflow := workflows[key]
	for _, rule := range workflow.rules {
		new_ranges := make([]Range, len(ranges))
		copy(new_ranges, ranges[:])

		switch rule.condition {
		case LESS_THAN:
			new_ranges[rule.index].end = rule.comparison - 1
			ranges[rule.index].start = rule.comparison
			result += count_combinations(rule.goes_to, new_ranges)
		case GREATER_THAN:
			new_ranges[rule.index].start = rule.comparison + 1
			ranges[rule.index].end = rule.comparison
			result += count_combinations(rule.goes_to, new_ranges)
		}
	}
	result += count_combinations(workflow.default_to, ranges)

	return result

}

func ranges_product(ranges []Range) int {
	product := 1
	for _, r := range ranges {
		product *= r.end - r.start + 1
	}
	return product
}

func part2(data string) string {
	_ = data
	sum := 0
	sum += count_combinations("in", []Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}})
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

	common.RunDay(19, part_1, part_2, time1, time2)
}
