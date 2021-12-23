package main

import (
	"os"
	"testing"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 739785,
	2: 444356092776315,
}

var vals = FileLoader("example.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

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
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}

func TestCache(t *testing.T) {
	if cachedPossibleUniverses != nil {
		t.Log("cached doesn't start nil")
		t.FailNow()
	}

	cached := getAllPossibleUniverses()

	if cachedPossibleUniverses == nil {
		t.Log("cached should no longer be nil")
		t.FailNow()
	}

	if cachedPossibleUniverses != cached {
		t.Log("cache should be identical")
		t.FailNow()
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
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}
