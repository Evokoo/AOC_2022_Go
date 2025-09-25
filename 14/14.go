package day14

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

type Point struct{ x, y int }
type Sand = Point
type Cave = ds.Set[Point]

func Solve(file string, part int) int {
	cave, depth := parseInput(file)

	for i := 0; ; i++ {
		grain := newParticle()
		if !cave.Has(grain) && grain.fall(cave, depth, part) {
			cave.Add(grain)
		} else {
			return i
		}
	}
}

func parseInput(file string) (Cave, int) {
	data := tools.ReadFile(file)
	cave := ds.NewSet[Point]()
	maxY := 0

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

				if current.y > maxY {
					maxY = current.y
				}
			}
			current = target
		}
	}
	return Cave(cave), maxY
}
func coordToPoint(s string) Point {
	digits := tools.QuickMatch(s, `\d+`)
	x, _ := strconv.Atoi(digits[0])
	y, _ := strconv.Atoi(digits[1])
	return Point{x, y}
}

func newParticle() Sand {
	return Sand{500, 0}
}
func (s *Sand) fall(cave Cave, maxDepth int, part int) bool {
	for {
		if part == 2 && s.y+1 == (maxDepth+2) {
			return true
		} else {
			if s.y > maxDepth {
				return false
			}
		}
		down := Sand{s.x, s.y + 1}
		left := Sand{s.x - 1, s.y + 1}
		right := Sand{s.x + 1, s.y + 1}

		switch {
		case !cave.Has(down):
			s.y++
		case !cave.Has(left):
			s.x--
			s.y++
		case !cave.Has(right):
			s.x++
			s.y++
		default:
			return true
		}
	}
}
