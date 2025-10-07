package day22

func Solve(file string) int {
	grid, moves := ParseInput(file)
	state := NewState(grid.start)

	for i, move := range moves {
		state.UpdateState(move, grid, i)
	}

	return state.GetPassword()

	return 0
}
