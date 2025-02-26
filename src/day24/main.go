package main

import (
	_ "embed"
	"fmt"
	"strings"

	"time"

	common "github.com/benjuh/aoc2023/common"
	"github.com/benjuh/aoc2023/util"
	"gonum.org/v1/gonum/mat"
)

var print = fmt.Printf

//go:embed data/data.txt
var input string

type DataInfo struct {
	px    []float64
	py    []float64
	pz    []float64
	vx    []float64
	vy    []float64
	vz    []float64
	lines int
}

func NewDataInfo() DataInfo {
	return DataInfo{
		px: make([]float64, 0),
		py: make([]float64, 0),
		pz: make([]float64, 0),
		vx: make([]float64, 0),
		vy: make([]float64, 0),
		vz: make([]float64, 0),
	}
}

var info DataInfo

func parse(data string) []HailStone {
	lines := common.GetLines(data)
	info = NewDataInfo()
	hail := []HailStone{}
	for _, line := range lines {
		line = strings.Replace(line, " @ ", ", ", 1)
		sections := strings.Split(line, ", ")
		x := util.StringToInt(strings.Trim(sections[0], " "))
		y := util.StringToInt(strings.Trim(sections[1], " "))
		z := util.StringToInt(strings.Trim(sections[2], " "))
		velx := util.StringToInt(strings.Trim(sections[3], " "))
		vely := util.StringToInt(strings.Trim(sections[4], " "))
		velz := util.StringToInt(strings.Trim(sections[5], " "))

		info.px = append(info.px, float64(x))
		info.py = append(info.py, float64(y))
		info.pz = append(info.pz, float64(z))
		info.vx = append(info.vx, float64(velx))
		info.vy = append(info.vy, float64(vely))
		info.vz = append(info.vz, float64(velz))

		ice := HailStone{
			Pos: P{float64(x), float64(y), float64(z)},
			Vel: P{float64(velx), float64(vely), float64(velz)},
		}
		hail = append(hail, ice)

	}
	info.lines = len(lines)

	return hail
}

func part1(data string) string {
	_ = parse(data)

	r1 := float64(200000000000000)
	r2 := float64(400000000000000)
	count := 0

	for i := range info.lines {
		a1 := info.vy[i] / info.vx[i]
		b1 := info.py[i] - a1*info.px[i]

		for j := i + 1; j < info.lines; j++ {
			a2 := info.vy[j] / info.vx[j]
			b2 := info.py[j] - a2*info.px[j]

			if a1 == a2 {
				continue
			}

			x := (b2 - b1) / (a1 - a2)
			y := a1*x + b1
			t1 := (x - info.px[i]) / info.vx[i]
			t2 := (x - info.px[j]) / info.vx[j]

			if t1 > 0 && t2 > 0 && x >= r1 && x <= r2 && y >= r1 && y <= r2 {
				count++
			}
		}
	}
	return fmt.Sprintf("part_1=%v", count)
}

type P struct{ x, y, z float64 }
type HailStone struct{ Pos, Vel P }

func part2(data string) string {
	hail := parse(data)
	A00 := diff(crossMatrix(hail[0].Vel), crossMatrix(hail[1].Vel))
	A03 := diff(crossMatrix(hail[1].Pos), crossMatrix(hail[0].Pos))

	// hail[0] with hail[2]
	// (p0 - p[2]) x (v0 - v[2]) == 0
	A30 := diff(crossMatrix(hail[0].Vel), crossMatrix(hail[2].Vel))
	A33 := diff(crossMatrix(hail[2].Pos), crossMatrix(hail[0].Pos))

	A := mat.NewDense(6, 6, []float64{
		A00[0], A00[1], A00[2], A03[0], A03[1], A03[2],
		A00[3], A00[4], A00[5], A03[3], A03[4], A03[5],
		A00[6], A00[7], A00[8], A03[6], A03[7], A03[8],
		A30[0], A30[1], A30[2], A33[0], A33[1], A33[2],
		A30[3], A30[4], A30[5], A33[3], A33[4], A33[5],
		A30[6], A30[7], A30[8], A33[6], A33[7], A33[8],
	})

	b0 := diff(hail[1].Pos.cross(hail[1].Vel).toF(), hail[0].Pos.cross(hail[0].Vel).toF())
	b3 := diff(hail[2].Pos.cross(hail[2].Vel).toF(), hail[0].Pos.cross(hail[0].Vel).toF())

	b := mat.NewVecDense(6, []float64{b0[0], b0[1], b0[2], b3[0], b3[1], b3[2]})

	var x mat.VecDense
	_ = x.SolveVec(A, b)

	rock := HailStone{
		Pos: P{x.At(0, 0), x.At(1, 0), x.At(2, 0)},
		Vel: P{x.At(3, 0), x.At(4, 0), x.At(5, 0)},
	}

	res := int64(rock.Pos.x) + int64(rock.Pos.y) + int64(rock.Pos.z)

	return fmt.Sprintf("part_1=%v", res)
}

func crossMatrix(p P) []float64 {
	return []float64{
		0, -p.z, p.y,
		p.z, 0, -p.x,
		-p.y, p.x, 0,
	}
}

func diff(a, b []float64) []float64 {
	res := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] - b[i]
	}
	return res
}

// https://wikimedia.org/api/rest_v1/media/math/render/svg/3242bd71d63c393d02302c7dbe517cd0ec352d31
// https://en.wikipedia.org/wiki/Cross_product#Coordinate_notation
func (p P) cross(p2 P) P {
	return P{
		p.y*p2.z - p.z*p2.y,
		p.z*p2.x - p.x*p2.z,
		p.x*p2.y - p.y*p2.x,
	}
}

func (p P) toF() []float64 {
	return []float64{p.x, p.y, p.z}
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

	common.RunDay(24, part_1, part_2, time1, time2)
}
