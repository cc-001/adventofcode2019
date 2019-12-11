package day11

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

const up = 0
const right = 1
const down = 2
const left = 3

type point struct {
	x, y int
}

type robot struct {
	memory []int
	pc int
	rb int
	facing int
	pos point
}

type panel struct {
	color, count int
}

type ship struct {
	initial_panel_color int
	panels map[point]panel
}

func (s ship) png(file string) {
	minx, miny := math.MaxInt32, math.MaxInt32
	maxx, maxy := math.MinInt32, math.MinInt32
	for k, _ := range s.panels {
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

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			c := 0
			if v, ok := s.panels[point{x:x, y:y}]; ok {
				c = v.color
			}
			if c == 0 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	scale := uint(4)
	img_scaled := resize.Resize(uint(width)*scale, uint(height)*scale, img, resize.NearestNeighbor)

	if err := png.Encode(f, img_scaled); err != nil {
		log.Fatal(err)
	}
}

func (r *robot) tick(s *ship) bool {
	input := make([]int, 1)
	if val, ok := s.panels[r.pos]; ok {
		input[0] = val.color
	} else if len(s.panels) == 0 {
		input[0] = s.initial_panel_color
	} else {
		input[0] = 0
	}

	output := make([]int, 0)
	r.memory = day05.Day05_solve(r.memory, input, &output, &r.pc, &r.rb,false)
	if r.pc < 0 {
		return false
	}
	r.memory = day05.Day05_solve(r.memory, input, &output, &r.pc, &r.rb,false)
	if r.pc < 0 {
		return false
	}

	if val, ok := s.panels[r.pos]; ok {
		val.color = output[0]
		val.count++
		s.panels[r.pos] = val
	} else {
		s.panels[r.pos] = panel{color:output[0],count:1}
	}

	if output[1] == 1 {
		r.facing = (r.facing + 1) % 4
	} else {
		if r.facing > 0 {
			r.facing--
		} else {
			r.facing = left
		}
	}

	switch r.facing {
	case up:
		r.pos.y -= 1
	case right:
		r.pos.x += 1
	case down:
		r.pos.y += 1
	case left:
		r.pos.x -= 1
	}
	return true
}

func (r *robot) load(file string) {
	r.memory = day02.Day02_parse(file)
}

func Day11_solve_part1(file string) int {
	s := ship{panels:make(map[point]panel)}
	r := robot{}
	r.load(file)
	for {
		if r.tick(&s) == false {
			return len(s.panels)
		}
	}
}

func Day11_solve_part2(file string, png string) {
	s := ship{panels:make(map[point]panel),initial_panel_color:1}
	r := robot{}
	r.load(file)
	for {
		if r.tick(&s) == false {
			break
		}
	}
	s.png(png)
}