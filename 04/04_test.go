package day04_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	day04 "github.com/Evokoo/AOC_2022_Go/04"
)

type Test struct {
	part   int
	file   string
	target int
}

var tests = []Test{
	{part: 1, file: "example.txt", target: 2},
	{part: 1, file: "input.txt", target: 530},
	{part: 2, file: "example.txt", target: 4},
	{part: 2, file: "input.txt", target: 903},
}

var _ = Describe("AOC 2022 - Day 04", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := day04.Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
