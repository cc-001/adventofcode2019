package day07

import (
	"advent2019/day05"
	"fmt"
	"math"
)

type slice []int
func (s slice) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func heap_permutation(a *slice, size int, n int, out *[]slice) {
	if size == 1 {
		val := make([]int, len(*a))
		copy(val, *a)
		*out = append(*out, val)
	}
	for i := 0; i < size; i++ {
		heap_permutation(a, size-1, n, out);
		if (size % 2 == 1) {
			a.swap(0, size-1)
		} else {
			a.swap(i, size-1)
		}
	}
}

type amplifier struct {
	suspend bool
	phase int
	memory []int
	output []int
	last_output int
	suspend_count int
	pc int
	rb int
}

func (a *amplifier) execute(input int) int {
	a.output = make([]int, 0)
	inputs := []int{a.phase, input}
	if a.suspend {
		if a.suspend_count > 0 {
			inputs[0] = input
		}
		a.memory = day05.Day05_solve(a.memory, inputs, &a.output, &a.pc, &a.rb, true, nil)
		a.suspend_count++
	} else {
		day05.Day05_solve(a.memory, inputs, &a.output, nil, nil,true, nil)
	}
	if len(a.output) > 0 {
		a.last_output = a.output[0]
	}
	return a.last_output
}

func Day07_solve_universal(program []int, phase_setting []int, suspend bool) int {
	amps := make([]amplifier, 0, len(phase_setting))
	for i := 0; i < len(phase_setting); i++ {
		amps = append(amps, amplifier{
			suspend: suspend,
			phase: phase_setting[i],
			memory: make([]int, len(program)),
		})
		copy(amps[i].memory, program)
	}

	input := 0
	done := false
	for {
		for i := 0; i < len(amps); i++ {
			input = amps[i].execute(input)
			if suspend && amps[i].pc == -1 {
				done = true
			}
		}
		if suspend {
			if done {
				return amps[len(amps)-1].last_output
			}
		} else {
			return input
		}
	}
}

func Day07_solve(program []int, num_amps int, initial_phase int, part2 bool) int {
	phase_seqs := make([]slice, 0)
	var initial slice
	for i := 0; i < num_amps; i++ {
		initial = append(initial, i + initial_phase)
	}
	heap_permutation(&initial, len(initial), len(initial), &phase_seqs)
	var thrust int
	best_thrust := math.MinInt32
	var best_seq slice
	for _, i := range phase_seqs {
		thrust = Day07_solve_universal(program, i, part2)
		if thrust > best_thrust {
			best_thrust = thrust
			best_seq = i
		}
	}
	fmt.Printf("best_seq: %v\nbest_thrust: %v\n", best_seq, best_thrust)
	return best_thrust
}