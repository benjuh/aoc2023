package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

func toList(s string) []int {
	list := strings.Split(s, ",")
	groups := make([]int, 0, len(list))
	for _, val := range list {
		n, _ := strconv.ParseUint(val, 10, 8)
		groups = append(groups, int(n))
	}
	return groups
}

func parse(data string) []string {
	lines := common.GetLines(data)
	return lines
}

func part1(data string) string {
	lines := parse(data)
	sum := 0
	for _, line := range lines {
		sections := strings.Split(line, " ")
		condition := sections[0]
		groups := toList(sections[1])
		sum += Cache{}.cachedArrangements(condition, groups)
	}

	return fmt.Sprintf("part_1=%v", sum)
}

type Cache map[string]int

func (c Cache) cachedArrangements(s string, num []int) int {
	s = strings.Trim(s, ".")

	key := s + fmt.Sprintf("%v", num)

	if val, ok := c[key]; ok {
		return val
	}

	result := c.arrangements(s, num)

	c[key] = result
	return result
}

func (c Cache) arrangements(s string, num []int) int {
	if s == "" {
		if len(num) == 0 {
			return 1
		}
		return 0
	}

	if len(num) == 0 {
		if strings.Contains(s, "#") {
			return 0
		}
		return 1
	}

	sum := 0

	if s[0] == '?' {
		sum += c.cachedArrangements(s[1:], num)
		s = "#" + s[1:]
	}

	if len(s) < num[0] {
		return sum
	}

	if strings.ContainsRune(s[:num[0]], '.') {
		return sum
	}

	if len(s) > num[0] {
		switch s[num[0]] {
		case '#':
			return sum
		case '?':
			s = s[:num[0]] + "." + s[num[0]+1:]
		}
	}

	return sum + c.cachedArrangements(s[num[0]:], num[1:])
}

func part2(data string) string {
	lines := parse(data)
	sum := 0
	for _, line := range lines {
		sections := strings.Split(line, " ")
		condition := sections[0]
		groups := sections[1]

		conditions := strings.Repeat(condition+"?", 5)
		conditions = conditions[:len(conditions)-1]

		groups_str := strings.Repeat(groups+",", 5)
		groups_str = groups_str[:len(groups_str)-1]

		sum += Cache{}.cachedArrangements(conditions, toList(groups_str))

	}

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

	common.RunDay(12, part_1, part_2, time1, time2)
}
