package common

type Point struct {
	X int
	Y int
}

func (p Point) Next(dx, dy int) Point {
	return Point{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}

type Delta struct {
	Dx int
	Dy int
}

type Directions [4]Delta

func GetDirections() Directions {
	return Directions{
		Delta{0, -1},
		Delta{1, 0},
		Delta{0, 1},
		Delta{-1, 0},
	}
}
