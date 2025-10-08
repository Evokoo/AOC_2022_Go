package day23

func Solve(file string, part int) int {
	elves := ParseInput(file)

	if part == 1 {
		for elves.round < 10 {
			elves.StartRound(elves.round)
			elves.EndRound()
		}
		return elves.CalcuateFreeSpace()
	}

	if part == 2 {
		for !elves.endSimulation {
			elves.StartRound(elves.round)
			elves.EndRound()
		}
		return elves.round
	}

	return 0
}
