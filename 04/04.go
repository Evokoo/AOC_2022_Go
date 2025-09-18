package day04

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Range struct {
	Start int
	End   int
}

func (a Range) Contains(b Range) bool {
	return a.Start <= b.Start && a.End >= b.End
}

func (a Range) Overlaps(b Range) bool {
	return a.Start <= b.End && b.Start <= a.End
}

func Solve(file string, part int) int {
	data := strings.Split(tools.ReadFile(file), "\n")
	partOne, partTwo := checkRanges(parseInput(data))

	switch part {
	case 1:
		return partOne
	case 2:
		return partTwo
	default:
		panic("Invalid part")
	}
}

func parseInput(data []string) [][2]Range {
	ranges := make([][2]Range, len(data))

	for i, line := range data {
		var rangePair [2]Range

		for j, pair := range strings.Split(line, ",") {
			values := strings.Split(pair, "-")
			start, _ := strconv.Atoi(values[0])
			end, _ := strconv.Atoi(values[1])
			rangePair[j] = Range{Start: start, End: end}
		}
		ranges[i] = rangePair
	}

	return ranges
}
func checkRanges(rangePairs [][2]Range) (int, int) {
	contained, overlapping := 0, 0
	for _, rangePair := range rangePairs {
		a, b := rangePair[0], rangePair[1]
		if a.Contains(b) || b.Contains(a) {
			contained++
		}
		if a.Overlaps(b) {
			overlapping++
		}
	}
	return contained, overlapping
}
