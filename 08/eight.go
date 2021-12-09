package main

import (
	"fmt"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadFile

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

type Pattern struct {
	value, length int
}

type Entry struct {
	patterns map[string]Pattern
	digits   [4]string
}

func parseLine(line string) (entry Entry) {
	parts := strings.Split(line, " | ")

	// why do I always have to initialize this "nil map"
	entry.patterns = map[string]Pattern{}

	for _, pattern := range strings.Fields(parts[0]) {
		// -1 is a mystery number
		sorted := utils.SortString(pattern)
		entry.patterns[sorted] = Pattern{
			length: len(sorted),
		}
	}

	for i, code := range strings.Fields(parts[1]) {
		sorted := utils.SortString(code)
		entry.digits[i] = sorted
	}

	return
}

// naive
func intersection(a, b string) string {
	achars := []rune(a)
	bchars := []rune(b)
	chars := []rune{}

	for _, x := range achars {
		for _, y := range bchars {
			if x == y {
				chars = append(chars, x)
			}
		}
	}

	return string(chars)
}

// naive
func contains(a, b string) bool {
	return len(intersection(a, b)) == len(a)
}

func (entry *Entry) getPatternByValue(value int) string {
	for pattern, pat := range entry.patterns {
		if pat.value == value {
			return pattern
		}
	}

	panic(fmt.Sprintf("no pattern for %d", value))
}

func (entry *Entry) getPatternsByLength(length int) (patterns []string) {
	for pattern, pat := range entry.patterns {
		if pat.length == length {
			patterns = append(patterns, pattern)
		}
	}

	if len(patterns) == 0 {
		panic(fmt.Sprintf("no patterns with length %d", length))
	}

	return
}

// kind of a mess working around UnaddressableFieldAssign
//
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal?utm_source=gopls#UnaddressableFieldAssign
func (entry *Entry) setPatternValue(pattern string, value int) {
	pat := entry.patterns[pattern]
	pat.value = value
	entry.patterns[pattern] = pat
}

// if any candidates contains all of pattern, then assign it the value
func (entry *Entry) deriveThroughIntersection(candidates []string, pattern string, value int) string {
	for _, candidate := range candidates {
		if contains(pattern, candidate) {
			entry.setPatternValue(candidate, value)

			return candidate
		}
	}
	panic(fmt.Sprintf("pattern not found: %s", pattern))
}

func removeFromStringArray(arr []string, value string) (out []string) {
	for _, val := range arr {
		if val != value {
			out = append(out, val)
		}
	}

	return
}

func (entry *Entry) derivePatterns() {
	// givens: 1, 4, 7, 8

	for pattern, pat := range entry.patterns {
		switch pat.length {
		case 2:
			pat.value = 1
		case 3:
			pat.value = 7
		case 4:
			pat.value = 4
		case 5:
			// 2, 3, 5
			pat.value = -1
		case 6:
			// 0, 6, 9
			pat.value = -1
		case 7:
			pat.value = 8
		}

		entry.patterns[pattern] = pat
	}

	// 2, 3 & 5 are all of length 5, but only 3 includes all segments from 1
	lenfive := entry.getPatternsByLength(5)
	matched := entry.deriveThroughIntersection(lenfive, entry.getPatternByValue(1), 3)
	lenfive = removeFromStringArray(lenfive, matched)

	// 9 contains 4
	lensix := entry.getPatternsByLength(6)
	matched = entry.deriveThroughIntersection(lensix, entry.getPatternByValue(4), 9)
	lensix = removeFromStringArray(lensix, matched)

	// 0 contains 1
	matched = entry.deriveThroughIntersection(lensix, entry.getPatternByValue(1), 0)
	lensix = removeFromStringArray(lensix, matched)

	// 6 remains
	six := lensix[0]
	entry.setPatternValue(six, 6)

	// 2 & 5 are left: 6 contains 5
	for _, candidate := range lenfive {
		// reversed from above
		if contains(candidate, six) {
			entry.setPatternValue(candidate, 5)
		} else {
			// only 2 remains
			entry.setPatternValue(candidate, 2)
		}
	}
}

// we concatenate each entry
func (entry *Entry) concatDigits() int {
	sum := ""
	for _, pattern := range entry.digits {
		sum += strconv.Itoa(entry.patterns[pattern].value)
	}

	out, _ := strconv.Atoi(sum)

	return out
}

func PartTwo(lines []string) (output int, err error) {
	for _, line := range lines {
		entry := parseLine(line)
		entry.derivePatterns()

		output += entry.concatDigits()
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
