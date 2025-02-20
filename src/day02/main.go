package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	common "github.com/benjuh/aoc2023/common"
)

const IS_TEST = false
const R = 0
const G = 1
const B = 2

//go:embed data/data.txt
var input string

type Game struct {
	id          int
	game_states [][3]int
}

var games []*Game

func parse(data string) []string {
	lines := common.GetLines(data)

	games = make([]*Game, len(lines))

	for i, line := range lines {
		var game Game
		var states [][3]int
		index := strings.Index(line, ":")
		game.id, _ = strconv.Atoi(line[5:index])
		games[i] = &game

		reduced := line[index+1:]
		reduced = strings.Replace(reduced, " ", "", -1)
		for _, s := range strings.Split(reduced, ";") {
			values := strings.Split(s, ",")
			curr_state := [3]int{0, 0, 0}
			for _, v := range values {
				is_red := strings.Index(v, "red")
				is_blue := strings.Index(v, "blue")
				is_green := strings.Index(v, "green")
				if is_red != -1 {
					red_val, _ := strconv.Atoi(v[0:is_red])
					curr_state[R] = red_val
				} else if is_green != -1 {
					green_val, _ := strconv.Atoi(v[0:is_green])
					curr_state[G] = green_val
				} else if is_blue != -1 {
					blue_val, _ := strconv.Atoi(v[0:is_blue])
					curr_state[B] = blue_val
				} else {
					panic("invalid state")
				}
			}
			states = append(states, curr_state)
		}
		game.game_states = states

	}

	return lines
}

func count_possible_games(red_limit, green_limit, blue_limit int) int {
	count := 0
	for _, g := range games {
		is_game_possible := true
		for _, s := range g.game_states {
			if s[R] > red_limit || s[G] > green_limit || s[B] > blue_limit {
				is_game_possible = false
				break
			}
		}
		if is_game_possible {
			count += g.id
		}
	}
	return count
}

func get_mininmum(states [][3]int) [3]int {
	max_r := states[0][R]
	max_g := states[0][G]
	max_b := states[0][B]
	for _, s := range states {
		if s[R] > max_r {
			max_r = s[R]
		}
		if s[G] > max_g {
			max_g = s[G]
		}
		if s[B] > max_b {
			max_b = s[B]
		}
	}

	return [3]int{max_r, max_g, max_b}
}

func get_power(state [3]int) int {
	return state[R] * state[G] * state[B]
}

func part1(data string) string {
	parse(data)
	sum := count_possible_games(12, 13, 14)
	return fmt.Sprintf("part_1=%v", sum)
}

func part2(data string) string {
	parse(data)
	powers := 0
	for _, g := range games {
		min := get_mininmum(g.game_states)
		powers += get_power(min)
	}
	return fmt.Sprintf("part_2=%v", powers)
}

func main() {
	part_1 := part1(input)
	part_2 := part2(input)

	common.RunDay(2, part_1, part_2)
}
