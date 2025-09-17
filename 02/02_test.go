package main

import (
	"fmt"
	"testing"
)

type Test struct {
	part   int
	file   string
	target int
}

var tests = []Test{
	{part: 1, file: "example.txt", target: 15},
	{part: 1, file: "input.txt", target: 14297},
	{part: 2, file: "example.txt", target: 12},
	{part: 2, file: "input.txt", target: 10498},
}

func TestAOC(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("\nTesting part %d with %s", tc.part, tc.file), func(t *testing.T) {
			result := Solve(tc.file, tc.part)
			if result != tc.target {
				t.Fatalf("Expected %d to equal %d", result, tc.target)
			}
		})
	}
}
