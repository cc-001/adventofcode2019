package day06

import (
	"testing"
)

func Test_day06_1(t *testing.T) {
	tree := Day06_parse("day06_test.txt")
	checksum := tree.checksum()
	if checksum != 42 {
		t.Errorf("Failed test, got: %d, want: 42", checksum)
	}
}

func Test_day06_part1(t *testing.T) {
	tree := Day06_parse("day06.txt")
	checksum := tree.checksum()
	if checksum != 322508 {
		t.Logf("Failed test, got: %d, want: 322508", checksum)
	}
}

func Test_day06_part2_0(t *testing.T) {
	tree := Day06_parse("day06_test2.txt")
	path_count := tree.path_count("YOU", "SAN")
	if path_count != 4 {
		t.Logf("Failed test, got: %d, want: 4", path_count)
	}
}

func Test_day06_part2_1(t *testing.T) {
	tree := Day06_parse("day06.txt")
	path_count := tree.path_count("YOU", "SAN")
	if path_count != 496 {
		t.Logf("Failed test, got: %d, want: 496", path_count)
	}
}