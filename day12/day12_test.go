package day12

import "testing"

func Test_day12_1(t *testing.T) {
	ans := Day12_solve_part1("day12_1.txt", 11)
	if ans != 179 {
		t.Errorf("Failed test, got: %v, want: 179", ans)
	}
}

func Test_day12_2(t *testing.T) {
	ans := Day12_solve_part1("day12_2.txt", 101)
	if ans != 1940 {
		t.Errorf("Failed test, got: %v, want: 1940", ans)
	}
}

func Test_day12_part1(t *testing.T) {
	ans := Day12_solve_part1("day12_input.txt", 1001)
	if ans != 9493 {
		t.Errorf("Failed test, got: %v, want: 9493", ans)
	}
}

func Test_day12_3(t *testing.T) {
	ans := Day12_solve_part2("day12_1.txt")
	if ans != 2772 {
		t.Errorf("Failed test, got: %v, want: 2772", ans)
	}
}

func Test_day12_4(t *testing.T) {
	ans := Day12_solve_part2("day12_2.txt")
	if ans != 4686774924 {
		t.Errorf("Failed test, got: %v, want: 4686774924", ans)
	}
}

func Test_day12_part2(t *testing.T) {
	ans := Day12_solve_part2("day12_input.txt")
	t.Logf("ans: %v", ans)
}