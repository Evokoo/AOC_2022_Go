package day08

import (
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Matrix [][]rune

func Solve(file string, part int) int {
	data := tools.ReadFile(file)
	matrix := parseInput(data)
	visible, viewScore := inspectMatrix(matrix)

	if part == 2 {
		return viewScore
	}

	return visible
}

func parseInput(data string) Matrix {
	var matrix Matrix
	for row := range strings.SplitSeq(data, "\n") {
		matrix = append(matrix, []rune(row))
	}
	return matrix
}

func inspectMatrix(matrix Matrix) (int, int) {
	size := len(matrix)
	visible, bestView := 0, 0

	for y := range size {
		for x := range size {
			height := matrix[y][x]

			directions := []struct{ dx, dy int }{
				{1, 0},  // right
				{-1, 0}, // left
				{0, 1},  // down
				{0, -1}, // up
			}

			score := 1
			isVisible := false

			for _, dir := range directions {
				cx, cy := x+dir.dx, y+dir.dy
				distance := 0

				for (cx >= 0 && cx < size) && (cy >= 0 && cy < size) {
					distance++

					if matrix[cy][cx] >= height {
						break
					}
					cx += dir.dx
					cy += dir.dy
				}
				score *= distance
				if cx < 0 || cx >= size || cy < 0 || cy >= size {
					isVisible = true
				}
			}

			if isVisible {
				visible++
			}
			if score > bestView {
				bestView = score
			}
		}
	}

	return visible, bestView
}
