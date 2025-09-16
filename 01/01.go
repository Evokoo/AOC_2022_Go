package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

func main() {
	fmt.Printf("EXAMPLE  :  %d\n", Solve("example.txt", 1))
	fmt.Printf("PART ONE :  %d\n", Solve("input.txt", 1))
	fmt.Printf("PART TWO :  %d\n", Solve("input.txt", 3))
}

func Solve(file string, n int) int {
	data := tools.ReadFile(file)
	elves := strings.Split(data, "\n\n")
	calories := make([]int, len(elves))

	for i, elf := range elves {
		var score int
		for _, s := range strings.Fields(elf) {
			if value, err := strconv.Atoi(s); err == nil {
				score += value
			}
		}
		calories[i] = score
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	var result int
	for _, amount := range calories[:n] {
		result += amount
	}

	return result
}
