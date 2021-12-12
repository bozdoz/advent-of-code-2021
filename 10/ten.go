package main

import (
	"fmt"
	"sort"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadFileAsLines

type Bracket struct {
	isOpen                          bool
	pair                            rune
	corruptScore, autocompleteScore int
}

var brackets = map[rune]Bracket{
	'(': {pair: ')', isOpen: true},
	'[': {pair: ']', isOpen: true},
	'{': {pair: '}', isOpen: true},
	'<': {pair: '>', isOpen: true},
	')': {pair: '(', corruptScore: 3, autocompleteScore: 1},
	']': {pair: '[', corruptScore: 57, autocompleteScore: 2},
	'}': {pair: '{', corruptScore: 1197, autocompleteScore: 3},
	'>': {pair: '<', corruptScore: 25137, autocompleteScore: 4},
}

type Stack []rune

func (stack *Stack) pop() (out rune) {
	lastIndex := len(*stack) - 1
	out = (*stack)[lastIndex]

	*stack = (*stack)[:lastIndex]

	return
}

type Line struct {
	isCorrupted, isIncomplete bool
	corruptedBy               rune
	incompleteBrackets        Stack
}

// this constructor is actually doing all the heavy lifting
func newLine(line string) Line {
	stack := Stack{}

	for _, char := range line {
		bracket := brackets[char]

		if bracket.isOpen {
			stack = append(stack, char)
		} else {
			lastChar := stack.pop()

			if lastChar != bracket.pair {
				return Line{
					isCorrupted: true,
					corruptedBy: char,
				}
			}
		}
	}

	return Line{
		isIncomplete:       true,
		incompleteBrackets: stack,
	}
}

// Find the first illegal character in each corrupted line of the navigation subsystem
func PartOne(content []string) (output int, err error) {
	for _, line := range content {
		l := newLine(line)

		if l.isCorrupted {
			output += brackets[l.corruptedBy].corruptScore
		}
	}

	return
}

// gets all pairs of incomplete brackets
func (line *Line) getAutocomplete() (autocomplete string) {
	incomplete := line.incompleteBrackets

	// in reverse order, because it's LIFO
	for i := len(incomplete) - 1; i >= 0; i-- {
		char := incomplete[i]
		autocomplete += string(brackets[char].pair)
	}

	return
}

func getAutocompleteScore(autocomplete string) (score int) {
	for _, char := range autocomplete {
		score *= 5
		score += brackets[char].autocompleteScore
	}

	return
}

// Find the completion string for each incomplete line, score the completion strings,
// sort the scores, and return the middle score (always odd number)
func PartTwo(content []string) (output int, err error) {
	points := []int{}

	for _, line := range content {
		l := newLine(line)

		if l.isIncomplete {
			score := getAutocompleteScore(l.getAutocomplete())
			points = append(points, score)
		}
	}

	sort.Ints(points)

	// always odd, so len / 2 should work
	output = points[len(points)/2]

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
