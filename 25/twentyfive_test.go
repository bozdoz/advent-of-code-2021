package main

import (
	"os"
	"testing"
)

// example input
var vals = FileLoader("example.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestIsEmpty(t *testing.T) {
	input := []string{
		"..",
		"..",
	}

	grid := parseInput(input)

	// all should be equivalent of 0,0
	tests := [...][2]int{
		{0, 0},
		{0, 2},
		{2, 0},
		{2, 2},
	}

	for _, test := range tests {
		y, x := test[0], test[1]

		if !grid.isEmpty(y, x) {
			t.Errorf("expected %d,%d to be empty", y, x)
		}
	}

	grid.cucumbers[0][0] = DOWN

	// reverse the tests
	for _, test := range tests {
		y, x := test[0], test[1]

		if grid.isEmpty(y, x) {
			t.Errorf("expected %d,%d to be %b, got %b", y, x, DOWN, grid.cucumbers[0][0])
		}
	}
}

func TestMoveLeft(t *testing.T) {
	input := []string{
		">.",
		"..",
	}

	grid := parseInput(input)

	if !grid.isEmpty(0, 1) {
		t.Errorf("expected %v, got %v", "empty", grid.get(0, 1))
	}

	grid.moveLeft(0, 0)

	if grid.isEmpty(0, 1) {
		t.Errorf("expected %v, got %c", "not empty", grid.get(0, 1))
	}

	grid.moveLeft(0, 1)

	if !grid.isEmpty(0, 1) {
		t.Errorf("expected %v, got %v", "empty", grid.get(0, 1))
	}

	if grid.isEmpty(0, 0) {
		t.Errorf("expected %v, got %c", "not empty", grid.get(0, 0))
	}
}

func TestMoveDown(t *testing.T) {
	input := []string{
		"v.",
		"..",
	}

	grid := parseInput(input)

	if !grid.isEmpty(1, 0) {
		t.Errorf("expected %v, got %v", "empty", grid.get(1, 0))
	}

	grid.moveDown(0, 0)

	if grid.isEmpty(1, 0) {
		t.Errorf("expected %v, got %c", "not empty", grid.get(1, 0))
	}

	grid.moveDown(1, 0)

	if !grid.isEmpty(1, 0) {
		t.Errorf("expected %v, got %v", "empty", grid.get(1, 0))
	}

	if grid.isEmpty(0, 0) {
		t.Errorf("expected %v, got %c", "not empty", grid.get(0, 0))
	}
}

func TestStep(t *testing.T) {
	grid := parseInput(vals)

	grid.step()

	expected := `
....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`

	if grid.String() != expected {
		t.Errorf("expected: %s, got: %s", expected, grid.String())
	}
}

func TestWrapAround(t *testing.T) {
	t.Run("Vertical wraparound", func(t *testing.T) {
		input := []string{
			".v.",
			"...",
			".v.",
			".v.",
			".v.",
			"...",
			".v.",
		}

		grid := parseInput(input)

		grid.step()

		expected := `
...
.v.
.v.
.v.
...
.v.
.v.`

		if grid.String() != expected {
			t.Errorf("expected %v got %v", expected, grid.String())
		}
	})

	t.Run("Horizontal wraparound", func(t *testing.T) {
		input := []string{"...>"}

		grid := parseInput(input)

		grid.step()

		// new lines
		expected := `
>...`

		if grid.String() != expected {
			t.Errorf("expected %v got %v", expected, grid.String())
		}
	})
}

func TestStepsTilStopped(t *testing.T) {
	grid := parseInput(vals)

	steps := grid.stepsTilStopped()

	if steps != 58 {
		t.Errorf("Expected %v got %v", 58, steps)
	}
}
