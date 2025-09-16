package main

import (
	"fmt"
	"testing"
)

type Test struct {
	part   string
	file   string
	elves  int
	target int
}

var tests = []Test{
	{part: "Example", file: "example.txt", elves: 1, target: 24000},
	{part: "Part One", file: "input.txt", elves: 1, target: 66186},
	{part: "Part Two", file: "input.txt", elves: 3, target: 196804},
}

func TestAOC(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("\nTesting %s with %s", tc.part, tc.file), func(t *testing.T) {
			result := Solve(tc.file, tc.elves)
			if result != tc.target {
				t.Fatalf("Expected %d to equal %d", result, tc.target)
			}
		})
	}
}
