package main

import (
	"fmt"
	"os"
	"testing"
)

var vals = FileLoader("input.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestBasic1(t *testing.T) {
	example := []string{
		"inp x",
		"mul x -1",
	}
	monad := parseInput(example)

	monad.input(4)

	got := monad.x
	want := -4

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestBasic2(t *testing.T) {
	example := []string{
		"inp z",
		"inp x",
		"mul z 3",
		"eql z x",
	}
	monad := parseInput(example)

	monad.input(3, 9)

	got := monad.z
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	monad.reset()

	monad.input(3, 6)

	got = monad.z
	want = 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestBasic3(t *testing.T) {
	example := []string{
		"inp w",
		"add z w",
		"mod z 2",
		"div w 2",
		"add y w",
		"mod y 2",
		"div w 2",
		"add x w",
		"mod x 2",
		"div w 2",
		"mod w 2",
	}
	monad := parseInput(example)

	monad.input(8 + 4 + 2 + 1)

	got := fmt.Sprintf("%d%d%d%d", monad.w, monad.x, monad.y, monad.z)
	want := "1111"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	monad.reset()
	monad.input(4 + 1)

	got = fmt.Sprintf("%d%d%d%d", monad.w, monad.x, monad.y, monad.z)
	want = "0101"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	monad.reset()
	monad.input(8)

	got = fmt.Sprintf("%d%d%d%d", monad.w, monad.x, monad.y, monad.z)
	want = "1000"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
