package day22

func Solve(file string, part int) int {
	grid, moves, state := ParseInput(file, part)

	for _, move := range moves {
		state.UpdateState(move, grid)
	}

	return state.GetPassword()

}
