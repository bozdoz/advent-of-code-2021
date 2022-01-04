package main

import (
	"fmt"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

func PartOne(content string) (output int, err error) {
	image, enhancer := parseInput(content)

	newImage := image.enhance(enhancer)
	nextImage := newImage.enhance(enhancer)

	return nextImage.litCount(), nil
}

func PartTwo(content string) (output int, err error) {
	image, enhancer := parseInput(content)

	for i := 0; i < 50; i++ {
		image = image.enhance(enhancer)
	}

	return image.litCount(), nil
}

func main() {
	// safe to assume
	filename := "input.txt"

	data := FileLoader(filename)

	answer, err := PartOne(data)

	if err != nil {
		fmt.Println("failed to parse PartOne", err)
		return
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(data)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
