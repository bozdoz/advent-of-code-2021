package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"bozdoz.com/aoc-2021/utils"
)

func transpose(lines []string) [][]int {
	transposed := make([][]int, len(lines[0]))

	for i, line := range lines {
		for j, char := range line {
			if len(transposed[j]) < 1 {
				transposed[j] = make([]int, len(lines))
			}

			val, err := strconv.Atoi(string(char))

			if err != nil {
				panic(err)
			}

			transposed[j][i] = val
		}
	}

	return transposed
}

// maybe should be utility
func binaryToInt(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func PartOne(lines []string) (int, error) {
	half := float64(len(lines)) / float64(2)

	transposed := transpose(lines)
	gamma := ""
	epsilon := ""

	for _, arr := range transposed {
		sum := float64(utils.Sum(arr...))

		if sum < half {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaint, err := binaryToInt(gamma)

	if err != nil {
		return -1, err
	}

	epsilonint, err := binaryToInt(epsilon)

	if err != nil {
		return -1, err
	}

	return int(gammaint * epsilonint), nil
}

// TODO: how on earth can this be more generic?
// might also be a utility function
func Filter(arr []string, fnc func(val string, i int) bool) (out []string) {
	for i, val := range arr {
		if fnc(val, i) {
			out = append(out, val)
		}
	}

	return
}

func weedOutBinaries(arr []string, check int) (string, error) {
	copy := arr[:]

	for bit := 0; bit < len(arr); bit++ {
		if len(copy) == 1 {
			return copy[0], nil
		}

		transposed := transpose(copy)
		section := transposed[bit]

		// TODO: why is division so difficult?
		half := float64(len(copy)) / float64(2)
		sum := float64(utils.Sum(section...))

		// keep is either 1 or 0
		keep := check
		if sum < half {
			keep = 1 - check
		}

		indices := map[int]bool{}

		for i, val := range section {
			if val == keep {
				indices[i] = true
			}
		}

		copy = Filter(copy, func(val string, i int) bool {
			return indices[i] || false
		})
	}

	return "", errors.New("could not weed out binaries")
}

// TODO: so many if err != nil checks
func PartTwo(lines []string) (int, error) {
	oxygen, err := weedOutBinaries(lines, 1)

	if err != nil {
		return -1, err
	}

	co2, err := weedOutBinaries(lines, 0)

	if err != nil {
		return -1, err
	}

	oxygenint, err := binaryToInt(oxygen)

	if err != nil {
		return -1, err
	}

	co2int, err := binaryToInt(co2)

	if err != nil {
		return -1, err
	}

	return int(oxygenint * co2int), nil
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
