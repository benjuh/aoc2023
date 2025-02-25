package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type Pulse bool

func (p Pulse) String() string {
	if p {
		return "high"
	}
	return "low"
}

const (
	LowPulse  Pulse = false
	HighPulse Pulse = true
)

type Message struct {
	pulse    Pulse
	src, dst string
}

var Button = Message{LowPulse, "button", "broadcaster"}

type Network map[string]Module

type NetworkModule struct {
	name         string
	destinations []string
}

type Module interface {
	Process(Message) []Message
}

type FlipFlop struct {
	NetworkModule
	status bool
}

func (f *FlipFlop) Process(msg Message) []Message {
	if msg.pulse == LowPulse {
		f.status = !f.status

		var queue []Message
		for _, dest := range f.destinations {
			queue = append(queue, Message{Pulse(f.status), f.name, dest})

		}
		return queue
	}
	return nil
}

type Conjunction struct {
	NetworkModule
	memory map[string]Pulse
}

func (c *Conjunction) Process(msg Message) []Message {
	c.memory[msg.src] = msg.pulse

	allHigh := true
	for _, v := range c.memory {
		if !v {
			allHigh = false
		}
	}

	var q []Message
	for _, d := range c.destinations {
		if allHigh {
			q = append(q, Message{LowPulse, c.name, d})
		} else {
			q = append(q, Message{HighPulse, c.name, d})
		}
	}
	return q
}

type Broadcast struct {
	NetworkModule
}

func (b *Broadcast) Process(msg Message) []Message {
	var q []Message
	for _, d := range b.destinations {
		q = append(q, Message{msg.pulse, b.name, d})
	}
	return q
}

func parse(data string) Network {
	lines := common.GetLines(data)
	nw := Network{}
	inputs := map[string][]string{}
	conjunctions := []string{}

	for _, line := range lines {
		sections := strings.Split(line, " -> ")
		name := ""
		dest := strings.Split(sections[1], ", ")

		var mod Module

		switch sections[0][0] {
		case '%':
			name = sections[0][1:]
			mod = &FlipFlop{
				NetworkModule: NetworkModule{
					name:         name,
					destinations: dest,
				},
				status: false,
			}
		case '&':
			name = sections[0][1:]
			mod = &Conjunction{
				NetworkModule: NetworkModule{
					name:         name,
					destinations: dest,
				},
				memory: make(map[string]Pulse),
			}
			conjunctions = append(conjunctions, name)
		case 'b':
			name = "broadcaster"
			mod = &Broadcast{
				NetworkModule: NetworkModule{
					name:         name,
					destinations: dest,
				},
			}
		}

		for _, d := range dest {
			inputs[d] = append(inputs[d], name)
		}

		nw[name] = mod
	}

	for _, c := range conjunctions {
		for _, in := range inputs[c] {
			con := nw[c].(*Conjunction)
			con.memory[in] = false
		}
	}

	return nw
}

func part1(data string) string {
	nw := parse(data)

	var queue []Message
	low, high := 0, 0

	count := func(m Message) {
		if m.pulse == LowPulse {
			low++
		} else {
			high++
		}
	}
	button_presses := 1000
	for _ = range button_presses {
		queue = append(queue, Button)
		var msg Message
		for len(queue) > 0 {
			msg = queue[0]
			queue = queue[1:]
			count(msg)

			if nw[msg.dst] != nil {
				messages := slices.Clone(nw[msg.dst].Process(msg))
				queue = append(queue, messages...)
			}
		}
	}
	return fmt.Sprintf("part_1=%v", low*high)
}

func part2(data string) string {
	nw := parse(data)

	firstOccurrence := map[string]int{
		"zc": 0,
		"mk": 0,
		"fp": 0,
		"xt": 0,
	}

	var res int

	count := len(firstOccurrence)

	var queue []Message
	for i := 1; ; i++ {
		queue = append(queue, Button)
		var msg Message
		for len(queue) > 0 {
			msg, queue = queue[0], queue[1:]

			if msg.dst == "rx" && msg.pulse == LowPulse {
				return fmt.Sprintf("part_2=%v", res)
			}

			first, exists := firstOccurrence[msg.src]

			if exists && msg.pulse == HighPulse && first == 0 {
				firstOccurrence[msg.src] = i
				count--
			}

			if count == 0 {
				var values []int
				for _, v := range firstOccurrence {
					values = append(values, v)
				}
				res = lcmm(values)
				return fmt.Sprintf("part_2=%v", res)
			}

			if nw[msg.dst] != nil {
				messages := slices.Clone(nw[msg.dst].Process(msg))
				queue = append(queue, messages...)
			}
		}
	}

}

func lcmm(values []int) int {
	lcm := func(a, b int) int { return a * b / util.Gcd(a, b) }

	result := 1
	for _, v := range values {
		result = lcm(result, v)
	}
	return result
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

	common.RunDay(20, part_1, part_2, time1, time2)
}
