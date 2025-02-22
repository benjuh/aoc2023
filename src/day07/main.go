package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

//go:embed data/data.txt
var input string

type CamelCard struct {
	hand   string
	bid    int
	score1 int
	score2 int
}

var camel_cards []*CamelCard

func parse(data string) []string {
	lines := common.GetLines(data)
	camel_cards = []*CamelCard{}
	for _, line := range lines {
		card_info := strings.Split(line, " ")
		bid, _ := strconv.Atoi(card_info[1])
		camel_cards = append(camel_cards, &CamelCard{
			hand:   card_info[0],
			bid:    bid,
			score1: hand_score(card_info[0]),
		})
	}

	return lines
}

var strength_of = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func hand_score(hand string) int {
	rank_counts := map[string]int{}

	for _, card := range hand {
		rank_counts[string(card)]++
	}

	unique_nums := len(rank_counts)
	switch unique_nums {
	case 1:
		return 6
	case 2:
		// either 4 of a kind or full house
		is_four_of_a_kind := false
		for _, v := range rank_counts {
			if v == 4 {
				is_four_of_a_kind = true
				break
			}
		}
		if is_four_of_a_kind {
			return 5
		} else {
			return 4
		}
	case 3:
		// either 3 of a kind or 2 pair
		is_three_of_a_kind := false
		for _, v := range rank_counts {
			if v == 3 {
				is_three_of_a_kind = true
				break
			}
		}
		if is_three_of_a_kind {
			return 3
		} else {
			return 2
		}
	case 4:
		return 1
	default:
		return 0
	}
}

func score_with_joker(hand string) int {
	rank_counts := map[string]int{}

	for _, card := range hand {
		rank_counts[string(card)]++
	}
	first_score := hand_score(hand)

	jokers := rank_counts["J"]
	switch jokers {
	case 1:
		switch first_score {
		case 0:
			return 1
		case 1:
			return 3
		case 2:
			return 4
		case 3:
			return 5
		case 4:
			return 5
		case 5:
			return 6
		}
	case 2:
		switch first_score {
		case 1:
			return 3
		case 2:
			return 5
		case 3:
			return 6
		case 4:
			return 6
		}
	case 3:
		switch first_score {
		case 3:
			return 5
		case 4:
			return 6
		}
	case 4:
		switch first_score {
		case 4:
			return 6
		case 5:
			return 6
		}
	}

	return first_score
}

func part1(data string) string {
	_ = parse(data)

	sort.Slice(camel_cards, func(i, j int) bool {
		if camel_cards[i].score1 > camel_cards[j].score1 {
			return true
		} else if camel_cards[i].score1 < camel_cards[j].score1 {
			return false
		} else {
			for idx, value := range camel_cards[i].hand {
				if strength_of[string(value)] < strength_of[string(camel_cards[j].hand[idx])] {
					return false
				} else if strength_of[string(value)] > strength_of[string(camel_cards[j].hand[idx])] {
					return true
				} else {
					continue
				}
			}

		}
		panic("should not reach here")
	})

	sum := 0
	for i, camel_card := range camel_cards {
		sum += camel_card.bid * (len(camel_cards) - i)
	}

	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	_ = parse(data)
	strength_of["J"] = 1
	for _, camel_card := range camel_cards {
		camel_card.score2 = score_with_joker(camel_card.hand)
	}

	sort.Slice(camel_cards, func(i, j int) bool {
		if camel_cards[i].score2 > camel_cards[j].score2 {
			return true
		} else if camel_cards[i].score2 < camel_cards[j].score2 {
			return false
		} else {
			for idx, value := range camel_cards[i].hand {
				if strength_of[string(value)] < strength_of[string(camel_cards[j].hand[idx])] {
					return false
				} else if strength_of[string(value)] > strength_of[string(camel_cards[j].hand[idx])] {
					return true
				} else {
					continue
				}
			}

		}
		panic("should not reach here")
	})

	sum := 0
	for i, camel_card := range camel_cards {
		sum += camel_card.bid * (len(camel_cards) - i)
	}
	return fmt.Sprintf("part_2=%v", sum)
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

	common.RunDay(7, part_1, part_2, time1, time2)
}
