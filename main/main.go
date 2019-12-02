package main

import (
	"advent2019/day02"
	"fmt"
)

func main() {
	part1_inputs := day02.Day02_parse("day02/day02.txt")
	part1_inputs[1] = 12
	part1_inputs[2] = 2
	fmt.Printf("Day02_part1_solve: %d\n", day02.Day02_part1_solve(part1_inputs)[0])
	part2_inputs := day02.Day02_parse("day02/day02.txt")
	fmt.Printf("Day02_part2_solve: %d\n", day02.Day02_part2_solve(part2_inputs))
}