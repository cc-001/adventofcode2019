package day11

import "testing"

func Test_day11_part1(t *testing.T) {
	out := Day11_solve_part1("day11.txt")
	if out != 2064 {
		t.Errorf("Failed test, got: %v, want: 2064", out)
	}
}

func Test_day11_part2(t *testing.T) {
	Day11_solve_part2("day11.txt", "part2.png")
}