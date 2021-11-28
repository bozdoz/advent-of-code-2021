package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func Load(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return lines
}

func LoadInts(filename string) []int {
	vals := Load(filename)
	nums := []int{}

	for _, val := range vals {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	return nums
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass the txt file as an arg")
		return
	}

	filename := os.Args[1]
	nums := LoadInts(filename)

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
