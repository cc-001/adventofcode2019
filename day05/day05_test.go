package day05

import (
	"advent2019/day02"
	"reflect"
	"testing"
)

func Test_day05_1(t *testing.T) {
	in := day02.Day02_parse_str("1,0,0,0,99")
	ans := day02.Day02_parse_str("2,0,0,0,99")
	in = Day05_solve(in, 0, nil)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day05_2(t *testing.T) {
	in := day02.Day02_parse_str("2,3,0,3,99")
	ans := day02.Day02_parse_str("2,3,0,6,99")
	in = Day05_solve(in, 0, nil)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day05_3(t *testing.T) {
	in := day02.Day02_parse_str("2,4,4,5,99,0")
	ans := day02.Day02_parse_str("2,4,4,5,99,9801")
	in = Day05_solve(in, 0, nil)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day05_4(t *testing.T) {
	in := day02.Day02_parse_str("1,1,1,4,99,5,6,0,99")
	ans := day02.Day02_parse_str("30,1,1,4,2,5,6,0,99")
	in = Day05_solve(in, 0, nil)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day05_5(t *testing.T) {
	in := day02.Day02_parse_str("1002,4,3,4,33")
	ans := day02.Day02_parse_str("1002,4,3,4,99")
	in = Day05_solve(in, 0, nil)
	if reflect.DeepEqual(in, ans) == false {
		t.Errorf("Failed test, got: %v, want: %v", in, ans)
	}
}

func Test_day05_6(t *testing.T) {
	in := day02.Day02_parse("day05.txt")
	output := make([]int, 0)
	Day05_solve(in, 1, &output)
	for i := 0; i < len(output)-1; i++ {
		if output[i] != 0 {
			t.Errorf("Failed test, got: %d at position %d, want: 0", output[i], i)
		}
	}
	if output[len(output)-1] != 11193703 {
		t.Errorf("Failed test, got: %d, want: %d", output[len(output)-1] , 11193703)
	}
}

func Test_day05_7(t *testing.T) {
	// equal
	in := day02.Day02_parse_str("3,9,8,9,10,9,4,9,99,-1,8")
	for i := -100; i < 100; i++ {
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i == 8 {
			if output[0] != 1 {
				t.Errorf("Failed test, got: %v, want: 1", output[0])
			}
		} else if output[0] != 0 {
			t.Errorf("Failed test, got: %v, want: 0, input: %d", output[0], i)
		}
	}
}

func Test_day05_8(t *testing.T) {
	// equal
	in := day02.Day02_parse_str("3,9,7,9,10,9,4,9,99,-1,8")
	for i := -100; i < 100; i++ {
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i < 8 {
			if output[0] != 1 {
				t.Errorf("Failed test, got: %v, want: 1", output[0])
			}
		} else if output[0] != 0 {
			t.Errorf("Failed test, got: %v, want: 0, input: %d", output[0], i)
		}
	}
}

func Test_day05_9(t *testing.T) {
	// equal, immediate
	in := day02.Day02_parse_str("3,3,1108,-1,8,3,4,3,99")
	for i := -100; i < 100; i++ {
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i == 8 {
			if output[0] != 1 {
				t.Errorf("Failed test, got: %v, want: 1", output[0])
			}
		} else if output[0] != 0 {
			t.Errorf("Failed test, got: %v, want: 0, input: %d", output[0], i)
		}
	}
}

func Test_day05_10(t *testing.T) {
	// less than, immediate
	in := day02.Day02_parse_str("3,3,1107,-1,8,3,4,3,99")
	for i := -100; i < 100; i++ {
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i < 8 {
			if output[0] != 1 {
				t.Errorf("Failed test, got: %v, want: 1", output[0])
			}
		} else if output[0] != 0 {
			t.Errorf("Failed test, got: %v, want: 0, input: %d", output[0], i)
		}
	}
}

func Test_day05_11(t *testing.T) {
	// jump
	for i := -100; i < 100; i++ {
		in := day02.Day02_parse_str("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i == 0 {
			if output[0] != 0 {
				t.Errorf("Failed test, got: %v, want: 0", output[0])
			}
		} else if output[0] != 1 {
			t.Errorf("Failed test, got: %v, want: 1, input: %d", output[0], i)
		}
	}
}

func Test_day05_12(t *testing.T) {
	// jump, immediate
	for i := -100; i < 100; i++ {
		in := day02.Day02_parse_str("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i == 0 {
			if output[0] != 0 {
				t.Errorf("Failed test, got: %v, want: 0", output[0])
			}
		} else if output[0] != 1 {
			t.Errorf("Failed test, got: %v, want: 1, input: %d", output[0], i)
		}
	}
}

func Test_day05_13(t *testing.T) {
	// larger example
	for i := -100; i < 100; i++ {
		in := day02.Day02_parse_str("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
		output := make([]int, 0)
		Day05_solve(in, i, &output)
		if i < 8 {
			if output[0] != 999 {
				t.Errorf("Failed test, got: %v, want: 999, input: %d", output[0], i)
			}
		} else if i > 8 {
			if output[0] != 1001 {
				t.Errorf("Failed test, got: %v, want: 1001, input: %d", output[0], i)
			}
		} else if output[0] != 1000 {
			t.Errorf("Failed test, got: %v, want: 1000, input: %d", output[0], i)
		}
	}
}

func Test_day05_14(t *testing.T) {
	// part2
	in := day02.Day02_parse("day05.txt")
	output := make([]int, 0)
	Day05_solve(in, 5, &output)
	if len(output) != 1 || output[0] != 12410607 {
		t.Errorf("Failed test, got: %v, want: %v", output, 12410607)
	}
}