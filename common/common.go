package common

import (
	"fmt"
	"strings"
	"time"
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
	fmt.Printf("\n[ Day %v ]\n", day)
}

func PrintAnswer(part1, part2 string, time1 time.Duration, time2 time.Duration) {
	green := "\033[32m"
	orange := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	var color1 string
	var color2 string
	if time1.Milliseconds() < 100 {
		color1 = green
	} else if time1.Milliseconds() < 500 {
		color1 = orange
	} else {
		color1 = red
	}

	if time2.Milliseconds() < 100 {
		color2 = green
	} else if time2.Milliseconds() < 500 {
		color2 = orange
	} else {
		color2 = red
	}

	sep1 := " ~ "
	sep2 := " ~ "

	len_1 := len(time1.String())
	len_2 := len(time2.String())

	if len_1 > len_2 {
		dist := len_1 - len_2
		sep2 = strings.Repeat(" ", dist) + sep2
	} else {
		dist := len_2 - len_1
		sep1 = strings.Repeat(" ", dist) + sep1
	}

	fmt.Printf("%s%s%s%s%s\n", color1, time1.String(), reset, sep1, part1)
	fmt.Printf("%s%s%s%s%s\n", color2, time2.String(), reset, sep2, part2)
}

func PrintExpected(part int, expected any, is_test bool) {
	if !is_test {
		return
	}
	fmt.Printf("[Part %v] Expected: %v\n", part, expected)
}

func RunDay(day int, part1 string, part2 string, time1 time.Duration, time2 time.Duration) {
	PrintHeader(day)
	PrintAnswer(part1, part2, time1, time2)
}
