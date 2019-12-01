package day01

import "testing"

func Test_day01_1(t *testing.T) {
	input := [4]int{12, 14, 1969, 100756}
	answers := [4]int{2, 2, 654, 33583}
	for i := 0; i < len(input); i++ {
		t.Logf("testing input: %d", input[i])
		tmp := Day01_part1_solve(input[i:i+1])
		if  tmp != answers[i] {
			t.Errorf("Failed test %d, got: %d, want: %d", i, tmp, answers[i])
		}
	}
}

func Test_day01_2(t *testing.T) {
	input := [3]int{14, 1969, 100756}
	answers := [3]int{2, 966, 50346}
	for i := 0; i < len(input); i++ {
		t.Logf("testing input: %d", input[i])
		tmp := Day01_part2_solve(input[i:i+1])
		if  tmp != answers[i] {
			t.Errorf("Failed test %d, got: %d, want: %d", i, tmp, answers[i])
		}
	}
}