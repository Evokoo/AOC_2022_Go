package day24

import (
	"fmt"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

func Solve(file string, part int) int {
	grid := parseInput(file)
	time := 0
	rounds := 1

	if part == 2 {
		rounds = 3
	}

	for round := 0; round < rounds; round++ {
		start := grid.entry
		goal := grid.exit
		if round%2 == 1 { // odd rounds: go back
			start, goal = goal, start
		}
		legTime := navigateGrid(Grid{grid.schedule, start, goal, grid.lcm}, time)
		time = legTime
	}

	// fmt.Printf("TIME for %s is %d\n", file, time)

	return time
}

type Point struct{ x, y int }
type Grid struct {
	schedule map[Point]ds.Set[int]
	entry    Point
	exit     Point
	lcm      int
}

func parseInput(file string) Grid {
	data := tools.ReadFile(file)
	rows := strings.Split(data, "\n")
	height, width := len(rows)-2, len(rows[0])-2

	grid := Grid{schedule: make(map[Point]ds.Set[int])}
	grid.lcm = tools.LCM(width, height)

	for y := 1; y < height+1; y++ {
		for x := 1; x < width+1; x++ {
			coord := Point{x, y}

			if _, found := grid.schedule[coord]; !found {
				grid.schedule[coord] = ds.NewSet[int]()
			}

			switch rows[y][x] {
			case '>':
				traceStormPath(Point{x, y}, Point{1, 0}, width, height, &grid.schedule, grid.lcm)
			case '<':
				traceStormPath(Point{x, y}, Point{-1, 0}, width, height, &grid.schedule, grid.lcm)
			case 'v':
				traceStormPath(Point{x, y}, Point{0, 1}, width, height, &grid.schedule, grid.lcm)
			case '^':
				traceStormPath(Point{x, y}, Point{0, -1}, width, height, &grid.schedule, grid.lcm)
			}
		}
	}

	grid.entry = Point{strings.IndexByte(rows[0], '.'), 0}
	grid.exit = Point{strings.IndexByte(rows[len(rows)-1], '.'), len(rows) - 1}
	grid.schedule[grid.entry] = ds.Set[int]{}
	grid.schedule[grid.exit] = ds.Set[int]{}

	return grid
}
func traceStormPath(storm, direction Point, width, height int, schedule *map[Point]ds.Set[int], lcm int) {
	pos := storm

	for t := 0; t < lcm; t++ {
		// Initialize set if missing
		if _, found := (*schedule)[pos]; !found {
			(*schedule)[pos] = ds.NewSet[int]()
		}

		// Mark storm at this position at time t
		(*schedule)[pos].Add(t)

		// Move storm
		pos.x += direction.x
		pos.y += direction.y

		// Wrap horizontally
		if pos.x < 1 {
			pos.x = width
		} else if pos.x > width {
			pos.x = 1
		}

		// Wrap vertically
		if pos.y < 1 {
			pos.y = height
		} else if pos.y > height {
			pos.y = 1
		}
	}
}

type State struct {
	position Point
	time     int
}

func navigateGrid(grid Grid, startTime int) int {
	queue := NewQueueWith(State{grid.entry, startTime})
	seen := make(map[State]bool)

	directions := []Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {0, 0}, // 4 moves + wait
	}

	for len(queue) > 0 {
		current := queue.Pop()
		// fmt.Printf("t=%d pos=(%d,%d)\n", current.time, current.position.x, current.position.y)

		// Reached exit
		if current.position == grid.exit {
			return current.time
		}

		for _, d := range directions {
			nextPos := Point{current.position.x + d.x, current.position.y + d.y}
			nextTime := current.time + 1
			nextState := State{nextPos, nextTime % grid.lcm}

			// Out of bounds: only entry/exit can be outside the main grid
			if _, ok := grid.schedule[nextPos]; !ok {
				// fmt.Println("OUT OF BOUNDS")
				continue
			}

			// Blocked by storm
			if grid.schedule[nextPos].Has(nextState.time) {
				// fmt.Println("STORM")
				continue
			}

			// Already visited this state
			if seen[nextState] {
				// fmt.Println("STATE ALREADY SEEN")
				continue
			}

			seen[nextState] = true
			queue.Push(State{nextPos, nextTime})
		}
	}

	fmt.Println("No path found")
	return -1
}
