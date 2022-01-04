package main

import (
	"strings"
	"sync"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

type ImageEnhancer string

type Image struct {
	pixels            []rune
	width, height     int
	infinitePixel     string
	nextInfinitePixel string
}

func parseInput(data string) (image *Image, enhancer ImageEnhancer) {
	parts := utils.SplitByEmptyNewline(data)

	for _, line := range strings.Split(parts[0], "\n") {
		enhancer += ImageEnhancer(line)
	}

	lines := strings.Split(parts[1], "\n")
	nextInfinitePixel := string(enhancer[0])
	height := len(lines)
	width := len(lines[0])
	pixels := make([]rune, height*width)

	image = &Image{
		pixels:            pixels,
		height:            height,
		width:             width,
		infinitePixel:     ".",
		nextInfinitePixel: nextInfinitePixel,
	}

	for i, line := range lines {
		for j, char := range line {
			val := '0'
			if char == '#' {
				val = '1'
			}
			image.pixels[i*width+j] = val
		}
	}

	return
}

func (image *Image) set(x, y int, val rune) {
	image.pixels[x*image.width+y] = val
}

func (image *Image) get(x, y int) rune {
	if x < 0 || x >= image.width || y < 0 || y >= image.height {
		if image.infinitePixel == "." {
			return '0'
		}
		return '1'
	}

	return image.pixels[x*image.width+y]
}

func (image *Image) isInfinitePixel(i, j int) bool {
	return false
}

func (image *Image) getBinaryForPixel(x, y int) int {
	// TODO: using strings.Builder is better for memory management
	// slightly faster benchmark times
	var binary strings.Builder

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			binary.WriteRune(image.get(i, j))
		}
	}

	val, _ := utils.BinaryToInt(binary.String())

	return val
}

type Worker struct {
	x, y, index int
}

func (ref *Image) enhancePixel(x, y int, stream chan Worker) {
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
func (image *Image) enhance(enhancer ImageEnhancer) (newImage *Image) {
	width := image.width + 2
	height := image.height + 2
	pixels := make([]rune, width*height)
	newImage = &Image{
		pixels:        pixels,
		width:         width,
		height:        height,
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

	wg.Add(image.height * image.width)
	for y := 0; y < image.height; y++ {
		for x := 0; x < image.width; x++ {
			go func(r, c int) {
				image.enhancePixel(r, c, stream)
				wg.Done()
			}(y, x)
		}
	}

	go func() {
		wg.Wait()
		close(stream)
	}()

	for data := range stream {
		x, y := data.x, data.y
		if enhancer[data.index] == '#' {
			newImage.set(x+1, y+1, '1')
		} else {
			newImage.set(x+1, y+1, '0')
		}
	}

	return
}

func (image Image) litCount() int {
	return strings.Count(string(image.pixels), "1")
}

func (image *Image) String() (output string) {
	for y := 0; y < image.height; y++ {
		for x := 0; x < image.width; x++ {
			if image.get(y, x) == '1' {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}

	return
}
