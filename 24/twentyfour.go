package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content []string) (output interface{}, err error) {
	program := parseInput(content)

	num := program.solveLargest()

	return num, nil
}

func PartTwo(content []string) (output interface{}, err error) {
	program := parseInput(content)

	num := program.solveSmallest()

	return num, nil
}

func main() {
	// safe to assume
	filename := "input.txt"

	data := FileLoader(filename)

	partFlag := flag.Int("part", -1, "pass a flag for -part")

	flag.Parse()

	if *partFlag < 2 {
		start := time.Now()
		answer, err := PartOne(data)

		if err != nil {
			fmt.Println("failed to parse PartOne", err)
			return
		}

		fmt.Printf("Part One: %d \n", answer)
		fmt.Printf("Time: %s \n", time.Since(start))
	}

	if *partFlag != 1 {
		start := time.Now()
		answer2, err := PartTwo(data)

		if err != nil {
			fmt.Println("failed to parse PartTwo", err)
			return
		}

		fmt.Printf("Part Two: %d \n", answer2)
		fmt.Printf("Time: %s \n", time.Since(start))
	}
}
