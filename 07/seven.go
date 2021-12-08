package main

import (
	"fmt"
	"sort"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadCSVInt

// maybe utility (math.Abs is float64)
func AbsDiff(a int, b int) int {
	diff := a - b

	if diff < 0 {
		diff *= -1
	}

	return diff
}

func calcLinearFuel(positions []int, ref int) (fuel int) {
	for _, val := range positions {
		fuel += AbsDiff(val, ref)
	}

	return
}

// maybe utility function
func sortedUnique(data []int) (unique []int) {
	set := map[int]bool{}

	for _, val := range data {
		set[val] = true
	}

	for key := range set {
		unique = append(unique, key)
	}

	sort.Ints(unique)

	return
}

func PartOne(content []int) (minfuel int, err error) {
	unique := sortedUnique(content)

	minfuel = calcLinearFuel(content, unique[0])

	// checks positions in dataset only (naive)
	for _, val := range unique[1:] {
		newfuel := calcLinearFuel(content, val)

		if newfuel < minfuel {
			minfuel = newfuel
		} else {
			break
		}
	}

	return
}

// adds all iterations of numbers by decrementing by 1:
// example: 4 = 4 + 3 + 2 + 1
// update: Should have been dist * (dist + 1) / 2
func getCumulativeFuel(dist int) (fuel int) {
	for ; dist > 0; dist-- {
		fuel += dist
	}

	return
}

// each change of 1 step in horizontal position costs 1 more unit of fuel than the last
func calcExponentialFuel(positions []int, ref int) (fuel int) {
	for _, val := range positions {
		fuel += getCumulativeFuel(AbsDiff(val, ref))
	}

	return
}

func PartTwo(content []int) (minfuel int, err error) {
	unique := sortedUnique(content)

	min := unique[0]
	max := unique[len(unique)-1]

	minfuel = calcExponentialFuel(content, min)

	// checks positions not in dataset
	for val := min + 1; val < max; val++ {
		newfuel := calcExponentialFuel(content, val)

		if newfuel < minfuel {
			minfuel = newfuel
		} else {
			break
		}
	}

	return
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
