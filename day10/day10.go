package day10

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

type Grid struct {
	vals map[Point]int
	width, height int
}

func (g Grid) Print() {
	point := Point{}
	for i := 0; i < g.height; i++ {
		point.x = 0
		str := ""
		for j := 0; j < g.width; j++ {
			str += fmt.Sprintf("%d", g.vals[point])
			point.x++
		}
		fmt.Printf("%s\n", str)
		point.y++
	}
}

func compare(a float64, b float64) bool {
	const EPSILON float64 = 1e-9
	if a-b > EPSILON || b-a > EPSILON {
		return false
	}
	return true
}

func (g Grid) intersect_points(p0 Point, p1 Point) int {
	// line eq between a and b
	var m, b float64
	var vert bool
	if p1.x - p0.x == 0 {
		vert = true
	} else {
		m = float64(p1.y-p0.y) / float64(p1.x-p0.x)
		b = float64(p0.y) - m*float64(p0.x)
	}

	var dx, dy int
	if p1.x > p0.x {
		dx = 1
	} else if p1.x < p0.x {
		dx = -1
	}
	if p1.y > p0.y {
		dy = 1
	} else if p1.y < p0.y {
		dy = -1
	}

	x := p0.x
	for {
		y := p0.y
		for {
			point:= Point{x:x, y:y}
			if g.vals[point] != 0 && point != p0 && point != p1 {
				if vert {
					if x == p0.x {
						return 1
					}
				}
				if compare(m*float64(x)+b, float64(y)) {
					return 1
				}
			}
			y += dy
			if y == p1.y {
				break
			}
		}
		x += dx
		if x == p1.x {
			break
		}
	}
	return 0
}

func (g Grid) Intersect() Grid {
	output := Grid{vals:make(map[Point]int), width: g.width, height: g.height}
	for p0, v0 := range(g.vals) {
		if v0 == 0 {
			continue
		}
		for p1, v1 := range(g.vals) {
			if v1 == 0  || p0 == p1 {
				continue
			}
			if g.intersect_points(p0, p1) == 0 {
				output.vals[p0] += 1
			}
		}
	}
	return output
}

func (g Grid) Max(point *Point) int {
	max := math.MinInt32
	for k, v := range(g.vals) {
		if v > max {
			if point != nil {
				*point = k
			}
			max = v
		}
	}
	return max
}

type PolarCoord struct {
	point Point
	ang float64
	radius float64
}

func normalize_ang(rads float64) float64 {
	deg := rads * 180 / math.Pi
	for {
		if deg >= 0 {
			return deg
		}
		deg += 360
	}
}

func (g Grid) Vaporize(point Point, count int) Point {
	coords := make([]PolarCoord, 0)
	for k, v := range g.vals {
		if v == 0 || k == point {
			continue
		}
		dx := float64(point.x - k.x)
		dy := float64(point.y - k.y)
		ang := normalize_ang(math.Atan2(-dx, dy))
		radius := math.Sqrt(dx*dx+dy*dy)
		coords = append(coords, PolarCoord{point:k,ang:ang,radius:radius})
	}

	sort.SliceStable(coords, func(i, j int) bool {
		if compare(coords[i].ang, coords[j].ang) {
			return coords[i].radius < coords[j].radius
		}
		return coords[i].ang < coords[j].ang
	})

	result := Point{}
	for i := 0; i < count; i++ {
		cur := coords[0]
		coords = append(coords[:0], coords[1:]...)
		if len(coords) > 1 {
			for loop := 0; loop < len(coords); loop++ {
				if coords[0].ang == cur.ang {
					tmp := coords[0]
					coords = append(coords[:0], coords[1:]...)
					coords = append(coords, tmp)
				} else {
					break
				}
			}
		}
		result = cur.point;
	}
	return result
}

func Day10_parse(fileName string) Grid {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	point := Point{}
	output := Grid{vals:make(map[Point]int)}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		point.x = 0
		bytes := scanner.Bytes()
		for i := 0; i < len(bytes); i++ {
			if bytes[i] == '.' {
				output.vals[point] = 0
			} else if bytes[i] == '#' {
				output.vals[point] = 1
			}
			point.x++
		}
		if output.width == 0 {
			output.width = len(output.vals)
		}
		point.y++
	}
	output.height = len(output.vals)/output.width
	return output
}