package day06_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2022_Go/06"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "exampleA.txt", target: 7},
	{part: 1, file: "exampleB.txt", target: 5},
	{part: 1, file: "input.txt", target: 1538},
	{part: 2, file: "exampleA.txt", target: 19},
	{part: 2, file: "exampleB.txt", target: 23},
	{part: 2, file: "input.txt", target: 2315},
}

var _ = Describe("AOC 2022 - Day 05", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			windowSize := 4
			if test.part == 2 {
				windowSize += 10
			}
			result := Solve(test.file, windowSize)
			Expect(result).To(Equal((test.target)))
		})
	}
})
