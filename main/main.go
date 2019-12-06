package main

import (
	"advent2019/day02"
	"advent2019/day05"
	"fmt"
)

func main() {
	in := day02.Day02_parse("day05/day05.txt")
	output := make([]int, 0)
	day05.Day05_solve(in, 5, &output)
	fmt.Printf("part2: %v", output)
}