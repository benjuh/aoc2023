package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
)

const (
	IS_TEST = false
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

var HashQueue []string

func parse(data string) []string {
	lines := common.GetLines(data)
	for _, line := range lines {
		strs := strings.SplitSeq(line, ",")
		for str := range strs {
			HashQueue = append(HashQueue, str)
		}
	}
	return lines
}

func HASH(s string) int {
	sum := 0
	for _, char := range s {
		sum += util.AsciiValue(string(char))
		sum = (sum * 17) % 256
	}
	return sum
}

func part1(data string) string {
	_ = parse(data)
	hashsum := 0
	for _, str := range HashQueue {
		hashsum += HASH(str)
	}
	return fmt.Sprintf("part_1=%v", hashsum)
}

type Box []Lense
type Lense struct {
	label        string
	focal_length int
}

func part2(data string) string {
	_ = parse(data)
	boxes := make([]Box, 256)
	label_to_box_index := map[string]int{}
	total_steps := len(HashQueue)
	for i := range total_steps {
		remove_operation := strings.Index(HashQueue[i], "-")
		if remove_operation != -1 {
			label := HashQueue[i][:remove_operation]
			box_index, ok := label_to_box_index[label]
			if ok {
				for i := range len(boxes[box_index]) - 1 {
					if boxes[box_index][i].label == label {
						boxes[box_index][i], boxes[box_index][i+1] = boxes[box_index][i+1], boxes[box_index][i]
					}
				}
				boxes[box_index] = boxes[box_index][:len(boxes[box_index])-1]
				delete(label_to_box_index, label)
			}
		}
		equals_operation := strings.Index(HashQueue[i], "=")
		if equals_operation != -1 {
			label := HashQueue[i][:equals_operation]
			focal_length, _ := strconv.Atoi(HashQueue[i][equals_operation+1:])
			box_index := HASH(label)
			old_box_index, ok := label_to_box_index[label]
			if ok {
				if old_box_index != box_index {
					panic("hashes should be the same...")
				}
				for i := range len(boxes[box_index]) {
					if boxes[box_index][i].label == label {
						boxes[box_index][i].focal_length = focal_length
					}
				}
			} else {
				boxes[box_index] = append(boxes[box_index], Lense{
					label:        label,
					focal_length: focal_length,
				})
				label_to_box_index[label] = box_index
			}
		}
	}
	focusing_power := 0
	for box_index, box := range boxes {
		for lense_index, lense := range box {
			focusing_power += (box_index + 1) * (lense_index + 1) * lense.focal_length
		}
	}

	return fmt.Sprintf("part_2=%v", focusing_power)
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

	common.RunDay(15, part_1, part_2, time1, time2)
}
