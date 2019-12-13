package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type vec [3]int

func vec_sub(a vec, b vec) vec {
	return vec{a[0]-b[0],a[1]-b[1],a[2]-b[2]}
}

func vec_add(a vec, b vec) vec {
	return vec{a[0]+b[0],a[1]+b[1],a[2]+b[2]}
}

func vec_sum_abs(a vec) int {
	return abs(a[0]) + abs(a[1]) + abs(a[2])
}

type particle struct {
	pos vec
	vel vec
}

type history struct {
	recs []particle
	cycle [3]int
}

func make_history(particles *[]particle) history {
	h := history{recs:make([]particle, len(*particles))}
	for i := 0; i < len(*particles); i++ {
		h.recs[i] = (*particles)[i]
	}
	return h
}

func (h *history) check_cycle(particles *[]particle, step int, coord int) bool {
	if h.cycle[coord] > 0 {
		return true
	}
	for i := 0; i < len(*particles); i++ {
		if (*particles)[i].pos[coord] != h.recs[i].pos[coord] || (*particles)[i].vel[coord] != h.recs[i].vel[coord] {
			return false
		}
	}
	h.cycle[coord] = step
	return true
}

func (h *history) check(particles *[]particle, step int) {
	for i := 0; i < len(h.cycle); i++ {
		h.check_cycle(particles, step, i)
	}
}

func (h history) complete() bool {
	for i := 0; i < len(h.cycle); i++ {
		if h.cycle[i] <= 0 {
			return false
		}
	}
	return true;
}

func (h history) print() {
	for i := 0; i < len(h.cycle); i++ {
		fmt.Printf("%d\n", h.cycle[i])
	}
}

func (h history) gather() []int {
	return h.cycle[:]
}

func gravity(a int) int {
	if a < 0 {
		return 1
	} else if a == 0 {
		return 0
	}
	return -1
}

func apply_gravity(particles *[]particle) {
	num := len(*particles)
	for i := 0; i < num; i++ {
		for j := i+1; j < num; j++ {
			d := vec_sub((*particles)[i].pos, (*particles)[j].pos)
			for k := 0; k < len(d); k++ {
				g := gravity(d[k])
				(*particles)[i].vel[k] += g
				(*particles)[j].vel[k] -= g
			}
		}
	}
}

func apply_velocity(particles *[]particle) {
	num := len(*particles)
	for i := 0; i < num; i++ {
		(*particles)[i].pos = vec_add((*particles)[i].pos, (*particles)[i].vel)
	}
}

func total_energy(particles *[]particle) int {
	num := len(*particles)
	e := 0
	for i := 0; i < num; i++ {
		p := vec_sum_abs((*particles)[i].pos)
		k := vec_sum_abs((*particles)[i].vel)
		e += p*k
	}
	return e
}

func step(particles *[]particle) {
	apply_gravity(particles);
	apply_velocity(particles);
}

func print(particles *[]particle) {
	for i := 0; i < len(*particles); i++ {
		p := (*particles)[i]
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n",
			p.pos[0], p.pos[1], p.pos[2],
			p.vel[0], p.vel[1], p.vel[2])
	}
}

func day12_parse(file_name string) []particle {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]particle, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := particle{}
		fmt.Sscanf(scanner.Text(),"<x=%d, y=%d, z=%d>", &p.pos[0], &p.pos[1], &p.pos[2])
		input = append(input, p)
	}
	return input
}

func Day12_solve_part1(file string, steps int) int {
	particles := day12_parse(file)
	ret := 0
	for i := 0; i < steps; i++ {
		ret = total_energy(&particles)
		step(&particles)
	}
	return ret
}

func gcd(a int, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func lcm(inputs []int) int {
	num := len(inputs)
	if num < 2 {
		log.Fatal()
	}

	res := inputs[0] * inputs[1] / gcd(inputs[0], inputs[1])
	if num >= 2 {
		tmp := make([]int, 2)
		for i := 2; i < len(inputs); i++ {
			tmp[0] = res
			tmp[1] = inputs[i]
			res = lcm(tmp)
		}
	}
	return res
}

/*
They're not sending their best problems folks.

It's obvious there are sub-cycles in it otherwise it is inane.  I decided in a couple mins there were smaller repeating
patterns otherwise it is intractable.  This is part of the issue I had with some of the problems in 2018.

It doesn't use any techniques of optimizing an actual N-body problem with large numbers of bodies like
distance approximations, gridding, etc...  This is more like a "gears" application of LCM.
*/
func Day12_solve_part2(file string) int {
	particles := day12_parse(file)
	history := make_history(&particles)
	count := 0
	for {
		count++
		step(&particles)
		history.check(&particles, count)
		if history.complete() {
			cycles := history.gather()
			log.Printf("cycles: %v", cycles)
			return lcm(cycles)
		}
	}
}