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
	{part: 1, file: "example.txt", target: 157},
	{part: 1, file: "input.txt", target: 8240},
	{part: 2, file: "example.txt", target: 70},
	{part: 2, file: "input.txt", target: 2587},
}

func TestAOC(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testing part %d with %s", tc.part, tc.file), func(t *testing.T) {
			var result int
			if tc.part == 1 {
				result = SolveOne(tc.file)
			} else {
				result = SolveTwo(tc.file)
			}
			if result != tc.target {
				t.Fatalf("Expected %d to equal %d", result, tc.target)
			}
		})
	}
}
