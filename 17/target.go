package main

import (
	"fmt"

	"bozdoz.com/aoc-2021/utils"
)

// Target is basically a bbox
type Bbox struct {
	xmin, xmax, ymin, ymax int
}
type Target Bbox

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

func (target *Target) contains(vec utils.Vector[int]) bool {
	return vec.X <= target.xmax && vec.X >= target.xmin &&
		vec.Y <= target.ymax && vec.Y >= target.ymin
}
