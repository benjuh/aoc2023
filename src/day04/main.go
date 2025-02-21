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

type Card struct {
	winning_numbers map[int]bool
	played_numbers  []int
}

var cards []Card

func parse(data string) []string {
	lines := common.GetLines(data)
	for _, line := range lines {
		colon_index := strings.Index(line, ":")
		usable := line[colon_index+1:]
		usable = strings.TrimSpace(usable)

		left_and_right := strings.Split(usable, "|")
		left := strings.TrimSpace(left_and_right[0])
		right := strings.TrimSpace(left_and_right[1])
		var card Card
		card.winning_numbers = make(map[int]bool)
		card.played_numbers = make([]int, 0)

		left_nums := strings.Split(left, " ")
		right_nums := strings.Split(right, " ")
		for _, num := range left_nums {
			val, _ := strconv.Atoi(string(num))
			if val == 0 {
				continue
			}
			card.winning_numbers[val] = true
		}

		for _, num := range right_nums {
			val, _ := strconv.Atoi(string(num))
			if val == 0 {
				continue
			}
			card.played_numbers = append(card.played_numbers, val)
		}
		cards = append(cards, card)
	}

	return lines
}

func part1(data string) string {
	_ = parse(data)

	sum := 0
	for _, scratchcard := range cards {
		local_sum := 0
		for _, num := range scratchcard.played_numbers {
			if scratchcard.winning_numbers[num] {
				if local_sum == 0 {
					local_sum = 1
				} else {
					local_sum *= 2
				}
			}
		}
		sum += local_sum

	}

	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	countCard := map[int]int{}

	for i, scratchcard := range cards {
		count := 0
		for _, num := range scratchcard.played_numbers {
			if scratchcard.winning_numbers[num] {
				count += 1
			}
		}
		countCard[i] += 1

		for j := i + 1; j < i+count+1; j++ {
			countCard[j] += countCard[i]
		}

	}

	total_won := 0
	for _, v := range countCard {
		total_won += v
	}

	return fmt.Sprintf("part_2=%v", total_won)
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

	common.RunDay(4, part_1, part_2, time1, time2)
}
