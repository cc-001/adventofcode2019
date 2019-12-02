package day02

import (
	"reflect"
	"testing"
)

func Test_day02_1(t *testing.T) {
	in := Day02_parse_str("1,0,0,0,99")
	ans := Day02_parse_str("2,0,0,0,99")
	in = Day02_part1_solve(in)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day02_2(t *testing.T) {
	in := Day02_parse_str("2,3,0,3,99")
	ans := Day02_parse_str("2,3,0,6,99")
	in = Day02_part1_solve(in)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day02_3(t *testing.T) {
	in := Day02_parse_str("2,4,4,5,99,0")
	ans := Day02_parse_str("2,4,4,5,99,9801")
	in = Day02_part1_solve(in)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day02_4(t *testing.T) {
	in := Day02_parse_str("1,1,1,4,99,5,6,0,99")
	ans := Day02_parse_str("30,1,1,4,2,5,6,0,99")
	in = Day02_part1_solve(in)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}