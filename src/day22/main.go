package main

import (
	_ "embed"
	"fmt"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type Point struct {
	x, y, z int
}

type Brick struct {
	id         int
	start, end Point
}

type Space struct {
	Map    map[Point]int
	Bricks map[int]*Brick
}

func NewSpace() *Space {
	return &Space{
		Map:    make(map[Point]int),
		Bricks: make(map[int]*Brick),
	}
}

func (s *Space) Add(b *Brick) {
	s.Bricks[b.id] = b
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			for z := b.start.z; z <= b.end.z; z++ {
				s.Map[Point{x, y, z}] = b.id
			}
		}
	}
}

var s *Space

func parse(data string) []string {
	s = NewSpace()
	lines := common.GetLines(data)

	for i, line := range lines {
		sections := strings.Split(line, "~")
		var start, end Point
		brick_1_sections := strings.Split(sections[0], ",")
		brick_2_sections := strings.Split(sections[1], ",")

		start.x = util.StringToInt(brick_1_sections[0])
		start.y = util.StringToInt(brick_1_sections[1])
		start.z = util.StringToInt(brick_1_sections[2])
		end.x = util.StringToInt(brick_2_sections[0])
		end.y = util.StringToInt(brick_2_sections[1])
		end.z = util.StringToInt(brick_2_sections[2])

		s.Add(&Brick{id: i, start: start, end: end})
	}

	return lines
}

func (s *Space) Fall() int {
	fallen := map[int]struct{}{}
	for changed := true; changed; {
		changed = false
		for _, b := range s.Bricks {
			for s.EmptyBelow(b) {
				s.Move(b, Point{0, 0, -1})
				changed = true
				fallen[b.id] = struct{}{}
			}
		}
	}
	return len(fallen)
}

func (s *Space) EmptyBelow(b *Brick) bool {
	if b.start.z == 1 {
		return false
	}
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			if _, full := s.Map[Point{x, y, b.start.z - 1}]; full {
				return false
			}
		}
	}
	return true
}

func (s *Space) Move(b *Brick, p Point) {
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			for z := b.start.z; z <= b.end.z; z++ {
				delete(s.Map, Point{x, y, z})
			}
		}
	}
	for x := b.start.x + p.x; x <= b.end.x+p.x; x++ {
		for y := b.start.y + p.y; y <= b.end.y+p.y; y++ {
			for z := b.start.z + p.z; z <= b.end.z+p.z; z++ {
				s.Map[Point{x, y, z}] = b.id
			}
		}
	}

	s.Bricks[b.id].start.x += p.x
	s.Bricks[b.id].start.y += p.y
	s.Bricks[b.id].start.z += p.z
	s.Bricks[b.id].end.x += p.x
	s.Bricks[b.id].end.y += p.y
	s.Bricks[b.id].end.z += p.z
}

func (s *Space) TopOf(b *Brick) []*Brick {
	var onTop []*Brick
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			if id, ok := s.Map[Point{x, y, b.end.z + 1}]; ok {
				onTop = append(onTop, s.Bricks[id])
			}
		}
	}
	return onTop
}

func (s *Space) SafeToDisintegrate(b *Brick) bool {
	for _, topB := range s.TopOf(b) {
		if s.DependsOn(topB, b) {
			return false
		}
	}
	return true
}

func (s *Space) DependsOn(topB, botB *Brick) bool {
	for x := topB.start.x; x <= topB.end.x; x++ {
		for y := topB.start.y; y <= topB.end.y; y++ {
			if id, full := s.Map[Point{x, y, topB.start.z - 1}]; full && id != botB.id {
				return false
			}
		}
	}
	return true
}

func part1(data string) string {
	_ = parse(data)
	_ = s.Fall()

	count := 0
	for _, b := range s.Bricks {
		if s.SafeToDisintegrate(b) {
			count++
		}
	}
	return fmt.Sprintf("part_1=%v", count)
}

func (s *Space) Delete(b *Brick) {
	for x := b.start.x; x <= b.end.x; x++ {
		for y := b.start.y; y <= b.end.y; y++ {
			for z := b.start.z; z <= b.end.z; z++ {
				delete(s.Map, Point{x, y, z})
			}
		}
	}
	delete(s.Bricks, b.id)
}

func (s *Space) DeepCopy() *Space {
	s2 := NewSpace()
	for k, v := range s.Bricks {
		s2.Bricks[k] = &Brick{
			id:    v.id,
			start: v.start,
			end:   v.end,
		}
	}
	for k, v := range s.Map {
		s2.Map[k] = v
	}
	return s2
}

func part2(data string) string {
	_ = parse(data)
	_ = s.Fall()

	count := 0
	for _, b := range s.Bricks {
		s2 := s.DeepCopy()
		s2.Delete(b)
		count += s2.Fall()
	}
	return fmt.Sprintf("part_2=%v", count)
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

	common.RunDay(22, part_1, part_2, time1, time2)
}
