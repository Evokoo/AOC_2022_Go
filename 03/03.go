package day03

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

func main() {
	fmt.Printf("PART ONE :  %d\n", SolveOne("input.txt"))
	fmt.Printf("PART TWO :  %d\n", SolveTwo("input.txt"))
}

func SolveOne(file string) string {
	lines := strings.Split(tools.ReadFile(file), "\n")
	items := make(map[rune]int)
	for _, line := range lines {
		sack := make(map[rune]struct{})
		m := len(line) / 2

		//Check left
		for _, item := range line[:m] {
			sack[item] = struct{}{}
		}
		//Check right
		for _, item := range line[m:] {
			if _, found := sack[item]; found {
				items[item]++
				break
			}
		}
	}
	sum := 0
	for item, count := range items {
		sum += runeToIndex(item) * count
	}
	return sum
}

func SolveTwo(file string) int {
	lines := strings.Split(tools.ReadFile(file), "\n")
	sum := 0

	for i := 0; i < len(lines); i = i + 3 {
		group := lines[i : i+3]
		shared := make(map[rune]int)

		for _, elf := range group {
			set := make(map[rune]struct{})
			for _, r := range elf {
				set[r] = struct{}{}
			}
			for r := range set {
				shared[r]++
			}
		}

		for r, count := range shared {
			if count == 3 {
				sum += runeToIndex(r)
			}
		}
	}
	return sum
}

func runeToIndex(r rune) int {
	if !unicode.IsLetter(r) {
		panic("Invalid rune")
	}
	if unicode.IsLower(r) {
		return int(r-'a') + 1
	} else {
		return int(r-'A') + 27
	}
}
