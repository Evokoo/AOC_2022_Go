package day17

func Solve(file string) int {
	pile := NewPile()
	jets := ParseJets(file)
	l := len(jets)

	for i, j := 0, 0; i < 1000000000000; i++ {
		block := NewBlock(i%5, &pile)

		for {
			dx := jets[j%l]
			block.HorizontalShift(dx, &pile)
			j++

			if !block.VerticalShift(&pile) {
				break
			}
		}
	}

	return pile.height
}
