package common

import (
	"fmt"
	"strings"
	"time"
)

var print = fmt.Printf

const (
	OKAY_TIME = 400
	BAD_TIME  = 750
)

func GetLines(input string) []string {
	input = strings.TrimSpace(input)
	strings := strings.Split(input, "\n")
	return strings
}

func PrintLines(input []string) {
	sep := ""
	maxLen := 0
	for _, line := range input {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	sep = strings.Repeat("-", maxLen/2-2)
	sep += "DATA"
	sep += strings.Repeat("-", maxLen/2-2)
	fmt.Printf("%s\n\n", sep)
	for _, line := range input {
		fmt.Printf("%s\n", line)
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", maxLen))
}

func PrintHeader(day int) {
	bold := "\033[1m"
	grey := "\033[90m"
	reset := "\033[0m"
	fmt.Printf("\n%s[%s %sDay %v %s]%s\n", grey, reset, bold, day, grey, reset)
}

func PrintAnswer(part1, part2 string, time1 time.Duration, time2 time.Duration) {
	bold := "\033[1m"
	green := "\033[32m"
	orange := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	var color1 string
	var color2 string
	if time1.Milliseconds() < OKAY_TIME {
		color1 = green
	} else if time1.Milliseconds() < BAD_TIME {
		color1 = orange
	} else {
		color1 = red
	}

	if time2.Milliseconds() < OKAY_TIME {
		color2 = green
	} else if time2.Milliseconds() < BAD_TIME {
		color2 = orange
	} else {
		color2 = red
	}

	fmt.Printf("%s%-10s  %s%s%s%s\n", color1, time1, reset, bold, part1, reset)
	fmt.Printf("%s%-10s  %s%s%s%s\n", color2, time2, reset, bold, part2, reset)
}

func PrintExpected(part int, expected any, is_test bool) {
	if !is_test {
		return
	}
	fmt.Printf("\n[Part %v] Expected: %v\n", part, expected)
}

func RunDay(day int, part1 string, part2 string, time1 time.Duration, time2 time.Duration) {
	PrintHeader(day)
	PrintAnswer(part1, part2, time1, time2)
}
