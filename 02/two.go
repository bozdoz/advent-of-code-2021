package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	
	"bozdoz.com/aoc-2021/utils"
)

func PartOne(vals []string) (int, error) {
	horizontal := 0
	depth := 0

	for _, val := range vals {
		v := strings.Fields(val)
		num, err := strconv.Atoi(v[1])

		if err != nil {
			return -1, err
		}

		switch v[0] {
			case "forward":
				horizontal += num
			case "backward":
				horizontal -= num
			case "up":
				depth -= num
			case "down":
				depth += num
		}
	}

	return horizontal * depth, nil
}

func PartTwo(vals []string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, val := range vals {
		v := strings.Fields(val)
		num, err := strconv.Atoi(v[1])

		if err != nil {
			return -1, err
		}

		switch v[0] {
			case "forward":
				horizontal += num
				depth += num * aim
			case "backward":
				horizontal -= num
				depth -= num * aim
			case "up":
				aim -= num
			case "down":
				aim += num
		}
	}

	return horizontal * depth, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass the txt file as an arg")
		return
	}

	filename := os.Args[1]
	vals := utils.LoadFile(filename)

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
