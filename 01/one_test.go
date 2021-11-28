package main

import "testing"

func TestLoading(t *testing.T) {
	ints := LoadInts("example.txt")

	if len(ints) != 6 {
		t.Log("example.txt should have 6 ints", ints)
		t.Fail()
	}
}

func TestExampleOne(t *testing.T) {
	expected := 514579
	ints := LoadInts("example.txt")
	val, err := PartOne(ints)

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
	expected := 241861950
	ints := LoadInts("example.txt")
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
