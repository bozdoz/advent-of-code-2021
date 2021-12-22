package main

import (
	"strings"
	"sync"

	"bozdoz.com/aoc-2021/utils"
)

type ImageEnhancer string

// pixel grid
type Image struct {
	pixels                    map[int]map[int]bool
	width, height, minx, miny int
	infinitePixel             string
	nextInfinitePixel         string
}

func parseInput(data string) (image *Image, enhancer ImageEnhancer) {
	parts := utils.SplitByEmptyNewline(data)

	for _, line := range strings.Split(parts[0], "\n") {
		enhancer += ImageEnhancer(line)
	}

	lines := strings.Split(parts[1], "\n")
	nextInfinitePixel := string(enhancer[0])

	image = &Image{
		pixels:            map[int]map[int]bool{},
		height:            len(lines),
		width:             len(lines[0]),
		minx:              0,
		miny:              0,
		infinitePixel:     ".",
		nextInfinitePixel: nextInfinitePixel,
	}

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				image.set(i, j)
			}
		}
	}

	return
}

func (image *Image) set(x, y int) {
	if (*image).pixels[x] == nil {
		(*image).pixels[x] = map[int]bool{}
	}
	(*image).pixels[x][y] = true
}

// first time using recover and defer?
func (image *Image) get(x, y int) bool {
	defer recover()

	return (*image).pixels[x][y]
}

func (image *Image) isInfinitePixel(i, j int) bool {
	return i < image.minx ||
		i > (image.width+image.minx) ||
		j < image.miny ||
		j > (image.height+image.miny)
}

func (image *Image) getBinaryForPixel(x, y int) int {
	binary := ""

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			lit := image.get(i, j)
			if lit || (image.infinitePixel != "." && image.isInfinitePixel(i, j)) {
				binary += "1"
			} else {
				binary += "0"
			}
		}
	}

	val, _ := utils.BinaryToInt(binary)

	return val
}

type Worker struct {
	x, y, index int
}

func (ref Image) enhancePixel(x, y int, stream chan Worker) {
	// x,y is lit
	// enhance all pixels around it
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			stream <- Worker{
				x:     i,
				y:     j,
				index: ref.getBinaryForPixel(i, j),
			}
		}
	}
}

// need to update pixels all around each lit pixel
func (image Image) enhance(enhancer ImageEnhancer) (newImage *Image) {
	newImage = &Image{
		pixels:        map[int]map[int]bool{},
		width:         image.width + 1,
		height:        image.height + 1,
		minx:          image.minx - 1,
		miny:          image.miny - 1,
		infinitePixel: image.nextInfinitePixel,
	}

	if newImage.infinitePixel == "." {
		// all dots is 0
		newImage.nextInfinitePixel = string(enhancer[0])
	} else {
		// all #'s is 511
		newImage.nextInfinitePixel = string(enhancer[511])
	}
	var wg sync.WaitGroup

	stream := make(chan Worker)

	for r, row := range image.pixels {
		for c := range row {
			wg.Add(1)
			go func(r, c int) {
				defer wg.Done()
				image.enhancePixel(r, c, stream)
			}(r, c)
		}
	}

	go func() {
		wg.Wait()
		close(stream)
	}()

	for data := range stream {
		x, y, index := data.x, data.y, data.index
		if enhancer[index] == '#' {
			newImage.set(x, y)
		}
	}

	return
}

func (image Image) litCount() (sum int) {
	for _, row := range image.pixels {
		sum += len(row)
	}

	return
}

func (image *Image) String() (output string) {
	for y := image.miny; y < image.height; y++ {
		for x := image.minx; x < image.width; x++ {
			if image.get(y, x) {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}

	return
}
