package day03

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func origin() Point {
	return Point{0, 0}
}

type GridVal [2]int
type Grid map[Point]GridVal

func grid_key(x int, y int) Point {
	return Point{x, y}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func update(x int, y int, check int, best *int, best_step *int, grid Grid, steps int) {
	key := grid_key(x, y)
	val := grid[key]
	intersection := val[1] == 0 && val[0] > 0
	val[check] = steps
	if check > 0 && key != origin() && intersection {
		dist := abs(x) + abs(y)
		if dist < *best {
			*best = dist
		}
		step_dist := val[0] + val[1]
		if step_dist < *best_step {
			*best_step = step_dist
		}
	}
	grid[key] = val
}

func parse_lines(input string, grid Grid, wire int, part2 bool) int {
	segments := strings.Split(input, ",")
	prev := Point{0,0}
	best := math.MaxInt32
	best_steps := math.MaxInt32
	steps := 0
	for _, seg := range segments {
		dir := seg[0]
		inc, _ := strconv.Atoi(seg[1:])
		switch dir {
		case 'U':
			for y := prev.y + 1; y <= prev.y + inc; y++ {
				steps++
				update(prev.x, y, wire, &best, &best_steps, grid, steps)
			}
			prev.y += inc
		case 'D':
			for y := prev.y - 1; y >= prev.y - inc; y-- {
				steps++
				update(prev.x, y, wire, &best, &best_steps, grid, steps)
			}
			prev.y -= inc
		case 'L':
			for x := prev.x - 1; x >= prev.x - inc; x-- {
				steps++
				update(x, prev.y, wire, &best, &best_steps, grid, steps)
			}
			prev.x -= inc
		case 'R':
			for x := prev.x + 1; x <= prev.x + inc; x++ {
				steps++
				update(x, prev.y, wire, &best, &best_steps, grid, steps)
			}
			prev.x += inc
		}
	}
	if part2 {
		return best_steps
	}
	return best
}

func Day03_part1_solve(wire_a string, wire_b string) int {
	grid := make(Grid)
	parse_lines(wire_a, grid, 0, false)
	return parse_lines(wire_b, grid, 1, false)
}

func Day03_part2_solve(wire_a string, wire_b string) int {
	grid := make(Grid)
	parse_lines(wire_a, grid, 0, true)
	return parse_lines(wire_b, grid, 1, true)
}