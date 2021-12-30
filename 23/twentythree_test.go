package main

import (
	"os"
	"testing"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 12521,
	2: 0,
}

var vals = FileLoader("example.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestHallwayClear(t *testing.T) {
	input := `BCBD
					  ADCA`
	burrow := parseInput(input)

	// move from sideroom to hallway
	grid := burrow.grid
	burrow.movePodTo(grid[1][2], 3, 0)

	if !burrow.isHallwayClear(0, 2) {
		t.Log("expected hallway clear, didn't get it")
		t.Fail()
	}

	if !burrow.isHallwayClear(2, 0) {
		t.Log("expected hallway clear, didn't get it")
		t.Fail()
	}

	if !burrow.isHallwayClear(4, 10) {
		t.Log("expected hallway clear, didn't get it")
		t.Fail()
	}

	if !burrow.isHallwayClear(10, 4) {
		t.Log("expected hallway clear, didn't get it")
		t.Fail()
	}

	if burrow.isHallwayClear(0, 3) {
		t.Log("expected hallway blocked 0 - 3, didn't get it")
		t.Fail()
	}

	if burrow.isHallwayClear(3, 10) {
		t.Log("expected hallway blocked 3 - 10, didn't get it")
		t.Fail()
	}

	burrow = parseInput(input)

	grid = burrow.grid
	// move D
	burrow.movePodTo(grid[1][8], 9, 0)
	// move A
	burrow.movePodTo(grid[2][8], 5, 0)

	log.Println(burrow)

	burrow.getNextStates()
}

func TestActivePods(t *testing.T) {
	// easy move AA and BB
	input := `BCCA
					  BDDA`
	burrow := parseInput(input)

	activePods := burrow.getActivePods()
	log.Println(activePods)

	if len(activePods) != 4 {
		t.Log("1. should start with 4 active pods, got", len(activePods))
		t.FailNow()
	}

	grid := burrow.grid
	// move from sideroom to hallway
	burrow.movePodTo(grid[1][2], 0, 0)
	burrow.movePodTo(grid[2][2], 1, 0)

	// move A's
	burrow.movePodTo(grid[1][8], 3, 0)

	log.Println(burrow)

	activePods = burrow.getActivePods()
	log.Println(activePods)

	if len(activePods) != 4 {
		t.Log("2. should have 4 active pods, got", len(activePods))
		t.FailNow()
	}

	burrow.movePodTo(grid[2][8], 5, 0)

	log.Println(burrow)

	activePods = burrow.getActivePods()

	if len(activePods) != 2 {
		t.Log("3. should still only have 2 active pods, got", len(activePods))
		t.Fail()
	}
}

func TestBurrowCopy(t *testing.T) {
	input := "BCCABDDA"
	burrow := parseInput(input)

	copy := burrow.Copy()

	if copy.amphipods[0] == burrow.amphipods[0] {
		t.Log("expected amphipods have new copies")
		t.Fail()
	}
	if copy.amphipods[0].x != burrow.amphipods[0].x {
		t.Log("expected amphipods don't have same values", copy.amphipods[0], burrow.amphipods[0])
		t.Fail()
	}

	if copy.grid == burrow.grid {
		t.Logf("expected grid to be copied by value")
		t.Fail()
	}

	if copy.grid[1][2] == burrow.grid[1][2] {
		t.Logf("expected grid pods to be copied by value")
		t.Fail()
	}

	if copy.grid[1][2]._type != burrow.grid[1][2]._type {
		t.Logf("expected grid pods to have same value")
		t.Fail()
	}

	twice := copy.Copy()

	if copy.amphipods[0] == twice.amphipods[0] {
		t.Log("expected amphipods have new copies twice")
		t.Fail()
	}

	burrow.cost = 1234
	copy = burrow.Copy()

	if copy.cost != burrow.cost {
		t.Logf("expected cost to be %v, got %v", burrow.cost, copy.cost)
		t.Fail()
	}
}

func TestValidHallway(t *testing.T) {
	input := `BCCA
					  BDDA`
	burrow := parseInput(input)

	positions := burrow.getValidHallwayPositionsFromRoom(2)

	// 4 positions are always invalid
	expected := 11 - 4

	if len(*positions) != expected {
		t.Logf("expected %v, got %v", expected, len(*positions))
		t.Fail()
	}
}

func TestLastMove(t *testing.T) {
	input := `ABCD
					  ABCD`
	burrow := parseInput(input)

	activePods := burrow.getActivePods()
	expected := 0

	if len(activePods) != expected {
		t.Logf("expected %v, got %v", expected, len(activePods))
		t.Fail()
	}

	grid := burrow.grid
	burrow.movePodTo(grid[1][2], 0, 0)
	log.Println(burrow)

	activePods = burrow.getActivePods()
	expected = 1

	if len(activePods) != expected {
		t.Logf("expected %v, got %v", expected, len(activePods))
		t.Fail()
	}
}

func TestCost1(t *testing.T) {
	input := `ABCD
					  ABCD`
	burrow := parseInput(input)

	cost := burrow.play()

	if cost != 0 {
		t.Logf("expected %v, got %v", 0, cost)
		t.Fail()
	}
}

func TestCost2(t *testing.T) {
	input := `BACD
		 		    ABCD`
	burrow := parseInput(input)

	cost := burrow.play()
	// A moves 6, B moves 4
	expected := 1*6 + 10*4

	if cost != expected {
		t.Logf("expected %v, got %v", expected, cost)
		t.Fail()
	}
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
