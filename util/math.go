package util

import (
	"fmt"

	"github.com/benjuh/aoc2023/common"
)

var print = fmt.Printf

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

func Sum(arr []int) int {
	var sum int
	for _, i := range arr {
		sum += i
	}
	return sum
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func Order(x, x2 int) (int, int) {
	if x <= x2 {
		return x, x2
	}
	return x2, x

}

func ShoelaceFormula(vertices []common.Point) int {
	var a int
	for i := range len(vertices) {
		j := (i + 1) % len(vertices)
		a += vertices[i].Y * vertices[j].X
		a -= vertices[i].X * vertices[j].Y
	}
	return Abs(a) / 2
}

func IsPointInPolygon(point common.Point, polygon []common.Point) bool {
	x := point.X
	y := point.Y
	n := len(polygon)
	inside := false

	p1x := polygon[0].X
	p1y := polygon[0].Y
	for i := range n + 1 {
		p2x := polygon[i%n].X
		p2y := polygon[i%n].Y
		if y > min(p1y, p2y) {
			if y <= max(p1y, p2y) {
				if x <= max(p1x, p2x) {
					var xinters int
					if p1y != p2y {
						xinters = (y-p1y)*(p2x-p1x)/(p2y-p1y) + p1x
					}
					if p1x == p2x || x <= xinters {
						inside = !inside
					}
				}
			}
		}
		p1x, p1y = p2x, p2y
	}
	return inside
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
