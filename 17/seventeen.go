package main

import (
	"fmt"
	"io/ioutil"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func parseTarget(data string) *Target {
	var xmin, xmax, ymin, ymax int
	count, err := fmt.Sscanf(data, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

	if count != 4 || err != nil {
		panic(fmt.Sprint("expected 4 values, found:", count, "\nError: \n", err))
	}

	return &Target{xmin, xmax, ymin, ymax}
}

// 6 would hit 21, because 1+2+3+4+5+6 == 21
// (drag decreases vel by 1 each tick)
func findMinXVelocity(num int) (min int) {
	inc := 0
	for min < num {
		inc += 1
		min += inc
	}

	return inc
}

// get all candidates for velocities and try them out
func (target *Target) practice(forEach func(probe *Probe, success bool)) {
	// shooting right at it, to hit it first tick
	xmax := target.xmax
	xmin := findMinXVelocity(target.xmin)

	// y will always come back down to 0 at the same velocity
	// that it went up with.  so ymin -5 means that max velocity is 4.
	// the tick that brings it to 0 will be velocity - 4, and the one
	// after that is -5 (the ymin).  Same with ymax.
	ymax := -target.ymin - 1
	ymin := target.ymin

	log.Println(xmax, xmin, ymax, ymin)

	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			probe := newProbe(0, 0, x, y)

			forEach(probe, probe.isLaunchSuccessful(target))
		}
	}
}

func PartOne(content string) (output int, err error) {
	target := parseTarget(content)

	log.Println(target)

	maxHeight := 0

	target.practice(func(probe *Probe, success bool) {
		if success && probe.maxHeight > maxHeight {
			maxHeight = probe.maxHeight
		}
	})

	return maxHeight, nil
}

func PartTwo(content string) (output int, err error) {
	target := parseTarget(content)

	log.Println(target)

	hitCount := 0

	target.practice(func(_ *Probe, success bool) {
		if success {
			hitCount++
		}
	})

	return hitCount, nil
}

func main() {
	// safe to assume
	filename := "input.txt"

	data := FileLoader(filename)

	answer, err := PartOne(data)

	if err != nil {
		fmt.Println("failed to parse PartOne", err)
		return
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(data)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
