package day12

import (
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

type Point struct{ x, y int }
type Grid struct {
	start  Point
	end    Point
	layout map[Point]int
}
type Climber struct {
	position  Point
	steps     int
	elevation int
}

type Queue []Climber

func (q *Queue) pop() Climber {
	removed := (*q)[0]
	(*q) = (*q)[1:]
	return removed
}
func (q *Queue) push(c Climber) {
	*q = append(*q, c)
}

func Solve(file string, part int) int {
	grid := parseInput(file)
	return traverseGird(grid)
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
func traverseGird(grid Grid) int {
	queue := Queue{Climber{grid.start, 0, 0}}
	seen := ds.NewSet[Point]()

	for len(queue) > 0 {
		current := queue.pop()

		if current.position == grid.end {
			return current.steps
		}

		for _, d := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			n := Point{current.position.x + d.x, current.position.y + d.y}

			if elevation, found := grid.layout[n]; found && !seen.Has(n) {
				if elevation <= current.elevation+1 {
					queue.push(Climber{n, current.steps + 1, elevation})
					seen.Add(n)
				}
			}
		}
	}

	panic("Something went wrong!")
}
