package day24

import "github.com/Evokoo/AOC_2022_Go/tools/ds"

// ========================
// CONSTANTS
// ========================
var DIRECTIONS = map[byte]Point{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
	'.': {0, 0},
}

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// GRID
// ========================
type Schedule = map[Point]ds.Set[int]

type Grid struct {
	schedule Schedule
	width    int
	height   int
	entry    Point
	exit     Point
	lcm      int
}

func (g *Grid) IsValid(p Point, t int) bool {
	if _, ok := g.schedule[p]; !ok {
		return false
	}
	if g.schedule[p].Has(t % g.lcm) {
		return false
	}
	return true
}
func (g *Grid) MapStormPath(start Point, key byte) {
	if key == '.' {
		return
	}

	direction := DIRECTIONS[key]
	pos := start

	for i := range g.lcm {
		if _, found := g.schedule[pos]; !found {
			g.schedule[pos] = ds.NewSet[int]()
		}

		g.schedule[pos].Add(i)

		pos.x = (pos.x-1+direction.x+g.width)%g.width + 1
		pos.y = (pos.y-1+direction.y+g.height)%g.height + 1
	}
}
func (g *Grid) ReverseDirection() {
	start := g.entry
	end := g.exit

	g.entry = end
	g.exit = start
}

// ========================
// STATE
// ========================
type State struct {
	pos  Point
	time int
}

func (s State) AtEnd(target Point) bool {
	return s.pos == target
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
