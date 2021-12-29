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
	input := "BCBDADCA"
	burrow := parseInput(input)

	// move from sideroom to hallway
	burrow.hallway[3] = burrow.siderooms[2][0]
	burrow.siderooms[2][0] = nil

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

	log.Println(burrow)

	if burrow.isHallwayClear(0, 3) {
		t.Log("expected hallway blocked 0 - 3, didn't get it")
		t.Fail()
	}

	if burrow.isHallwayClear(3, 10) {
		t.Log("expected hallway blocked 3 - 10, didn't get it")
		t.Fail()
	}
}

func TestSideRoomOtherTypes(t *testing.T) {
	sideroom := []*Amphipod{}

	if sideRoomHasOtherTypes(sideroom, AMBER) {
		t.Log("expected no other types, didn't get it")
		t.Fail()
	}

	sideroom = []*Amphipod{
		{_type: AMBER},
	}

	if sideRoomHasOtherTypes(sideroom, AMBER) {
		t.Log("expected no other types, didn't get it")
		t.Fail()
	}

	sideroom = []*Amphipod{
		{_type: AMBER},
		{_type: BRONZE},
	}

	if !sideRoomHasOtherTypes(sideroom, AMBER) {
		t.Log("expected other types, didn't get it")
		t.Fail()
	}

	sideroom = []*Amphipod{
		{_type: BRONZE},
		{_type: BRONZE},
	}

	if !sideRoomHasOtherTypes(sideroom, AMBER) {
		t.Log("expected other types, didn't get it")
		t.Fail()
	}
}

func TestFirstInSideRoom(t *testing.T) {
	sideroom := []*Amphipod{}

	first := getFirstPodInSideroom(sideroom)

	if first != nil {
		t.Log("expected nil in sideroom, didn't get it")
		t.Fail()
	}

	sideroom = []*Amphipod{
		nil,
		{_type: BRONZE},
	}

	first = getFirstPodInSideroom(sideroom)

	if first == nil || first._type != BRONZE {
		t.Log("expected bronze pod")
		t.Fail()
	}

	sideroom = []*Amphipod{
		{_type: BRONZE},
		{_type: AMBER},
	}

	first = getFirstPodInSideroom(sideroom)

	if first == nil || first._type != BRONZE {
		t.Log("expected bronze pod")
		t.Fail()
	}
}

func TestActivePods(t *testing.T) {
	// easy move AA and BB
	input := "BCCABDDA"
	burrow := parseInput(input)

	activePods := burrow.getActivePods()
	log.Println(activePods)

	if len(activePods) != 4 {
		t.Log("1. should start with 4 active pods, got", len(activePods))
		t.FailNow()
	}

	// move from sideroom to hallway
	burrow.swapSideRoomtoHallway(0, 0, 0)
	burrow.swapSideRoomtoHallway(0, 1, 1)
	// move A's
	burrow.swapSideRoomtoHallway(3, 0, 3)

	log.Println(burrow)

	activePods = burrow.getActivePods()
	log.Println(activePods)

	if len(activePods) != 4 {
		t.Log("2. should have 4 active pods, got", len(activePods))
		t.FailNow()
	}

	burrow.swapSideRoomtoHallway(3, 1, 5)

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

	if copy.siderooms[0][0] == burrow.siderooms[0][0] {
		t.Log("expected siderooms to have new copies")
		t.Fail()
	}

	twice := copy.Copy()

	if copy.siderooms[0][0] == twice.siderooms[0][0] {
		t.Log("expected siderooms to have new copies twice")
		t.Fail()
	}

	copy.swapSideRoomtoHallway(0, 0, 0)
	third := copy.Copy()

	if copy.siderooms[0][0] != third.siderooms[0][0] {
		t.Log("expected nil to be equal to nil")
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
	input := "BCCABDDA"
	burrow := parseInput(input)

	positions := burrow.getValidHallwayPositionsFromRoom(0)

	// 4 positions are always invalid
	expected := 11 - 4

	if len(positions) != expected {
		t.Logf("expected %v, got %v", expected, len(positions))
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

	burrow.swapSideRoomtoHallway(0, 0, 0)
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
