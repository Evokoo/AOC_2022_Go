package day22

func generate3DWrapMap(n int, grid Grid) {
	for i := 1; i <= n; i++ {
		// A
		grid.AddWrap(Point{2*n + i, 0}, Point{i, 4 * n}, 0, 0)              // A Top edge -> F
		grid.AddWrap(Point{3*n + 1, i}, Point{2 * n, 3*n + 1 - i}, 90, 180) // A Right edge -> D
		grid.AddWrap(Point{2*n + i, n + 1}, Point{2 * n, n + i}, 180, 90)   // A Bottom edge -> C
		// B
		grid.AddWrap(Point{n + i, 0}, Point{1, 3*n + i}, 0, 90)    // B Top edge -> F
		grid.AddWrap(Point{n, i}, Point{1, 3*n + 1 - i}, 270, 180) // B Left edge -> E
		// C
		grid.AddWrap(Point{n, n + i}, Point{i, 2*n + 1}, 270, -90)      // C Left edge -> E
		grid.AddWrap(Point{2*n + 1, n + i}, Point{2*n + i, n}, 90, -90) // C Right edge -> A
		// D
		grid.AddWrap(Point{n*2 + 1, 2*n + i}, Point{3 * n, n + 1 - i}, 90, 180) // D Right edge -> A
		grid.AddWrap(Point{n + i, 3*n + 1}, Point{n, 3*n + i}, 180, 90)         // D Bottom edge -> F
		// E
		grid.AddWrap(Point{i, 2 * n}, Point{n + 1, n + i}, 0, 90)          // E Top edge -> C
		grid.AddWrap(Point{0, 2*n + i}, Point{n + 1, n + 1 - i}, 270, 180) // E Left edge -> B
		// F
		grid.AddWrap(Point{0, 3*n + i}, Point{n + i, 1}, 270, -90)        // F Left edge -> B
		grid.AddWrap(Point{i, 4*n + 1}, Point{2*n + i, 1}, 180, 0)        // F Bottom edge -> A
		grid.AddWrap(Point{n + 1, 3*n + i}, Point{n + i, 3 * n}, 90, -90) // F right edge -> D
	}
}

func generate2DWrapMap(n int, grid Grid) {
	for i := 1; i <= n; i++ {
		// A
		grid.AddWrap(Point{2*n + i, 0}, Point{2*n + i, n}, 0, 0)       // A Top edge
		grid.AddWrap(Point{3*n + 1, i}, Point{n + 1, i}, 90, 0)        // A Right edge
		grid.AddWrap(Point{2*n + i, n + 1}, Point{2*n + i, 1}, 180, 0) // A Bottom edge
		// B
		grid.AddWrap(Point{n + i, 0}, Point{n + i, 3 * n}, 0, 0) // B Top edge
		grid.AddWrap(Point{n, i}, Point{n * 3, i}, 270, 0)       // B Left edge
		// C
		grid.AddWrap(Point{n, n + i}, Point{2 * n, n + i}, 270, 0)  // C Left edge
		grid.AddWrap(Point{2*n + 1, n + i}, Point{n, n + i}, 90, 0) // C Right edge
		// D
		grid.AddWrap(Point{n*2 + 1, 2*n + i}, Point{1, 2*n + i}, 90, 0) // D Right edge
		grid.AddWrap(Point{n + i, 3*n + 1}, Point{n + i, 1}, 180, 0)    // D Bottom edge
		// E
		grid.AddWrap(Point{i, 2 * n}, Point{i, 4 * n}, 0, 0)           // E Top edge
		grid.AddWrap(Point{0, 2*n + i}, Point{2 * n, 2*n + i}, 270, 0) // E Left edge
		// F
		grid.AddWrap(Point{0, 3*n + i}, Point{n, 3*n + i}, 270, 0)    // F Left edge
		grid.AddWrap(Point{i, 4*n + 1}, Point{i, 2*n + 1}, 180, 0)    // F Bottom edge
		grid.AddWrap(Point{n + 1, 3*n + i}, Point{1, 3*n + i}, 90, 0) // F right edge
	}
}
