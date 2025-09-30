package day18

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// CONSTANTS
// ========================
var Directions = []Point3D{
	{1, 0, 0},  //Positive X
	{-1, 0, 0}, //Negative X
	{0, 1, 0},  //Positive Y
	{0, -1, 0}, //Negative Y
	{0, 0, 1},  //Positive Z
	{0, 0, -1}, //Negative Z
}

var DimensionOffset = 6

// ========================
// POINT3D
// ========================
type Point3D struct{ x, y, z int }

func (p Point3D) SurfaceArea(droplet Droplet) int {
	area := 6
	for _, point := range p.GetAdjacent() {
		if droplet.Has(point) {
			area--
		}
	}
	return area
}

func (p Point3D) GetAdjacent() []Point3D {
	var adjacent []Point3D
	for _, dir := range Directions {
		adjacent = append(adjacent, Point3D{p.x + dir.x, p.y + dir.y, p.z + dir.z})
	}
	return adjacent
}

// ========================
// DROPLET
// ========================
type Droplet struct {
	blocks map[Point3D]struct{}
	xRange Range
	yRange Range
	zRange Range
}

func NewDroplet(file string) Droplet {
	droplet := Droplet{
		blocks: make(map[Point3D]struct{}),
		xRange: NewRange(),
		yRange: NewRange(),
		zRange: NewRange(),
	}

	data := tools.ReadFile(file)
	for line := range strings.SplitSeq(data, "\n") {
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		droplet.UpdateDimensions(x, y, z)
		droplet.Add(Point3D{x, y, z})
	}
	return droplet
}

func (d Droplet) Has(point Point3D) bool {
	_, found := d.blocks[point]
	return found
}

func (d Droplet) Add(point Point3D) {
	d.blocks[point] = struct{}{}
}

func (d *Droplet) UpdateDimensions(x, y, z int) {
	if x < d.xRange.min {
		d.xRange.min = x - DimensionOffset
	}
	if x > d.xRange.max {
		d.xRange.max = x + DimensionOffset
	}
	if y < d.yRange.min {
		d.yRange.min = y - DimensionOffset
	}
	if y > d.yRange.max {
		d.yRange.max = y + DimensionOffset
	}
	if z < d.zRange.min {
		d.zRange.min = z - DimensionOffset
	}
	if z > d.zRange.max {
		d.zRange.max = z + DimensionOffset
	}
}

func (d Droplet) TotalSurfaceArea() int {
	area := 0
	for point := range d.blocks {
		area += point.SurfaceArea(d)
	}
	return area
}

func (d Droplet) ExternalSurfaceArea() (area int) {
	start := Point3D{d.xRange.min, d.yRange.min, d.zRange.min}
	queue := NewQueue(start)
	seen := NewSet(start, d)

	for !queue.IsEmpty() {
		current := queue.Pop()
		for _, point := range current.GetAdjacent() {
			if point.x < d.xRange.min || point.x > d.xRange.max ||
				point.y < d.yRange.min || point.y > d.yRange.max ||
				point.z < d.zRange.min || point.z > d.zRange.max {
				continue
			}
			if d.Has(point) {
				area++
				continue
			}
			if !seen.Has(point) {
				seen.Add(point)
				queue.Push(point)
			}
		}
	}
	return area
}

// ========================
// RANGE
// ========================
type Range struct{ min, max int }

func NewRange() Range {
	return Range{min: math.MaxInt, max: math.MinInt}
}

// ========================
// SET
// ========================
type Set map[Point3D]struct{}

func NewSet(initalValue Point3D, d Droplet) Set {
	cap := (d.xRange.max - d.xRange.min + 1) * (d.yRange.max - d.yRange.min + 1) * (d.zRange.max - d.zRange.min + 1)
	set := make(Set, cap)
	set.Add(initalValue)
	return set
}

func (s Set) Has(point Point3D) bool {
	_, found := s[point]
	return found
}

func (s *Set) Add(point Point3D) {
	(*s)[point] = struct{}{}
}

// ========================
// QUEUE
// ========================
type Queue[T any] []T

func NewQueue[T any](item T) Queue[T] {
	return Queue[T]{item}
}
func (q *Queue[T]) Pop() T {
	removed := (*q)[0]
	(*q) = (*q)[1:]
	return removed
}
func (q *Queue[T]) Push(c T) {
	*q = append(*q, c)
}
func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}
