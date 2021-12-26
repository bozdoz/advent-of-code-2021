package main

import (
	"fmt"

	"bozdoz.com/aoc-2021/utils"
)

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	isOn                   bool
	intersections          []*Cube
}

type Grid struct {
	cubes []*Cube
}

func (grid *Grid) parseInstructions(data []string, shouldClamp bool) {
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

		grid.cubes = append(grid.cubes, cube)
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

func countLitCubes(cubes []*Cube) (sum int) {
	// resolve intersections backwards
	for i := len(cubes) - 1; i >= 0; i-- {
		cube := cubes[i]

		// TODO: really?!
		if !cube.isOn {
			continue
		}

		// get all overlapping cubes (forwards)
		for _, next := range cubes[i+1:] {
			if next == nil {
				continue
			}

			intersection := cube.intersection(next)

			if intersection == nil {
				continue
			}

			// TODO: always takes the parent property?
			intersection.isOn = true

			// if there is an intersection save it, and
			// reverse all intersections of the cubes intersections
			cube.intersections = append(cube.intersections, intersection)
		}

		sum += cube.volume()
		sum -= countLitCubes(cube.intersections)
	}

	return
}

// TODO: would love to do a spatial index on these cubes
func (grid *Grid) count() int {
	return countLitCubes(grid.cubes)
}
