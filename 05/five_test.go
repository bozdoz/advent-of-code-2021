package main

import (
	"testing"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 5,
	2: 12,
}

// different puzzles require different file loaders
var fileLoader = utils.LoadAsLines

func TestExampleOne(t *testing.T) {
	expected, ok := answers[1]

	if !ok {
		return
	}

	vals := fileLoader("example.txt")
	val, err := PartOne(vals)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}
}

func TestExampleTwo(t *testing.T) {
	expected, ok := answers[2]

	if !ok {
		return
	}

	vals := fileLoader("example.txt")
	val, err := PartTwo(vals)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}
}
