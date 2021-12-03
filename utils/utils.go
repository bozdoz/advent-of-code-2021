package utils

import (
	"bufio"
	"os"
	"strconv"
)

func LoadFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// I don't know what this line does
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func LoadInts(filename string) []int {
	vals := LoadFile(filename)
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

func Sum(nums ...int) (s int) {
	for _, val := range nums {
		s += val
	}

	return
}