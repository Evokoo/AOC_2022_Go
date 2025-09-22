package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// Solution
func Solve(file string) int {
	data := tools.ReadFile(file)
	CPU := newCPU()
	CRT := newCRT()

	for line := range strings.SplitSeq(data, "\n") {
		command := newCommand(line)
		for range command.cycles {
			CRT.drawPixel(CPU, "+")
			CPU.recordSignal()
			CPU.tick()
		}
		CPU.updateRegister(command.value)
	}

	// Part 2 solution
	CRT.print()

	// Part 1 solution
	return CPU.signal
}

// CPU
type CPU struct {
	register int
	cycle    int
	signal   int
	target   int
}

func newCPU() CPU {
	return CPU{1, 0, 0, 20}
}
func (cpu *CPU) recordSignal() {
	if (cpu.cycle + 1) == cpu.target {
		cpu.signal += cpu.register * (cpu.cycle + 1)
		cpu.target += 40
	}
}
func (cpu *CPU) tick() {
	cpu.cycle++
}
func (cpu *CPU) updateRegister(value int) {
	cpu.register += value
}

// CRT
type CRT [][]string

func newCRT() CRT {
	pixels := make([][]string, 6)

	for i := range pixels {
		row := make([]string, 40)
		for j := range row {
			row[j] = " "
		}
		pixels[i] = row
	}

	return pixels
}
func (crt *CRT) drawPixel(cpu CPU, pixel string) {
	row := cpu.cycle / 40
	col := cpu.cycle % 40

	if col >= cpu.register-1 && col <= cpu.register+1 {
		(*crt)[row][col] = pixel
	}
}
func (crt *CRT) print() {
	for _, row := range *crt {
		fmt.Println(strings.Join(row, ""))
	}
}

// Commands
type Command struct {
	cycles int
	value  int
}

func newCommand(line string) Command {
	parts := strings.Split(line, " ")
	if parts[0][0] == 'n' {
		return Command{1, 0}
	} else {
		value, _ := strconv.Atoi(parts[1])
		return Command{2, value}
	}
}
