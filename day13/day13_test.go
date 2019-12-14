package day13

import "testing"

func Test_day13_part1(t *testing.T) {
	ans := Day13_solve_part1()
	if ans != 253 {
		t.Errorf("Failed test, got: %v, want: 253", ans)
	}
}

func Test_day13_part2(t *testing.T) {
	ans := Day13_solve_part2()
	if ans != 12263 {
		t.Errorf("Failed test, got: %v, want: 12263", ans)
	}
}
