package day06

import (
	"github.com/Evokoo/AOC_2022_Go/tools"
)

func Solve(file string, windowSize int) int {
	data := tools.ReadFile(file)
	return locateMarker(data, windowSize)
}

func locateMarker(data string, size int) int {
outer:
	for i := 0; i < len(data)-size; i++ {
		charMap := make(map[rune]struct{}, size)

		for _, char := range data[i : i+size] {
			if _, found := charMap[char]; found {
				continue outer
			} else {
				charMap[char] = struct{}{}
			}
		}
		return i + size
	}

	panic("Marker not found")
}
