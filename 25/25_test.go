package day25_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2022_Go/25"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target string
}

var tests = []Test{
	{part: 1, file: "example.txt", target: "2=-1=0"},
	{part: 1, file: "input.txt", target: "20=02=120-=-2110-0=1"},
	// {part: 2, file: "example.txt", target: -1},
	// {part: 2, file: "input.txt", target: -1},
}

var _ = Describe("AOC 2022 - Day 25", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file)
			Expect(result).To(Equal((test.target)))
		})
	}
})
