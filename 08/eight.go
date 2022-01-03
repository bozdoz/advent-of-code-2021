package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

// how many times do digits 1, 4, 7, or 8 appear?
func PartOne(lines []string) (output int, err error) {
	var counts [10]int

	for _, line := range lines {
		output_values := strings.Fields(strings.Split(line, " | ")[1])

		for _, val := range output_values {
			counts[len(val)]++
		}
	}

	return counts[2] + counts[4] + counts[3] + counts[7], nil
}

const (
	unknown = -1
)

type Pattern string

type Entry struct {
	patterns map[Pattern]int
	digits   [4]Pattern
}

func newPattern(p string) Pattern {
	p = utils.SortString(p)

	return Pattern(p)
}

func newEntry(line string) (entry Entry) {
	parts := strings.Split(line, " | ")

	// why do I always have to initialize this "nil map"
	entry.patterns = map[Pattern]int{}

	for _, pattern := range strings.Fields(parts[0]) {
		entry.patterns[newPattern(pattern)] = unknown
	}

	for i, code := range strings.Fields(parts[1]) {
		entry.digits[i] = newPattern(code)
	}

	return
}

// all parts of b are inside of a
func (a *Pattern) contains(b Pattern) bool {
outer:
	for _, x := range b {
		for _, y := range *a {
			if x == y {
				continue outer
			}
		}
		return false
	}

	return true
}

// if any candidates contains all of pattern, then assign it the value
func (entry *Entry) deriveThroughIntersection(candidates []Pattern, pattern Pattern, value int) Pattern {
	for _, candidate := range candidates {
		if candidate.contains(pattern) {
			entry.patterns[candidate] = value

			return candidate
		}
	}
	panic(fmt.Sprintf("pattern not found: %s", pattern))
}

// created purely to make the remove method
type PatternArr []Pattern

func (arr PatternArr) remove(pattern Pattern) (out PatternArr) {
	for _, val := range arr {
		if val != pattern {
			out = append(out, val)
		}
	}

	return
}

func (entry *Entry) derivePatterns() {
	// save one and four because they are used to get other values
	var one, four Pattern
	var lenfive, lensix PatternArr

	// givens: 1, 4, 7, 8
	for pattern := range entry.patterns {
		val := unknown
		switch len(pattern) {
		case 2:
			val = 1
			one = pattern
		case 3:
			val = 7
		case 4:
			val = 4
			four = pattern
		case 5:
			// 2, 3, 5
			lenfive = append(lenfive, pattern)
		case 6:
			// 0, 6, 9
			lensix = append(lensix, pattern)
		case 7:
			val = 8
		}

		entry.patterns[pattern] = val
	}

	// 2, 3 & 5 are all of length 5, but only 3 includes all segments from 1
	matched := entry.deriveThroughIntersection(lenfive, one, 3)
	lenfive = lenfive.remove(matched)

	// 9 contains 4
	matched = entry.deriveThroughIntersection(lensix, four, 9)
	lensix = lensix.remove(matched)

	// 0 contains 1
	matched = entry.deriveThroughIntersection(lensix, one, 0)

	// 6 remains
	six := lensix.remove(matched)[0]
	entry.patterns[six] = 6

	// 2 & 5 are left: 6 contains 5
	for _, candidate := range lenfive {
		// reversed from above
		if six.contains(candidate) {
			entry.patterns[candidate] = 5
		} else {
			// only 2 remains
			entry.patterns[candidate] = 2
		}
	}
}

// we concatenate each entry
func (entry *Entry) concatDigits() (int, error) {
	sum := ""
	for _, pattern := range entry.digits {
		sum += strconv.Itoa(entry.patterns[pattern])
	}

	return strconv.Atoi(sum)
}

func PartTwo(lines []string) (output int, err error) {
	for _, line := range lines {
		entry := newEntry(line)
		entry.derivePatterns()

		sum, err := entry.concatDigits()

		if err != nil {
			return output, err
		}

		output += sum
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
