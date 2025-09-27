package tools

import (
	"os"
	"regexp"
)

type Point struct{ X, Y int }

func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

func ManhattanDistance(a Point, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
