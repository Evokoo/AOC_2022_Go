package day23

func Solve(file string, part int) int {
	elves := ParseInput(file)

	if part == 1 {
		for round := 0; round < 10; round++ {
			moves := NewMoveMap()

			for _, elf := range elves.list {
				if elf.WillMove(&elves) {
					elf.ProposeMove(&elves, &moves, round)
				}
			}
			for point, proposed := range moves {
				if len(proposed) == 1 {
					proposed[0].Update(point)
				}
			}

			elves.UpdateLocations()
		}
		return elves.Print()
	}

	if part == 2 {
		for round := 0; ; round++ {
			moves := NewMoveMap()
			updates := 0

			for _, elf := range elves.list {
				if elf.WillMove(&elves) {
					elf.ProposeMove(&elves, &moves, round)
					updates++
				}
			}

			if updates == 0 {
				return round + 1
			}

			for point, proposed := range moves {
				if len(proposed) == 1 {
					proposed[0].Update(point)
				}
			}

			elves.UpdateLocations()
		}
	}

	// for i := 0; ; i++ {
	// 	moves := NewMoveMap()
	// 	updates := 0

	// 	for _, elf := range elves.list {
	// 		if elf.WillMove(&elves) {
	// 			updates++
	// 			elf.ProposeMove(&elves, &moves, i)
	// 		}
	// 	}

	// 	if updates == 0 {
	// 		fmt.Println(i + 1)
	// 		panic("End?")
	// 	}

	// 	// fmt.Println("Round:", i)

	// }

	// elves.Print()

	return 0
}
