package day25

import (
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// PARSER
// ========================
func GetInputSum(file string) int {
	data := tools.ReadFile(file)
	sum := 0
	for code := range strings.SplitSeq(data, "\n") {
		sum += ToDecimal(code)
	}
	return sum
}

// ========================
// DECIMAL
// ========================
var Convert = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'-': -1,
	'=': -2,
}

func ToDecimal(n string) int {
	sum, col := 0, 1

	for i := len(n) - 1; i >= 0; i-- {
		sum += Convert[n[i]] * col
		col *= 5
	}
	return sum
}

// ========================
// SNAFU
// ========================
func ToSNAFU(n int) string {
	if n == 0 {
		return "0"
	}

	digits := []byte{}
	for n != 0 {
		r := n % 5
		n /= 5
		if r > 2 {
			r -= 5
			n += 1
		}

		switch r {
		case -2:
			digits = append([]byte{'='}, digits...)
		case -1:
			digits = append([]byte{'-'}, digits...)
		case 0:
			digits = append([]byte{'0'}, digits...)
		case 1:
			digits = append([]byte{'1'}, digits...)
		case 2:
			digits = append([]byte{'2'}, digits...)
		}
	}

	return string(digits)
}
