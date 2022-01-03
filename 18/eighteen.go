package main

import (
	"fmt"
	"io/ioutil"

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

func PartOne(content []string) (output int, err error) {
	curPair := parsePairs(content[0])

	for _, data := range content[1:] {
		curPair = curPair.add(parsePairs(data))
	}

	return curPair.getMagnitude(), nil
}

func PartTwo(content []string) (output int, err error) {
	pairs := []*Pair{}

	max := 0

	for _, data := range content {
		pairs = append(pairs, parsePairs(data))
	}

	for _, pair := range pairs {
		for _, next := range pairs {
			if pair == next {
				continue
			}
			// lazy copy
			pairCopy := parsePairs(pair.String())
			// lazy copy
			nextCopy := parsePairs(next.String())

			mag := pairCopy.add(nextCopy).getMagnitude()

			if mag > max {
				max = mag
			}
		}
	}

	return max, nil
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
