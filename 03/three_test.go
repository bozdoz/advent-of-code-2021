package main

import (
	"testing"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

func TestExampleOne(t *testing.T) {
	expected := 198
	vals := utils.LoadAsLines("example.txt")
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
	expected := 230
	vals := utils.LoadAsLines("example.txt")
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
