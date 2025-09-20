package day07

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type FilePath []int

func (fp *FilePath) push(dir int) {
	*fp = append(*fp, dir)
}
func (fp *FilePath) pop() {
	*fp = (*fp)[:len(*fp)-1]
}

func Solve(file string, part int) int {
	data := tools.ReadFile(file)
	dirs := traverseFileTree(data)

	dir := math.MaxInt
	sum := 0

	for _, size := range dirs {
		if size <= 100000 {
			sum += size
		}

		if (70000000-dirs[0])+size >= 30000000 && size < dir {
			dir = size
		}
	}
	if part == 2 {
		return dir
	}
	return sum
}
func traverseFileTree(data string) map[int]int {
	var path FilePath
	var id int

	volumes := make(map[int]int)

	for command := range strings.SplitSeq(data, "\n") {
		c := strings.Fields(command)

		if c[0] == "$" && c[1] == "cd" {
			if c[2] == ".." {
				path.pop()
			} else {
				path.push(id)
				id++
			}
		}
		if size, err := strconv.Atoi(c[0]); err == nil {
			for _, dir := range path {
				volumes[dir] += size
			}
		}
	}
	return volumes
}
