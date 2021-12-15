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
	// TODO: find out how to get a logger to
	// always log Llongfile: followed by "\n"
	log.SetFlags(log.Llongfile)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content string) (output int, err error) {
	log.Println("-- PART ONE --")

	polymer := newPolymer(content)

	elements := polymer.pairInsertion(10)

	log.Println("\n", polymer.template, newElements(polymer.template))

	// first step
	log.Println("\n", "NCNBCHB\n", polymer.pairInsertion(1))

	// 10th step
	log.Println("\n", "Step 10", "\n", elements)

	min, max := elements.getMinMax()

	return max - min, nil
}

func PartTwo(content string) (output int, err error) {
	log.Println("-- PART TWO --")

	polymer := newPolymer(content)

	elements := polymer.pairInsertion(40)

	min, max := elements.getMinMax()

	return max - min, nil
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
