package day08

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

type layer struct {
	pixels []int
}

func (l layer) png(file string, width int, height int) {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y * width + x
			if l.pixels[idx] == 0 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

type eimage struct {
	width int
	height int
	layers []layer
}

func image_create(width int, height int, input []int) eimage {
	layer_pixels := width*height
	num_layers := len(input)/layer_pixels
	image := eimage{
		width: width,
		height: height,
		layers: make([]layer, 0, num_layers),
	}
	for i := 0; i < num_layers; i++ {
		layer := layer{pixels: make([]int, layer_pixels)}
		start := i*layer_pixels
		end := start+layer_pixels
		copy(layer.pixels, input[start:end])
		image.layers = append(image.layers, layer)
	}
	return image
}

func (i eimage) render() layer {
	layer_pixels := i.width*i.height
	output := layer{make([]int, layer_pixels)}
	written := map[int]bool{}
	for _, layer := range i.layers {
		for i := 0; i < layer_pixels; i++ {
			if _, has_val := written[i]; !has_val && layer.pixels[i] != 2 {
				written[i] = true
				output.pixels[i] = layer.pixels[i]
			}
		}
	}
	return output
}

func pixels_load(file string) []int {
	buf, _ := ioutil.ReadFile(file)
	str := string(buf)
	pixels := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {
		pixel, _ := strconv.Atoi(str[i:i+1])
		pixels = append(pixels, pixel)
	}
	return pixels
}

func Day08_part1_solve(width int, height int, input []int) int {
	image := image_create(width, height, input)
	best := math.MaxInt32
	var ans int
	for _, i := range image.layers {
		counts := [3]int{0, 0, 0}
		for _, j := range i.pixels {
			counts[j]++
		}
		if counts[0] < best {
			best = counts[0]
			ans = counts[1]*counts[2]
		}
	}
	return ans
}