package common

import (
	"fmt"
	"strings"
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

func PrintAnswer(part1, part2 string) {
	fmt.Printf("%s\n%s\n\n", part1, part2)
}

func PrintExpected(part int, expected any, is_test bool) {
	if !is_test {
		return
	}
	fmt.Printf("[Part %v] Expected: %v\n", part, expected)
}

func RunDay(day int, part1 string, part2 string) {
	PrintHeader(day)
	PrintAnswer(part1, part2)
}
