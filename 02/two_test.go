package main

import (
	"testing"

	"bozdoz.com/aoc-2021/utils"
)

func TestLoading(t *testing.T) {
	expected := 3
	vals := utils.LoadFile("example.txt")

	if len(vals) != expected {
		t.Log("example.txt should have 6 vals", vals)
		t.Fail()
	}
}

func TestExampleOne(t *testing.T) {
	expected := 2
	vals := utils.LoadFile("example.txt")
	val, err := PartOne(vals)
	t.Log("done")

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
	expected := 1
	vals := utils.LoadFile("example.txt")
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
