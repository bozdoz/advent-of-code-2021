package main

import (
	"fmt"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	isOn                   bool
}

type Cubes []*Cube

func (cubes *Cubes) parseInstructions(data []string, shouldClamp bool) {
	for _, line := range data {
		var ok bool
		var onoff string
		var x1, x2, y1, y2, z1, z2 int

		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &onoff, &x1, &x2, &y1, &y2, &z1, &z2)

		if shouldClamp {
			x1, x2, ok = getClampedDimensions(x1, x2)

			if !ok {
				continue
			}

			y1, y2, ok = getClampedDimensions(y1, y2)

			if !ok {
				continue
			}

			z1, z2, ok = getClampedDimensions(z1, z2)

			if !ok {
				continue
			}
		}

		cube := &Cube{
			x1:   x1,
			x2:   x2,
			y1:   y1,
			y2:   y2,
			z1:   z1,
			z2:   z2,
			isOn: onoff == "on",
		}

		*cubes = append(*cubes, cube)
	}

	return
}

func getClampedDimensions(v1, v2 int) (int, int, bool) {
	v1, v1ok := getClampedDimension(v1)
	v2, v2ok := getClampedDimension(v2)

	return v1, v2, v1ok && v2ok
}

func getClampedDimension(val int) (int, bool) {
	return utils.MaxInt(utils.MinInt(val, 50), -50), val >= -50 && val <= 50
}

func (cube *Cube) volume() int {
	vol := ((cube.x2 + 1) - cube.x1) *
		((cube.y2 + 1) - cube.y1) *
		((cube.z2 + 1) - cube.z1)

	if vol == 0 {
		// this never happens (except in fuzzing)
		return 1
	}

	if vol < 0 {
		// TODO: is this right?
		return -vol
	}

	return vol
}

func (cube *Cube) intersection(other *Cube) *Cube {
	x1 := utils.MaxInt(cube.x1, other.x1)
	x2 := utils.MinInt(cube.x2, other.x2)

	if x2 < x1 {
		return nil
	}

	y1 := utils.MaxInt(cube.y1, other.y1)
	y2 := utils.MinInt(cube.y2, other.y2)

	if y2 < y1 {
		return nil
	}

	z1 := utils.MaxInt(cube.z1, other.z1)
	z2 := utils.MinInt(cube.z2, other.z2)

	if z2 < z1 {
		return nil
	}

	return &Cube{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
		z1: z1,
		z2: z2,
	}
}

func (cubes *Cubes) count() (sum int) {
	// resolve intersections backwards
	for i := len(*cubes) - 1; i >= 0; i-- {
		cube := (*cubes)[i]

		// the volume of "off" cubes are never counted
		// only subtracted from the volumes of "on" cubes
		if !cube.isOn {
			continue
		}

		intersections := &Cubes{}

		// get all overlapping cubes (forwards)
		for _, next := range (*cubes)[i+1:] {
			intersection := cube.intersection(next)

			if intersection == nil {
				// did not intersect
				continue
			}

			// in recursive calls, "isOn" is synonymous with "shouldCountVolume"
			// as in, even if it's "off" we should calculate the total volume
			// of the intersections
			shouldCountVolume := true
			intersection.isOn = shouldCountVolume

			// if there is an intersection save it, and
			// reverse all intersections of the cubes intersections
			// i.e. don't count intersecting parts twice, and don't
			// discount intersecting parts twice.
			*intersections = append(*intersections, intersection)
		}

		sum += cube.volume()
		sum -= intersections.count()
	}

	return
}
