package day22

<<<<<<< HEAD
func Solve(file string) int {
	grid, moves := ParseInput(file)
	state := NewState(grid.start)

	for i, move := range moves {
		state.UpdateState(move, grid, i)
=======
func Solve(file string, part int) int {
	grid, moves, state := ParseInput(file, part)

	for _, move := range moves {
		state.UpdateState(move, grid)
>>>>>>> day-22-debug
	}

	return state.GetPassword()

}
