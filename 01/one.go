package main

import (
	"fmt"
	"os"

	"bozdoz.com/aoc-2021/utils"
)

func PartOne(nums []int) (int, error) {
	last := nums[0]
	count := 0

	for _, val := range nums {
		if val > last {
			count += 1
		}
		last = val
	}

	return count, nil
}

func sum(nums []int) int {
	s := 0

	for _, val := range nums {
		s += val
	}

	return s
}

func PartTwo(nums []int) (int, error) {
	measurement := 3
	wins := [][]int{}
	max := len(nums) - measurement + 1

	for i := range nums[:max] {
		wins = append(wins, nums[i:i+measurement])
	}

	last := sum(wins[0])
	count := 0
	for _, arr := range wins[1:] {
		val := sum(arr)
		if val > last {
			count += 1
		}
		last = val
	}

	return count, nil
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
