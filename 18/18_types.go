package day18

import (
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

// ========================
// POINT3D
// ========================
type Point3D struct{ x, y, z int }

func (p Point3D) SurfaceArea(droplet Droplet) int {
	area := 6
	for _, dir := range Directions {
		n := Point3D{p.x + dir.x, p.y + dir.y, p.z + dir.z}

		if droplet.Has(n) {
			area--
		}
	}
	return area
}

// ========================
// DROPLET
// ========================
type Droplet map[Point3D]struct{}

func NewDroplet(file string) Droplet {
	droplet := make(Droplet)
	data := tools.ReadFile(file)

	for line := range strings.SplitSeq(data, "\n") {
		xyz := strings.Split(line, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		droplet.Add(Point3D{x, y, z})
	}

	return droplet
}

func (d Droplet) Has(point Point3D) bool {
	_, found := d[point]
	return found
}

func (d Droplet) Add(point Point3D) {
	d[point] = struct{}{}
}

func (d Droplet) TotalSurfaceArea() int {
	area := 0
	for point := range d {
		area += point.SurfaceArea(d)
	}
	return area
}
