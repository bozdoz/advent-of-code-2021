package main

import (
	"fmt"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

func PartOne(content []string) (output int, err error) {
	caveSys := newCaveSystem(content)

	output = caveSys.findAllPaths()

	return
}

func PartTwo(content []string) (output int, err error) {
	caveSys := newCaveSystem(content)

	caveSys.viewSingleSmallCaveTwice = true

	output = caveSys.findAllPaths()

	return
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
