package day08

import (
	"testing"
)

const width = 25
const height = 6

func Test_day08_1(t *testing.T) {
	pixels := pixels_load("day08.txt")
	ans := Day08_part1_solve(width, height, pixels)
	t.Logf("part1: %d", ans)
	if ans != 2413 {
		t.Errorf("Failed test, got: %d, want: 2413", ans)
	}
}

func Test_day08_2(t *testing.T) {
	pixels := pixels_load("day08.txt")
	image := image_create(width, height, pixels)
	layer := image.render()
	layer.png("part2.png", width, height)
}