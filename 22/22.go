package day22

import "fmt"

func Solve(file string, part int) int {
	grid, moves := ParseInput(file, part)
	state := NewState(grid.start)

	// state.pos = Point{6, 6}
	// state.bearing = 0

	for i, move := range moves {
		fmt.Println(move)
		state.UpdateState(move, grid, i)
	}

	fmt.Println(state.GetPassword())

	return -1

}
