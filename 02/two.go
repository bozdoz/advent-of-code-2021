package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

func withVals(vals []string, fnc func(direction string, val int)) error {
	for _, val := range vals {
		v := strings.Fields(val)
		num, err := strconv.Atoi(v[1])

		if err != nil {
			return err
		}

		fnc(v[0], num)
	}

	return nil
}

func PartOne(vals []string) (int, error) {
	horizontal := 0
	depth := 0

	err := withVals(vals, func(direction string, val int) {
		switch direction {
		case "forward":
			horizontal += val
		case "backward":
			horizontal -= val
		case "up":
			depth -= val
		case "down":
			depth += val
		}
	})

	if err != nil {
		return -1, err
	}

	return horizontal * depth, nil
}

func PartTwo(vals []string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	err := withVals(vals, func(direction string, val int) {
		switch direction {
		case "forward":
			horizontal += val
			depth += val * aim
		case "backward":
			horizontal -= val
			depth -= val * aim
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	})

	if err != nil {
		return -1, err
	}

	return horizontal * depth, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass the txt file as an arg")
		return
	}

	filename := os.Args[1]
	vals := utils.LoadFileAsLines(filename)

	answer, err := PartOne(vals)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(vals)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
