package day23

import (
	"math"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// DIRECITONS
// ========================
var Direction = []Point{
	{0, -1},  //N
	{0, 1},   //S
	{-1, 0},  //W
	{1, 0},   //E
	{1, -1},  //NE
	{-1, -1}, //NW
	{1, 1},   //SE
	{-1, 1},  //SW

}

const (
	N = iota
	S
	W
	E
	NE
	NW
	SE
	SW
)

// ========================
// MOVE MAP
// ========================
type MoveMap map[Point][]*Point

func NewMoveMap() MoveMap {
	return make(map[Point][]*Point)
}
func (m *MoveMap) Update(key Point, point *Point) {
	if _, found := (*m)[key]; found {
		existing := (*m)[key]
		existing = append(existing, point)
		(*m)[key] = existing
	} else {
		(*m)[key] = []*Point{point}
	}
}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}
func (e *Point) Update(point Point) {
	(*e).x = point.x
	(*e).y = point.y
}
func (e *Point) WillMove(evles *Evles) bool {
	for _, d := range Direction {
		if evles.Has(Point{e.x + d.x, e.y + d.y}) {
			return true
		}
	}

	return false
}
func (e *Point) GetAdjacent(index int) Point {
	d := Direction[index]
	return Point{e.x + d.x, e.y + d.y}
}
func (e *Point) ProposeMove(elves *Evles, moves *MoveMap, round int) {
	for i := 0; i < 4; i++ {
		dir := (i + round) % 4

		var check []int
		switch dir {
		case 0: // N
			check = []int{N, NE, NW}
		case 1: // S
			check = []int{S, SE, SW}
		case 2: // W
			check = []int{W, SW, NW}
		case 3: // E
			check = []int{E, NE, SE}
		}

		if elves.CheckPoints(*e, check) {
			// fmt.Println(dir)
			moves.Update(e.GetAdjacent(dir), e)
			return
		}
	}
}

// ========================
// ELVES
// ========================
type Evles struct {
	list     []*Point
	location map[Point]struct{}
}

func (e *Evles) Add(x, y int) {
	point := &Point{x, y}
	e.list = append(e.list, point)
	e.location[*point] = struct{}{}
}
func (e *Evles) Has(point Point) bool {
	_, found := e.location[point]
	return found
}
func (e *Evles) UpdateLocations() {
	clear(e.location)
	for _, elf := range (*e).list {
		e.location[*elf] = struct{}{}
	}
}
func (e *Evles) CheckPoints(origin Point, directions []int) bool {
	for _, dir := range directions {
		d := Direction[dir]
		if e.Has(Point{origin.x + d.x, origin.y + d.y}) {
			return false
		}
	}
	return true
}
func (e *Evles) Print() int {
	// Find bounds
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	for elf := range e.location {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}

	// Add padding if you like
	// padding := 0
	// minX -= padding
	// minY -= padding
	// maxX += padding
	// maxY += padding

	// // Print grid row by row
	// for y := minY; y <= maxY; y++ {
	// 	for x := minX; x <= maxX; x++ {
	// 		if e.Has(Point{x, y}) {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	rows := maxY - minY + 1
	cols := maxX - minX + 1

	return rows*cols - len(e.list)
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Evles {
	data := tools.ReadFile(file)
	elves := Evles{
		list:     []*Point{},
		location: make(map[Point]struct{}),
	}

	for y, row := range strings.Split(data, "\n") {
		for x, col := range row {
			if col == '#' {
				elves.Add(x, y)
			}
		}
	}
	return elves
}
