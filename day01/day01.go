package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day01_parse() []int {
	file, err := os.Open("day01/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			input = append(input, i)
		}
	}
	return input
}

func Day01_part1_solve(inputs []int) int {
	sum := 0
	for _, i := range inputs {
		sum += (i / 3) - 2
	}
	return sum
}

func Day01_part2_solve(inputs []int) int {
	sum := 0
	for _, i := range inputs {
		prev := i
		for {
			fuel := (prev / 3) - 2
			if fuel <= 0 {
				break;
			} else {
				sum += fuel
				prev = fuel
			}
		}
	}
	return sum
}