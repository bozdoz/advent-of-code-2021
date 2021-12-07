package main

import (
	"fmt"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

// map of counter -> count
type State [9]int

func loadState(data string) (state State, err error) {
	// get first line
	lines := strings.Split(data, "\n")
	vals := strings.Split(lines[0], ",")

	for _, val := range vals {
		i, err := strconv.Atoi(val)

		if err != nil {
			return state, err
		}

		state[i]++
	}

	return
}

// incrementDay decrements fish counters
func incrementDay(state *State) {
	populating := state[0]

	// moves all counters to previous index
	for i, val := range state[1:] {
		state[i] = val
	}

	// 0-day counters are both moved to 8 and copied to 6
	state[8] = populating
	state[6] += populating
}

func fastforward(state *State, days int) {
	for i := 0; i < days; i++ {
		incrementDay(state)
	}
}

func PartOne(content string) (output int, err error) {
	state, err := loadState(content)

	fastforward(&state, 80)

	return utils.Sum(state[:]...), err
}

func PartTwo(content string) (output int, err error) {
	state, err := loadState(content)

	fastforward(&state, 256)

	return utils.Sum(state[:]...), err
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
