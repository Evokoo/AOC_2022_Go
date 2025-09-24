package day12

import (
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

func Solve(file string, part int) int {
	grid := parseInput(file)

	var climber Climber
	if part == 1 {
		climber = Climber{grid.start, 0, 0, true}
	} else {
		climber = Climber{grid.end, 0, 25, false}
	}

	return traverseGrid(climber, grid)
}
func parseInput(file string) Grid {
	data := tools.ReadFile(file)
	rows := strings.Split(data, "\n")
	grid := Grid{layout: make(map[Point]int)}

	for y := range rows {
		for x := range rows[y] {
			cell, point := rows[y][x], Point{x, y}
			switch cell {
			case 'S':
				grid.layout[point] = 0
				grid.start = point
			case 'E':
				grid.layout[point] = 25
				grid.end = point
			default:
				grid.layout[point] = int(cell - 'a')
			}
		}
	}
	return grid
}
func traverseGrid(climber Climber, grid Grid) int {
	queue := Queue{climber}
	seen := ds.NewSet[Point]()

	for len(queue) > 0 {
		c := queue.pop()

		if (c.ascending && c.position == grid.end) || (!c.ascending && c.elevation == 0) {
			return c.steps
		}

		c.moveToAdjacent(grid, &queue, &seen)
	}

	panic("Goal not reached!")
}

// POINT
type Point struct{ x, y int }

// GRID
type Grid struct {
	start  Point
	end    Point
	layout map[Point]int
}

// CLIMBER
type Climber struct {
	position  Point
	steps     int
	elevation int
	ascending bool
}

func (c Climber) moveToAdjacent(grid Grid, queue *Queue, seen *ds.Set[Point]) {
	for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		next := Point{c.position.x + dir.x, c.position.y + dir.y}

		if elevation, found := grid.layout[next]; found && !(*seen).Has(next) {
			isValid := (c.ascending && elevation <= c.elevation+1) || (!c.ascending && elevation >= c.elevation-1)
			if isValid {
				(*queue).push(Climber{next, c.steps + 1, elevation, c.ascending})
				(*seen).Add(next)
			}
		}
	}
}

// QUEUE
type Queue []Climber

func (q *Queue) pop() Climber {
	removed := (*q)[0]
	(*q) = (*q)[1:]
	return removed
}
func (q *Queue) push(c Climber) {
	*q = append(*q, c)
}
