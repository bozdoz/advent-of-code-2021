package main

import (
	"os"
	"testing"
)

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestTicking(t *testing.T) {
	probe := newProbe(0, 0, 7, 2)

	probe.tick()

	expectedPosition := Vector[int]{7, 2}
	expectedVelocity := Vector[int]{6, 1}

	if !probe.position.isEqualTo(expectedPosition) {
		t.Logf("Answer should be %v, but got %v", expectedPosition, probe.position)
		t.Fail()
	}

	if !probe.velocity.isEqualTo(expectedVelocity) {
		t.Logf("Answer should be %v, but got %v", expectedPosition, probe.velocity)
		t.Fail()
	}

	// tick again
	probe.tick()

	expectedPosition = Vector[int]{13, 3}
	expectedVelocity = Vector[int]{5, 0}

	if !probe.position.isEqualTo(expectedPosition) {
		t.Logf("Answer should be %v, but got %v", expectedPosition, probe.position)
		t.Fail()
	}

	if !probe.velocity.isEqualTo(expectedVelocity) {
		t.Logf("Answer should be %v, but got %v", expectedPosition, probe.velocity)
		t.Fail()
	}
}

func TestVectorAdd(t *testing.T) {
	a := Vector[int]{2, 3}
	b := Vector[int]{-1, 10}

	a.add(b)
	expected := Vector[int]{1, 13}

	if !a.isEqualTo(expected) {
		t.Logf("Answer should be %v, but got %v", expected, a)
		t.Fail()
	}

	// b is unchanged
	expected = Vector[int]{-1, 10}

	if !b.isEqualTo(expected) {
		t.Logf("Answer should be %v, but got %v", expected, b)
		t.Fail()
	}
}

func TestTargetContain(t *testing.T) {
	target := Target{20, 30, -10, -5}
	good := []Vector[int]{
		{20, -5},
		{25, -7},
		{30, -10},
	}

	bad := []Vector[int]{
		{19, -5},
		{31, -5},
		{15, -4},
		{19, -4},
	}

	for _, actual := range good {
		if !target.contains(actual) {
			t.Logf("Answer should be %v, but wasn't: %v", true, actual)
			t.Fail()
		}
	}

	for _, actual := range bad {
		if target.contains(actual) {
			t.Logf("Answer should be %v, but wasn't: %v", false, actual)
			t.Fail()
		}
	}
}

func TestAngle(t *testing.T) {
	vec := Vector[int]{5, 5}

	angle := vec.angleDegrees()

	if angle != 45 {
		t.Logf("Answer should be %v, but wasn't: %v", 45, angle)
		t.Fail()
	}

	vec = Vector[int]{0, 5}

	angle = vec.angleDegrees()

	if angle != 90 {
		t.Logf("Answer should be %v, but wasn't: %v", 90, angle)
		t.Fail()
	}

	vec = Vector[int]{5, 0}

	angle = vec.angleDegrees()

	if angle != 0 {
		t.Logf("Answer should be %v, but wasn't: %v", 0, angle)
		t.Fail()
	}
}

func TestProbeLaunch(t *testing.T) {
	target := Target{20, 30, -10, -5}
	probe := newProbe(0, 0, 20, -10)
	hit := probe.isLaunchSuccessful(&target)
	if !hit {
		t.Logf("Answer should be %v, but wasn't: %v", true, hit)
		t.Fail()
	}
}

func TestPartOne(t *testing.T) {
	content := FileLoader("example.txt")
	val, err := PartOne(content)
	expected := 45

	if err != nil {
		t.Log("Expected no error, got:", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Expected %v, got %v", expected, val)
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	content := FileLoader("example.txt")
	val, err := PartTwo(content)
	expected := 112

	if err != nil {
		t.Log("Expected no error, got:", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Expected %v, got %v", expected, val)
		t.Fail()
	}
}

func TestPartTwoFailedExample(t *testing.T) {
	content := FileLoader("example.txt")
	target := parseTarget(content)

	probe := newProbe(0, 0, 6, 0)

	hitTarget := probe.isLaunchSuccessful(target)

	if !hitTarget {
		t.Logf("Expected %v, got %v", true, hitTarget)
		t.Fail()
	}
}
