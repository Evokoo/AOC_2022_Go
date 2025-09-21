package day09_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2022_Go/09"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "exampleA.txt", target: 13},
	{part: 1, file: "input.txt", target: 6044},
	{part: 2, file: "exampleB.txt", target: 36},
	{part: 2, file: "input.txt", target: 2384},
}

var _ = Describe("AOC 2022 - Day 09", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
