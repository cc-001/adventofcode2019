package day09

import (
	"advent2019/day02"
	"advent2019/day05"
	"fmt"
	"reflect"
	"testing"
)

func Test_day09_1(t *testing.T) {
	in := day02.Day02_parse_str("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	output := make([]int, 0)
	day05.Day05_solve_adapter(in, 0, &output)
	if reflect.DeepEqual(in, output) == false {
		t.Errorf("Failed test, got: %v, want: %v", output, in)
	}
}

func Test_day09_2(t *testing.T) {
	in := day02.Day02_parse_str("1102,34915192,34915192,7,4,7,99,0")
	output := make([]int, 0)
	day05.Day05_solve_adapter(in, 0, &output)
	str := fmt.Sprintf("%d", output[0])
	if len(str) != 16 {
		t.Errorf("Failed test, got: %v, want: %v", len(str), 16)
	}
}

func Test_day09_3(t *testing.T) {
	in := day02.Day02_parse_str("104,1125899906842624,99")
	output := make([]int, 0)
	day05.Day05_solve_adapter(in, 0, &output)
	if output[0] != 1125899906842624 {
		t.Errorf("Failed test, got: %v, want: %v", output[0], 1125899906842624)
	}
}

func Test_day09_4(t *testing.T) {
	in := day02.Day02_parse("day09.txt")
	output := make([]int, 0)
	day05.Day05_solve_adapter(in, 1, &output)
	if len(output) != 1 || output[0] != 3235019597 {
		t.Errorf("Failed test, got: %v, want: %v", output[0], 3235019597)
	}
}

func Test_day09_5(t *testing.T) {
	in := day02.Day02_parse("day09.txt")
	output := make([]int, 0)
	day05.Day05_solve_adapter(in, 2, &output)
	if len(output) != 1 || output[0] != 80274 {
		t.Errorf("Failed test, got: %v, want: %v", output[0], 80274)
	}
}