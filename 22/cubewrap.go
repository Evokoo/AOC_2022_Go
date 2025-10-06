package day22

import "fmt"

//Add Bearings

// type Wrap struct {
// 	from    Point
// 	to      Point
// 	bearing int
// }

func AC(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{i + 3*n, 0 - 1}
		toB := Point{n - i, n}
		fmt.Printf("C -> A: %v -> %v\n", fromA, toB)

		fromB := Point{i, n - 1}
		toA := Point{4*n - i, 0}
		fmt.Printf("A -> C: %v -> %v\n", fromB, toA)
	}
}

func AE(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{0 + i, 2*n + 1}
		toB := Point{3*n - i, 3 * n}
		fmt.Printf("A -> E: %v -> %v\n", fromA, toB)

		fromB := Point{2*n + i, 3*n + 1}
		toA := Point{n - i, 2 * n}
		fmt.Printf("E -> A: %v -> %v\n", fromB, toA)
	}
}

func BC(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{2*n - 1, i}
		toB := Point{i + n, n}
		fmt.Printf("C -> B: %v -> %v\n", fromA, toB)

		fromB := Point{i + n, i - 1}
		toA := Point{2 * n, i}
		fmt.Printf("B -> C: %v -> %v\n", fromB, toA)
	}
}

func BE(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{n + i, n*2 + 1}
		toB := Point{2 * n, 4*n - i}
		fmt.Printf("B -> E: %v -> %v\n", fromA, toB)

		fromB := Point{2*n - 1, 2*n + i}
		toA := Point{2*n - i, 2 * n}
		fmt.Printf("E -> B: %v -> %v\n", fromB, toA)
	}
}

func DF(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{3*n + 1, n + i}
		toB := Point{4*n - i, 2 * n}
		fmt.Printf("D -> F: %v -> %v\n", fromA, toB)

		fromB := Point{3*n + i, 2*n - 1}
		toA := Point{3 * n, 2*n - i}
		fmt.Printf("F -> D: %v -> %v\n", fromB, toA)
	}
}

func AF(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{0 - 1, n + i}
		toB := Point{4*n - i, 3 * n}
		fmt.Printf("A -> F: %v -> %v\n", fromA, toB)

		fromB := Point{3*n + i, 3*n + 1}
		toA := Point{0, 2*n - i}
		fmt.Printf("F -> A: %v -> %v\n", fromB, toA)
	}
}

func CF(n int) {
	for i := 0; i <= n; i++ {
		fromA := Point{n*3 + 1, i}
		toB := Point{4 * n, 3*n - i}
		fmt.Printf("C -> F: %v -> %v\n", fromA, toB)

		// fromB := Point{4*n + 1, 2*n + i}
		// toA := Point{3 * n, n - i}
		// fmt.Printf("F -> C: %v -> %v\n", fromB, toA)
	}
}

// Flat Wraps

// func generateExampleFlatMap(n int, grid Grid) {
// 	for i := 1; i <= n; i++ {
// 		// A
// 		grid.AddWrap(Point{0, n + i}, Point{3 * n, n + i}, 0) // A left edge -> D
// 		grid.AddWrap(Point{i, n}, Point{i, 2 * n}, 0)         // A Top edge -> A
// 		grid.AddWrap(Point{i, 2*n + 1}, Point{i, n}, 0)       // A Bottom edge -> A
// 		// B
// 		grid.AddWrap(Point{n + i, n}, Point{n + i, 2 * n}, 0)   // B Top edge -> B
// 		grid.AddWrap(Point{n + i, 2*n + 1}, Point{n + i, n}, 0) // B Bottom edge -> B
// 		// C
// 		grid.AddWrap(Point{2 * n, i}, Point{3 * n, i}, 0)         // C Left edge -> C
// 		grid.AddWrap(Point{3*n + 1, i}, Point{2 * n, i}, 0)       // C Right edge -> C
// 		grid.AddWrap(Point{2*n + i, 0}, Point{2*n + i, 3 * n}, 0) // C Top edge -> E
// 		// D
// 		grid.AddWrap(Point{3*n + 1, n + i}, Point{1, n + i}, 0) // D Right edge -> A
// 		// E
// 		grid.AddWrap(Point{2*n + 1, 2*n + i}, Point{4 * n, 2*n + i}, 0) // E Left edge -> F
// 		grid.AddWrap(Point{2*n + i, 3 * n}, Point{2*n + i, 1}, 0)       // E Bottom edge -> C
// 		// F
// 		grid.AddWrap(Point{3*n + i, 2 * n}, Point{3*n + i, 3 * n}, 0)   // F Top edge -> F
// 		grid.AddWrap(Point{3*n + i, 3 * n}, Point{3*n + i, 2 * n}, 0)   // F Bottom edge -> F
// 		grid.AddWrap(Point{4*n + 1, 2*n + i}, Point{2 * n, 2*n + i}, 0) // F Right edge -> E
// 	}
// }

func generateExampleCubeMap(n int, grid Grid) {
	// n += 1
	for i := 1; i <= n; i++ {
		// A
		// grid.AddWrap(Point{i, n - 1}, Point{4*n - i, 0}, 0, 180)         // A Top edge -> C
		// grid.AddWrap(Point{-1, n + i}, Point{4*n - i, 3 * n}, 270, 90)   // A Left edge -> F
		// grid.AddWrap(Point{i, 2*n + 1}, Point{3*n - i, 3 * n}, 180, 180) // A Bottom edge -> E
		// // B
		// grid.AddWrap(Point{n + i, n}, Point{2*n + 1, i}, 0, 90)              // B Top edge -> C
		// grid.AddWrap(Point{n + i, 2*n + 1}, Point{2 * n, 3*n - i}, 180, -90) // B Bottom edge -> E
		// // // C
		// // grid.AddWrap(Point{2*n + i, 0 - 1}, Point{n - i, n}, 0) // C Top edge -> A
		// grid.AddWrap(Point{2 * n, i}, Point{n + i, n + 1}, 270, -90) // C Left edge -> B
		// // grid.AddWrap(Point{3 * n, i}, Point{4 * n, 3*n - i}, 0) // C Right edge -> F
		// // D
		// grid.AddWrap(Point{3*n + 1, n + i}, Point{4*n + 1 - i, 2*n + 1}, 90, 90) // D Right edge -> F
		// // E
		// grid.AddWrap(Point{2*n - 1, 2*n + i}, Point{2*n - i, 2 * n}, 270, 90)    // E Left edge -> B
		// grid.AddWrap(Point{2*n + i, 3*n + 1}, Point{n + 1 - i, 2 * n}, 180, 180) // E Bottom edge -> A
		// // F
		// grid.AddWrap(Point{3*n + i, 2 * n}, Point{3 * n, 2*n - i}, 0) // F Top edge -> D
		// grid.AddWrap(Point{4 * n, 2*n + i}, Point{3 * n, n - i}, 0)   // F Right edge -> C
		// grid.AddWrap(Point{3 * n, 3*n + 1}, Point{0, 2*n - i}, 0)     // F Bottom edge -> A
	}
}

func generateInputCubeMap(n int, grid Grid) {
	// n += 1
	for i := 1; i <= n; i++ {
		// A
		grid.AddWrap(Point{2*n + 1 + i, 0}, Point{2 * n, 4*n + 1 - i}, 0, -90) // A Top edge -> E
		grid.AddWrap(Point{3*n + 1, i}, Point{2 * n, 3*n + 1 - i}, 90, 180)    // A Right edge -> D
		grid.AddWrap(Point{2*n + 1 + i, n + 1}, Point{2 * n, n + i}, 180, 90)  // A Bottom edge -> C
		// B
		grid.AddWrap(Point{n + 1 + i, 0}, Point{n + 1 + i, 4 * n}, 0, 0) // B Top edge -> E
		grid.AddWrap(Point{n, i}, Point{n + 1 - i, 4 * n}, 270, 90)      // B Left edge -> F
		// C
		grid.AddWrap(Point{n, n + 1 + i}, Point{1, 4*n + 1 - i}, 270, 180)      // C Left edge -> F
		grid.AddWrap(Point{2*n + 1, n + 1 + i}, Point{2*n + 1 + i, n}, 90, -90) // C Right edge -> A
		// D
		grid.AddWrap(Point{n, 2*n + 1 + i}, Point{i, 3*n + 1}, 270, -90)            // D Left edge -> F
		grid.AddWrap(Point{2*n + 1, 2*n + 1 + i}, Point{3 * n, n + 1 - i}, 90, 180) // D Right edge -> A
		// E
		grid.AddWrap(Point{2*n + 1, 3*n + i}, Point{3*n + 1 - i, 1}, 90, 90) // E Right edge -> A
		grid.AddWrap(Point{n + 1 + i, 4*n + 1}, Point{n + 1 + i, 1}, 180, 0) // E Bottom edge -> B
		// F
		grid.AddWrap(Point{i, 4*n + 1}, Point{n + 1, n + 1 - i}, 180, -90) // F Bottom edge -> B
		grid.AddWrap(Point{0, 3 * n}, Point{n + 1, 2*n + 1 - i}, 270, 180) // F Left edge -> C
		grid.AddWrap(Point{i, 3 * n}, Point{n + 1, 4*n + 1 - i}, 0, 90)    // F Top edge -> D
	}
}

// To Do
// Add exit bearing to wraps to ensure no overlapping clashes
