package main

import "github.com/bozdoz/advent-of-code-2021/types"

type Probe struct {
	position, velocity, forceVec types.Vector[int]
	maxHeight                    int
}

var forceVec = types.NewVector(-1, -1)

func newProbe(px, py, vx, vy int) *Probe {
	return &Probe{
		position: types.NewVector(px, py),
		velocity: types.NewVector(vx, vy),
		forceVec: forceVec,
	}
}

// advance to next step
func (probe *Probe) tick() {
	probe.position = probe.position.Add(probe.velocity)

	// "x velocity...does not change if it is already 0"
	if probe.velocity.X == 0 {
		probe.forceVec.X = 0
	}
	probe.velocity = probe.velocity.Add(probe.forceVec)

	if probe.position.Y > probe.maxHeight {
		probe.maxHeight = probe.position.Y
	}
}

func (probe *Probe) isLaunchSuccessful(target *Target) bool {
	for !probe.missedTarget(target) {
		probe.tick()

		if probe.isInTarget(target) {
			return true
		}
	}

	return false
}

// in Target, shopping...
func (probe *Probe) isInTarget(target *Target) bool {
	return target.contains(probe.position)
}

func (probe *Probe) missedTarget(target *Target) bool {
	return probe.position.X > target.xmax || probe.position.Y < target.ymin
}
