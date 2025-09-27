package day24

import (
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

func Solve(file string, part int) int {
	grid := parseInput(file)
	time, rounds := 0, 1

	if part == 2 {
		rounds = 3
	}

	for round := 0; round < rounds; round++ {
		if round%2 != 0 {
			grid.ReverseDirection()
			time = traverseGrid(grid, time)
			grid.ReverseDirection()
		} else {
			time = traverseGrid(grid, time)
		}
	}

	return time
}

func parseInput(file string) Grid {
	data := tools.ReadFile(file)
	rows := strings.Split(data, "\n")

	grid := Grid{
		schedule: make(Schedule),
		height:   len(rows) - 2,
		width:    len(rows[0]) - 2,
		entry:    Point{strings.IndexByte(rows[0], '.'), 0},
		exit:     Point{strings.IndexByte(rows[len(rows)-1], '.'), len(rows) - 1},
	}

	grid.lcm = tools.LCM(grid.width, grid.height)
	grid.schedule[grid.entry] = ds.NewSet[int]()
	grid.schedule[grid.exit] = ds.NewSet[int]()

	for y := 1; y < grid.height+1; y++ {
		for x := 1; x < grid.width+1; x++ {
			coord := Point{x, y}
			if _, found := grid.schedule[coord]; !found {
				grid.schedule[coord] = ds.NewSet[int]()
			}
			grid.MapStormPath(coord, rows[y][x])
		}
	}
	return grid
}

func traverseGrid(grid Grid, startTime int) int {
	queue := NewQueue(State{grid.entry, startTime})
	seen := make(map[State]bool)

	for len(queue) > 0 {
		current := queue.Pop()

		if current.AtEnd(grid.exit) {
			return current.time
		}

		for _, d := range DIRECTIONS {
			nextPos := Point{current.pos.x + d.x, current.pos.y + d.y}
			nextTime := current.time + 1
			nextState := State{nextPos, nextTime % grid.lcm}

			if grid.IsValid(nextPos, nextTime) && !seen[nextState] {
				seen[nextState] = true
				queue.Push(State{nextPos, nextTime})
			}

		}
	}

	panic("Path not found!")
}
