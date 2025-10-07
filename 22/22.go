package day22

func Solve(file string, part int) int {
	grid, moves := ParseInput(file, part)
	state := NewState(grid.start)

	for i, move := range moves {
		state.UpdateState(move, grid, i)
	}

	return state.GetPassword()

}
