package day13

import (
	"advent2019/day02"
	"advent2019/day05"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

type point struct {
	x, y int
}

func point_add(a point, b point) point {
	return point{a.x+b.x,a.y+b.y}
}

func point_sub(a point, b point) point {
	return point{a.x-b.x,a.y-b.y}
}

const tile_wall = 1
const tile_block = 2
const tile_paddle = 3
const tile_ball = 4

var invalid point = point{-1,-1}

type game struct {
	screen map[point]int
	prev_ball point
	prev_input int
	score int
	inputs int
	intersect point
	images []*image.Image
}

func (g *game) consume_outputs(output []int) {
	score_coord := point{-1,0}
	for i := 0; i < len(output); i += 3 {
		point := point{output[i],output[i+1]}
		if point == score_coord {
			g.score = output[i+2]
		} else  {
			g.screen[point] = output[i+2]
		}
	}
}

func (g game) count_blocks() int {
	count := 0
	for _, v := range g.screen {
		if v == tile_block {
			count++
		}
	}
	return count
}

func (g game) get_point(tile int) point {
	for k, v := range g.screen {
		if v == tile {
			return k
		}
	}
	return invalid
}

func (g game) render(file string) {
	minx, miny := math.MaxInt32, math.MaxInt32
	maxx, maxy := math.MinInt32, math.MinInt32
	for k, _ := range g.screen {
		if k.x < minx {
			minx = k.x
		}
		if k.y < miny {
			miny = k.y
		}
		if k.x > maxx {
			maxx = k.x
		}
		if k.y > maxy {
			maxy = k.y
		}
	}

	width := (maxx-minx)+1
	height := (maxy-miny)+1
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	colors := [5]color.NRGBA{{0,0,0,255},
		{113,113,113,255},
		{180,122,48,255},
		{191,71,72,255},
		{191,71,72,255}}

	blocks := [6]color.NRGBA{{191,71,72,255},
		{194,105,58,255},
		{172,118,44,255},
		{149,158,42,255},
		{76,157,77,255},
		{60,64,168,255}}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if v, ok := g.screen[point{x:x, y:y}]; ok {
				if v == 2 {
					img.Set(x, y, blocks[y%6])
				} else {
					img.Set(x, y, colors[v])
				}
			}
		}
	}

	if g.intersect != invalid {
		img.Set(g.intersect.x, g.intersect.y, color.White)
	}

	scale := uint(8)
	img_scaled := resize.Resize(uint(width)*scale, uint(height)*scale, img, resize.NearestNeighbor)
	g.images = append(g.images, &img_scaled)

	if err := png.Encode(f, img_scaled); err != nil {
		log.Fatal(err)
	}
}

func Day13_solve_part1() int {
	input := day02.Day02_parse("day13.txt")
	output := make([]int, 0)
	day05.Day05_solve_adapter(input, 0, &output)
	game := game{screen:make(map[point]int)}
	game.consume_outputs(output)
	return game.count_blocks()
}

var g_game game = game{screen:make(map[point]int),intersect:invalid,prev_ball:invalid}

func game_input(cur point) int {
	paddle := g_game.get_point(tile_paddle)
	if paddle == invalid {
		return 0
	}

	val0 := 0
	if cur.x > paddle.x {
		val0 = 1
	}

	val1 := 0
	if cur.x < paddle.x {
		val1 = 1
	}
	return val0 - val1
}

func program_input() int {
	cur := g_game.get_point(tile_ball)
	g_game.prev_input = game_input(cur)
	g_game.prev_ball = cur
	//g_game.render(fmt.Sprintf("%d.png", g_game.inputs))
	g_game.inputs++
	return g_game.prev_input
}

func Day13_solve_part2() int {
	program := day02.Day02_parse("day13.txt")
	program[0] = 2

	pc := 0
	rb := 0
	output := make([]int, 0, 3)

	zero := false
	for {
		program = day05.Day05_solve(program, nil, &output, &pc, &rb, false, program_input)
		if pc < 0 {
			log.Fatal()
		}

		if len(output) == 3 {
			g_game.consume_outputs(output)
			if g_game.inputs > 0 && g_game.count_blocks() == 0 {
				if zero {
					return g_game.score
				} else {
					zero = true
				}
			}
			output = make([]int, 0, 3)
		}
	}
}