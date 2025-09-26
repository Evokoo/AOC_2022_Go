package day15

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Point = tools.Point
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
		coords := make([]int, 4)
		for i, digit := range tools.QuickMatch(line, `-*\d+`) {
			n, _ := strconv.Atoi(digit)
			coords[i] = n
		}
		sensor := Point{X: coords[0], Y: coords[1]}
		beacon := Point{X: coords[2], Y: coords[3]}
		output = append(output, Pair{sensor, beacon, tools.ManhattanDistance(beacon, sensor)})
	}
	return output
}
func locateBeacons(pairs []Pair, row int) int {
	var ranges []Range

	for _, pair := range pairs {
		if (row <= pair.sensor.Y+pair.power) && (row >= pair.sensor.Y-pair.power) {
			d := tools.Abs(pair.sensor.Y - row)
			w := tools.Abs(pair.power - d)
			ranges = append(ranges, Range{pair.sensor.X - w, pair.sensor.X + w})
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

	occupied := 0
	for _, r := range merged {
		occupied += r.end - r.start
	}
	return occupied
}
