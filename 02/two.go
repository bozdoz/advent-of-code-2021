package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

func WithMatchedStrings(vals []string, forEach func(parts []string)) {
	re := regexp.MustCompile(`(\d+)\-(\d+)\s(\w):\s(\w+)$`)

	for _, val := range vals {
		parts := re.FindAllStringSubmatch(val, -1)[0]

		forEach(parts)
	}
}

type Policy struct {
	min, max       int
	char, password string
}

func goodPassword(pass Policy) bool {
	count := strings.Count(pass.password, pass.char)

	return count >= pass.min && count <= pass.max
}

func PartOne(vals []string) (int, error) {
	count := 0

	WithMatchedStrings(vals, func(parts []string) {
		min, _ := strconv.Atoi(parts[1])
		max, _ := strconv.Atoi(parts[2])

		policy := Policy{
			min:      min,
			max:      max,
			char:     parts[3],
			password: parts[4],
		}

		if goodPassword(policy) {
			count += 1
		}
	})

	return count, nil
}

type Policy2 struct {
	first, second  int
	char, password string
}

func goodPassword2(pass Policy2) bool {
	first := string(pass.password[pass.first-1]) == pass.char
	second := string(pass.password[pass.second-1]) == pass.char

	if first && second {
		return false
	}

	return first || second
}

func PartTwo(vals []string) (int, error) {
	count := 0

	WithMatchedStrings(vals, func(parts []string) {
		first, _ := strconv.Atoi(parts[1])
		second, _ := strconv.Atoi(parts[2])

		policy := Policy2{
			first:    first,
			second:   second,
			char:     parts[3],
			password: parts[4],
		}

		if goodPassword2(policy) {
			count += 1
		}
	})

	return count, nil
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