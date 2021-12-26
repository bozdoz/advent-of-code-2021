package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content []string) (output int, err error) {
	// clamped to -50,50
	clampData := true
	grid := &Cubes{}
	grid.parseInstructions(content, clampData)

	return grid.count(), nil
}

func PartTwo(content []string) (output int, err error) {
	clampData := false
	grid := &Cubes{}
	grid.parseInstructions(content, clampData)

	return grid.count(), nil
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

	start := time.Now()

	answer2, err := PartTwo(data)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
	diff := time.Since(start)
	fmt.Printf("Time: %s\n", diff)
}
