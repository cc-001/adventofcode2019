package day05

import (
	"fmt"
	"strconv"
)

func expand(addr int, memory *[]int) {
	if addr >= len(*memory) {
		delta := (addr - len(*memory)) + 1
		*memory = append(*memory, make([]int, delta)...)
		// log.Printf("expanding memory by: %d", delta)
	}
}

func get_addr(param int, mode [3]int, pc int, rb int, memory *[]int, write bool) int {
	var read_addr, addr int
	read_addr = pc+param
	expand(read_addr, memory)
	addr = (*memory)[read_addr]
	if mode[param-1] == 2 {
		addr += rb
	}
	if write {
		expand(addr, memory)
	}
	return addr
}

func fetch(param int, mode [3]int, pc int, rb int, memory *[]int) int {
	addr := get_addr(param, mode, pc, rb, memory, false)
	if mode[param-1] == 1 {
		return addr
	}
	if addr >= len(*memory) {
		return 0
	}
	return (*memory)[addr]
}

func Day05_solve_adapter(memory []int, input int, output *[]int) []int {
	new_input := make([]int, 0, 1)
	new_input = append(new_input, input)
	return Day05_solve(memory, new_input, output, nil, nil,true, nil)
}

func Day05_solve(memory []int, input []int, output *[]int, pc_out *int, rb_out *int, inc_input bool, get func()int) []int {
	input_ctr := 0
	pc := 0
	rb := 0
	if pc_out != nil {
		pc = *pc_out
	}
	if rb_out != nil {
		rb = *rb_out
	}
	for {
		instruction := fmt.Sprintf("%05d", memory[pc])
		opcode, _ := strconv.Atoi(instruction[3:])
		mode := [3]int{}
		for i := 2; i >= 0; i-- {
			ch := instruction[i:i+1]
			if ch == "1" {
				mode[2-i] = 1;
			} else if ch == "2" {
				mode[2-i] = 2
			} else {
				mode[2-i] = 0
			}
		}

		switch opcode {
		case 1:
			addr := get_addr(3, mode, pc, rb, &memory, true)
			memory[addr] = fetch(1, mode, pc, rb, &memory) + fetch(2, mode, pc, rb, &memory)
			pc += 4
		case 2:
			addr := get_addr(3, mode, pc, rb, &memory, true)
			memory[addr] = fetch(1, mode, pc, rb, &memory) * fetch(2, mode, pc, rb, &memory)
			pc += 4
		case 3:
			addr := get_addr(1, mode, pc, rb, &memory, true)
			if get != nil {
				memory[addr] = get()
			} else {
				memory[addr] = input[input_ctr]
			}
			if inc_input {
				input_ctr++
			}
			pc += 2
		case 4:
			*output = append(*output, fetch(1, mode, pc, rb, &memory))
			pc += 2
			if rb_out != nil {
				*rb_out = rb
			}
			if pc_out != nil {
				// suspend
				*pc_out = pc
				return memory
			}
		case 5:
			val := fetch(1, mode, pc, rb, &memory)
			if val != 0 {
				pc = fetch(2, mode, pc, rb, &memory)
			} else {
				pc += 3
			}
		case 6:
			val := fetch(1, mode, pc, rb, &memory)
			if val == 0 {
				pc = fetch(2, mode, pc, rb, &memory)
			} else {
				pc += 3
			}
		case 7:
			a := fetch(1, mode, pc, rb, &memory)
			b := fetch(2, mode, pc, rb, &memory)
			addr := get_addr(3, mode, pc, rb, &memory, true)
			if a < b {
				memory[addr] = 1
			} else {
				memory[addr] = 0
			}
			pc += 4
		case 8:
			a := fetch(1, mode, pc, rb, &memory)
			b := fetch(2, mode, pc, rb, &memory)
			addr := get_addr(3, mode, pc, rb, &memory, true)
			if a == b {
				memory[addr] = 1
			} else {
				memory[addr] = 0
			}
			pc += 4
		case 9:
			val := fetch(1, mode, pc, rb, &memory)
			rb += val
			pc += 2
		case 99:
			if pc_out != nil {
				// terminate signal
				*pc_out = -1
			}
			return memory
		}
	}
}