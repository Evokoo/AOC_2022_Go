package day10

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type CPU struct {
	register       int
	cycle          int
	signalStrength int
}
type Command struct {
	cycles int
	value  int
}

func Solve(file string) int {
	cpu := newCPU()
	target := 20

	for _, command := range parseCommands(file) {
		for range command.cycles {
			if cpu.cycle == target {
				cpu.signalStrength += cpu.register * cpu.cycle
				target += 40
			}
			cpu.cycle++
		}

		cpu.register += command.value
	}

	return cpu.signalStrength
}

func newCPU() CPU {
	return CPU{1, 1, 0}
}
func parseCommands(file string) (commands []Command) {
	data := tools.ReadFile(file)

	for line := range strings.SplitSeq(data, "\n") {
		parts := strings.Split(line, " ")
		if parts[0][0] == 'n' {
			commands = append(commands, Command{1, 0})
		} else {
			value, _ := strconv.Atoi(parts[1])
			commands = append(commands, Command{2, value})
		}
	}
	return
}
