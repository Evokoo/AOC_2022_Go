package day20_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2022_Go/20"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "example.txt", target: 3},
	{part: 1, file: "input.txt", target: 8302},
	{part: 2, file: "example.txt", target: 1623178306},
	{part: 2, file: "input.txt", target: 656575624777},
}

var _ = Describe("AOC 2022 - Day 20", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
