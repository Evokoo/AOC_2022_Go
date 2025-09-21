package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Move struct {
	direction string
	steps     int
}

func Solve(file string, part int) int {
	data := tools.ReadFile(file)
	path := make(map[Point]struct{})

	var rope Rope
	if part == 1 {
		rope = NewRope(2, Point{0, 0})
	} else {
		rope = NewRope(10, Point{0, 0})
	}
	for line := range strings.SplitSeq(data, "\n") {
		move := parseMove(line)
		rope.Update(move, path)
	}
	return len(path)
}

func parseMove(line string) Move {
	parts := strings.Split(line, " ")
	if n, err := strconv.Atoi(parts[1]); err == nil {
		return Move{direction: parts[0], steps: n}
	}
	panic(fmt.Sprintf("%s is not a valid number", parts[1]))
}
