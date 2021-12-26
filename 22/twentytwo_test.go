package main

import (
	"os"
	"testing"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 590784,
	2: 2758514936282235,
}

var vals = FileLoader("example.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func makeCube(args ...int) *Cube {
	return &Cube{
		x1: args[0],
		x2: args[1],
		y1: args[2],
		y2: args[3],
		z1: args[4],
		z2: args[5],
	}
}

func TestVolume(t *testing.T) {
	cube := makeCube(10, 10, 10, 10, 10, 10)
	vol := cube.volume()

	if vol != 1 {
		t.Logf("got %v, wanted %v", vol, 1)
		t.FailNow()
	}

	cube = makeCube(10, 10, 10, 10, 10, 11)
	vol = cube.volume()

	if vol != 2 {
		t.Logf("got %v, wanted %v", vol, 2)
		t.FailNow()
	}

	cube = makeCube(10, 10, 10, 11, 10, 11)
	vol = cube.volume()

	if vol != 4 {
		t.Logf("got %v, wanted %v", vol, 4)
		t.FailNow()
	}

	cube = makeCube(10, 11, 10, 11, 10, 11)
	vol = cube.volume()

	if vol != 8 {
		t.Logf("got %v, wanted %v", vol, 8)
		t.FailNow()
	}
}

func getVol(x1, x2, y1, y2, z1, z2 int) int {
	return makeCube(x1, x2, y1, y2, z1, z2).volume()
}

func FuzzVolume(f *testing.F) {
	f.Add(10, 10, 10, 10, 10, 10)

	f.Fuzz(func(t *testing.T, x1, x2, y1, y2, z1, z2 int) {
		out := getVol(x1, x2, y1, y2, z1, z2)

		if out < 1 {
			t.Fatalf("shouldn't have a negative cube, got %v", out)
		}
	})
}

func TestExampleSmall(t *testing.T) {
	expected := 39

	grid := &Grid{}

	grid.parseInstructions(FileLoader("examplesmall.txt"), true)

	count := grid.count()

	if count != expected {
		t.Logf("%v not equal %v", count, expected)
		t.Fail()
	}
}

func TestLightsOn(t *testing.T) {
	grid := &Grid{}
	input := []string{}
	expected := makeCube(-41, 9, -7, 43, -33, 15).volume()

	for range make([]int, 10) {
		input = append(input, "on x=-41..9,y=-7..43,z=-33..15")
	}

	grid.parseInstructions(input, true)
	count := grid.count()

	if count != expected {
		t.Logf("expected %v, got %v", expected, count)
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

	val, err := PartTwo(FileLoader("examplelarge.txt"))

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}
