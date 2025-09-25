package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

func Solve(file string) int {
	cave := parseInput(file)

	fmt.Println(cave)
	fmt.Println(len(cave))
	return 0
}

type Point struct{ x, y int }
type Cave ds.Set[Point]

func parseInput(file string) Cave {
	data := tools.ReadFile(file)
	cave := ds.NewSet[Point]()

	for line := range strings.SplitSeq(data, "\n") {
		coords := strings.Split(line, " -> ")
		current := coordToPoint(coords[0])
		cave.Add(current)

		for _, coord := range coords[1:] {
			target := coordToPoint(coord)

			for current != target {
				switch {
				case current.x < target.x:
					current.x++
				case current.x > target.x:
					current.x--
				case current.y < target.y:
					current.y++
				case current.y > target.y:
					current.y--
				}
				cave.Add(current)
			}
			current = target
		}
	}
	return Cave(cave)
}
func coordToPoint(s string) Point {
	digits := tools.QuickMatch(s, `\d+`)
	x, _ := strconv.Atoi(digits[0])
	y, _ := strconv.Atoi(digits[1])
	return Point{x, y}
}
