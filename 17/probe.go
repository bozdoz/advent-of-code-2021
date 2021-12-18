package main

type Probe struct {
	position, velocity, forceVec Vector[int]
	maxHeight                    int
}

type Target struct {
	xmin, xmax, ymin, ymax int
}

var forceVec = Vector[int]{
	x: -1,
	y: -1,
}

func newProbe(px, py, vx, vy int) *Probe {
	return &Probe{
		position: Vector[int]{px, py},
		velocity: Vector[int]{vx, vy},
		forceVec: forceVec,
	}
}

// advance to next step
func (probe *Probe) tick() {
	probe.position.add(probe.velocity)

	// "x velocity...does not change if it is already 0"
	if probe.velocity.x == 0 {
		probe.forceVec.x = 0
	}
	probe.velocity.add(probe.forceVec)

	if probe.position.y > probe.maxHeight {
		probe.maxHeight = probe.position.y
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

func (target *Target) contains(vec Vector[int]) bool {
	return vec.x <= target.xmax && vec.x >= target.xmin &&
		vec.y <= target.ymax && vec.y >= target.ymin
}

func (probe *Probe) missedTarget(target *Target) bool {
	return probe.position.x > target.xmax || probe.position.y < target.ymin
}
