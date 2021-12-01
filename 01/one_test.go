package main

import (
	"testing"

	"bozdoz.com/aoc-2021/utils"
)

func TestLoading(t *testing.T) {
	expected := 10
	vals := utils.LoadInts("example.txt")

	if len(vals) != expected {
		t.Logf("example.txt should have %d ints", expected)
		t.Log(vals)
		t.Fail()
	}
}

func TestExampleOne(t *testing.T) {
	expected := 7
	vals := utils.LoadInts("example.txt")
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
	expected := 5
	ints := utils.LoadInts("example.txt")
	val, err := PartTwo(ints)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}
}
