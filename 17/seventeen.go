package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content string) (output int, err error) {
	target := parseTarget(content)

	log.Println(target)

	maxHeight := 0

	target.practice(func(probe *Probe, success bool) {
		if success && probe.maxHeight > maxHeight {
			maxHeight = probe.maxHeight
		}
	})

	return maxHeight, nil
}

func PartTwo(content string) (output int, err error) {
	target := parseTarget(content)

	log.Println(target)

	hitCount := 0

	target.practice(func(_ *Probe, success bool) {
		if success {
			hitCount++
		}
	})

	return hitCount, nil
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
