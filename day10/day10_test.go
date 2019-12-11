package day10

import (
	"testing"
)

func Test_day10_1(t *testing.T) {
	in := Day10_parse("day10_0.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 8
	ans_coord := Point{x:3,y:4}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func Test_day10_2(t *testing.T) {
	in := Day10_parse("day10_1.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 33
	ans_coord := Point{x:5,y:8}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func Test_day10_3(t *testing.T) {
	in := Day10_parse("day10_2.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 35
	ans_coord := Point{x:1,y:2}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func Test_day10_4(t *testing.T) {
	in := Day10_parse("day10_3.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 41
	ans_coord := Point{x:6,y:3}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func Test_day10_5(t *testing.T) {
	in := Day10_parse("day10_4.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 210
	ans_coord := Point{x:11,y:13}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func Test_day10_part1(t *testing.T) {
	in := Day10_parse("day10_input.txt")
	out := in.Intersect()
	coord := Point{}
	max := out.Max(&coord)
	ans := 256
	ans_coord := Point{x:29,y:28}
	if max != ans || coord != ans_coord {
		t.Errorf("Failed test, got: %v - %v, want: %v - %v", max, coord, ans, ans_coord)
	}
}

func test_vaporize(t *testing.T, g Grid, c int, ans Point) {
	tmp := g.Vaporize(Point{x:11,y:13}, c)
	if tmp != ans {
		t.Errorf("Failed test %d, got: %v, want: %v", c, tmp, ans)
	}
}
func Test_day10_6(t *testing.T) {
	in := Day10_parse("day10_4.txt")
	test_vaporize(t, in, 1, Point{11,12})
	test_vaporize(t, in, 2, Point{12,1})
	test_vaporize(t, in, 3, Point{12,2})
	test_vaporize(t, in, 10, Point{12,8})
	test_vaporize(t, in, 20, Point{16,0})
	test_vaporize(t, in, 50, Point{16,9})
	test_vaporize(t, in, 100, Point{10,16})
	test_vaporize(t, in, 199, Point{9,6})
	test_vaporize(t, in, 200, Point{8,2})
	test_vaporize(t, in, 201, Point{10,9})
	test_vaporize(t, in, 299, Point{11,1})
}

func Test_day10_7(t *testing.T) {
	in := Day10_parse("day10_input.txt")
	tmp := in.Vaporize(Point{x:29,y:28}, 200)
	ans := tmp.x * 100 + tmp.y
	if ans != 1707 {
		t.Errorf("Failed test, got: %v, want: 1707", ans)
	}
}