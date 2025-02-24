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
