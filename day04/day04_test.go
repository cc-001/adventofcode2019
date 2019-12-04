package day04

import (
	"testing"
)

func Test_day04_1(t *testing.T) {
	ans := Day04_solve(273025, 767253, false)
	if ans != 910 {
		t.Errorf("Failed test, got: %v, want: %v", ans, 910)
	}

	ans = Day04_solve(273025, 767253, true)
	if ans != 598 {
		t.Errorf("Failed test, got: %v, want: %v", ans, 598)
	}
}