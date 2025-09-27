package day24

import (
	"fmt"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

func Solve(file string) int {
	parseInput(file)

	return 0
}

type Point struct{ x, y int }
type GridMap = map[Point]ds.Set[int]

func parseInput(file string) {
	data := tools.ReadFile(file)
	rows := strings.Split(data, "\n")
	height, width := len(rows)-2, len(rows[0])-2

	gridMap := make(GridMap)

	for y := 1; y < height+1; y++ {
		for x := 1; x < width+1; x++ {
			coord := Point{x, y}

			if _, found := gridMap[coord]; !found {
				gridMap[coord] = ds.NewSet[int]()
			}

			switch rows[y][x] {
			case '>':
				traceStormPath(Point{x, y}, Point{1, 0}, width, &gridMap, width, height)
			case '<':
				traceStormPath(Point{x, y}, Point{-1, 0}, width, &gridMap, width, height)
			case 'v':
				traceStormPath(Point{x, y}, Point{0, 1}, height, &gridMap, width, height)
			case '^':
				traceStormPath(Point{x, y}, Point{0, -1}, height, &gridMap, width, height)
			}
		}
	}

	fmt.Println(gridMap)
}

func traceStormPath(storm, direction Point, steps int, gridMap *GridMap, width, height int) {
	for i := 0; i < steps; i++ {
		if _, found := (*gridMap)[storm]; !found {
			(*gridMap)[storm] = ds.NewSet[int]()
		}

		(*gridMap)[storm].Add(i)

		// move
		storm.x += direction.x
		storm.y += direction.y

		// wrap horizontally
		if storm.x < 1 {
			storm.x = width
		} else if storm.x > width {
			storm.x = 1
		}

		// wrap vertically
		if storm.y < 1 {
			storm.y = height
		} else if storm.y > height {
			storm.y = 1
		}
	}
}
