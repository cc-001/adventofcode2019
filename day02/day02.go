package day02

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Day02_parse_str(s string) []int {
	strs := strings.Split(s, ",")
	input := make([]int, 0, len(strs))
	for _, str := range strs {
		if i, err := strconv.Atoi(str); err == nil {
			input = append(input, i)
		}
	}
	return input
}

func Day02_parse(fileName string) []int {
	buf, _ := ioutil.ReadFile(fileName)
	return Day02_parse_str(string(buf))
}

func Day02_part1_solve(program []int) []int {
	pc := 0
	for {
		switch program[pc] {
		case 1:
			program[program[pc+3]] = program[program[pc+1]] + program[program[pc+2]]
		case 2:
			program[program[pc+3]] = program[program[pc+1]] * program[program[pc+2]]
		case 99:
			return program
		}
		pc += 4
	}
}

func Day02_part2_solve(program []int) int {
	// stupid O(N^2), stupid memory, stupid
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			input := make([]int, len(program))
			copy(input, program)
			input[1] = noun
			input[2] = verb
			output := Day02_part1_solve(input)
			if output[0] == 19690720 {
				return 100 * noun + verb
			}
		}
	}
	return -1
}