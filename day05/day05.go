package day05

import (
	"fmt"
	"strconv"
)

func fetch(param int, mode [2]int, pc int, memory *[]int) int {
	if mode[param-1] > 0 {
		return (*memory)[pc+param]
	}
	return (*memory)[(*memory)[pc+param]]
}

func Day05_solve(memory []int, input int, output *[]int) []int {
	pc := 0
	for {
		instruction := fmt.Sprintf("%04d", memory[pc])
		opcode, _ := strconv.Atoi(instruction[2:])
		mode := [2]int{}
		for i := 1; i >= 0; i-- {
			if instruction[i:i+1] == "1" {
				mode[1-i] = 1;
			}
		}

		switch opcode {
		case 1:
			memory[memory[pc+3]] = fetch(1, mode, pc, &memory) + fetch(2, mode, pc, &memory)
			pc += 4
		case 2:
			memory[memory[pc+3]] = fetch(1, mode, pc, &memory) * fetch(2, mode, pc, &memory)
			pc += 4
		case 3:
			memory[memory[pc+1]] = input
			pc += 2
		case 4:
			*output = append(*output, fetch(1, mode, pc, &memory))
			pc += 2
		case 5:
			val := fetch(1, mode, pc, &memory)
			if val != 0 {
				pc = fetch(2, mode, pc, &memory)
			} else {
				pc += 3
			}
		case 6:
			val := fetch(1, mode, pc, &memory)
			if val == 0 {
				pc = fetch(2, mode, pc, &memory)
			} else {
				pc += 3
			}
		case 7:
			a := fetch(1, mode, pc, &memory)
			b := fetch(2, mode, pc, &memory)
			if a < b {
				memory[memory[pc+3]] = 1
			} else {
				memory[memory[pc+3]] = 0
			}
			pc += 4
		case 8:
			a := fetch(1, mode, pc, &memory)
			b := fetch(2, mode, pc, &memory)
			if a == b {
				memory[memory[pc+3]] = 1
			} else {
				memory[memory[pc+3]] = 0
			}
			pc += 4
		case 99:
			return memory
		}
	}
}
