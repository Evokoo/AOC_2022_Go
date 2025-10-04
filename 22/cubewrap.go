package day22

import "fmt"

//Add Bearings

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
		toB := Point{4 * n, 4*n - i}
		fmt.Printf("C -> F: %v -> %v\n", fromA, toB)

		fromB := Point{4*n + 1, 2*n + i}
		toA := Point{3 * n, n - i}
		fmt.Printf("F -> C: %v -> %v\n", fromB, toA)
	}
}
