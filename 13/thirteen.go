package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

// discard logs when script is run (overwritten in test file)
func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(ioutil.Discard)
}

// how many dots visible after the first fold
func PartOne(content string) (output int, err error) {
	log.Println("-- Part One --")

	paper := newPaper(content)

	log.Println(paper.Board())

	for _, instruction := range paper.foldInstructions[0:1] {
		paper.fold(instruction)
		log.Println(paper.Board())
	}

	return paper.countDots(), nil
}

// the folded paper reveals eight capital letters
func PartTwo(content string) (output int, err error) {
	paper := newPaper(content)

	for _, instruction := range paper.foldInstructions {
		paper.fold(instruction)
	}

	// today I decided to suppress logs at runtime
	// today the puzzle demands I output some ascii art
	fmt.Println(paper.Board())

	return
}

// TODO: maybe these main scripts should be extracted to some other utility
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
