package main

import (
	"fmt"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	// log.SetOutput(ioutil.Discard)
}

func PartOne(content []string) (output int, err error) {
	// iterations: 20
	cave := newCave(content)

	// log.Println(cave.String())

	// log.Println(cave.DisplayScores())

	cave.findAllPaths()

	return cave.end.distance, nil
}

func PartTwo(content []string) (output int, err error) {
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
