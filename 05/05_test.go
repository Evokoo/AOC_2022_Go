package day05_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2022_Go/05"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	part   int
	file   string
	target string
}

var tests = []Test{
	{part: 1, file: "example.txt", target: "CMZ"},
	{part: 1, file: "input.txt", target: "TDCHVHJTG"},
	{part: 2, file: "example.txt", target: "MCD"},
	{part: 2, file: "input.txt", target: "NGCMPJLHV"},
}

var _ = Describe("AOC 2022 - Day 05", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
