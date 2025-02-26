package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

// Node represents a node of a graph
type Node struct {
	label          string
	connectedNodes []*Node
}

// Edge connecting two nodes
type Edge struct {
	u *Node
	v *Node
}

func parse(data string) map[string]*Node {
	res := map[string]*Node{}
	lines := common.GetLines(data)

	for _, line := range lines {
		sections := strings.Split(line, ": ")
		label := sections[0]

		next_sections := strings.Split(sections[1], " ")
		source, ok := res[label]
		if !ok {
			source = &Node{label: label}
			res[label] = source
		}

		for _, other := range next_sections {
			target, ok := res[other]
			if !ok {
				target = &Node{label: other}
				res[other] = target
			}
			source.connectedNodes = append(source.connectedNodes, target)
			target.connectedNodes = append(target.connectedNodes, source)
		}
	}

	return res
}

func part1(data string) string {
	graph := parse(data)
	var ans string
	for {
		cut, res := minCut(graph)
		if cut == 6 {
			ans = strconv.Itoa(res)
			break
		}
	}

	return fmt.Sprintf("part_1=%v", ans)
}

func minCut(nodeMap map[string]*Node) (int, int) {
	vertices := len(nodeMap)
	sets := make([][]*Node, 0, len(nodeMap))
	var edges []Edge
	for _, u := range nodeMap {
		sets = append(sets, []*Node{u})
		for _, v := range u.connectedNodes {
			edges = append(edges, Edge{u: u, v: v})
		}
	}

	for vertices > 2 {
		i := rand.Intn(len(edges))
		set1 := find(sets, edges[i].u)
		set2 := find(sets, edges[i].v)

		if set1 != set2 {
			vertices--
			sets = union(sets, set1, set2)
		}
	}

	cut := 0
	for _, edge := range edges {
		set1 := find(sets, edge.u)
		set2 := find(sets, edge.v)
		if set1 != set2 {
			cut++
		}
	}

	return cut, len(sets[0]) * len(sets[1])
}

func find(sets [][]*Node, u *Node) int {
	for i, set := range sets {
		if slices.Contains(set, u) {
			return i
		}
	}
	panic("node not found")
}

func union(sets [][]*Node, set1 int, set2 int) [][]*Node {
	sets[set1] = append(sets[set1], sets[set2]...)
	return append(sets[:set2], sets[set2+1:]...)
}

func part2(data string) string {
	// _ = parse(data)
	return fmt.Sprintf("part_2=%v", "DONE!")
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

	common.RunDay(25, part_1, part_2, time1, time2)
}
