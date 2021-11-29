package main

import (
	"errors"
	"fmt"
	"os"

	"bozdoz.com/aoc-2021/utils"
)

func PartOne(nums []int) (int, error) {
	for _, first := range nums {
		for _, second := range nums {
			if sum := first + second; sum == 2020 {
				return first * second, nil
			}
		}
	}

	return -1, errors.New("failed to find valid numbers")
}

func PartTwo(nums []int) (int, error) {
	for _, first := range nums {
		for _, second := range nums {
			for _, third := range nums {
				if sum := first + second + third; sum == 2020 {
					return first * second * third, nil
				}
			}
		}
	}

	return -1, errors.New("failed to find valid numbers")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass the txt file as an arg")
		return
	}

	filename := os.Args[1]
	nums := utils.LoadInts(filename)

	answer, err := PartOne(nums)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(nums)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
