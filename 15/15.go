package day15

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Point struct{ x, y int }
type Range struct{ start, end int }
type Pair struct {
	sensor Point
	beacon Point
	power  int
}

func Solve(file string) int {
	row := 10

	if file == "input.txt" {
		row = 2000000
	}

	pairs := parseInput(file)
	positions := locateBeacons(pairs, row)
	return positions
}

func parseInput(file string) []Pair {
	data := tools.ReadFile(file)
	var output []Pair

	for line := range strings.SplitSeq(data, "\n") {
		var x1, y1, x2, y2 int

		for i, digit := range tools.QuickMatch(line, `\d+`) {
			n, _ := strconv.Atoi(digit)
			switch i {
			case 0:
				x1 = n
			case 1:
				y1 = n
			case 2:
				x2 = n
			case 3:
				y2 = n
			}
		}
		sensor := Point{x1, y1}
		beacon := Point{x2, y2}
		output = append(output, Pair{sensor, beacon, manhattanDistance(beacon, sensor)})
	}

	return output
}
func locateBeacons(pairs []Pair, row int) int {
	var ranges []Range

	for _, pair := range pairs {
		if (row <= pair.sensor.y+pair.power) && (row >= pair.sensor.y-pair.power) {
			d := abs(pair.sensor.y - row)
			w := abs(pair.power - d)
			ranges = append(ranges, Range{pair.sensor.x - w, pair.sensor.x + w})
		}
	}

	return mergeRanges(ranges)
}
func mergeRanges(ranges []Range) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}

	for _, curr := range ranges[1:] {
		last := &merged[len(merged)-1]

		if curr.start <= last.end {
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}

	total := 0

	for _, r := range merged {
		total += r.end - r.start
	}

	return total
}
func manhattanDistance(a Point, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
