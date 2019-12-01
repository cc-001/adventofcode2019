package main

import (
	"advent2019/day01"
	"fmt"
)

func main() {
	part1_inputs := day01.Day01_parse()
	fmt.Printf("Day01_part1_solve: %d\n", day01.Day01_part1_solve(part1_inputs))
	fmt.Printf("Day01_part2_solve: %d\n", day01.Day01_part2_solve(part1_inputs))
}