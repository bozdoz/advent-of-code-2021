package main

import (
	"testing"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 26397,
	2: 288957,
}

var vals = FileLoader("example.txt")

func TestExampleOne(t *testing.T) {
	expected, ok := answers[1]

	if !ok {
		return
	}

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
