package utils

import (
	"math"
	"testing"
)

func TestSortString(t *testing.T) {
	alpha := "defcghba"
	expected := "abcdefgh"
	sorted := SortString(alpha)

	if sorted != expected {
		t.Logf("Answer should be %s, but got %s", expected, sorted)
		t.Fail()
	}
}

var vals = []int{5, 3, 4, 6, 7, 4, 3, 4, 6, 3, 1, 3, 4, 5}

const float64EqualityThreshold = 1e-6

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestMean(t *testing.T) {
	expected := 4.142857
	val := Mean(vals)

	if !almostEqual(val, expected) {
		t.Logf("Answer should be %f, but got %f", expected, val)
		t.Fail()
	}
}

func TestMedian(t *testing.T) {
	expected := 4.0
	val := Median(vals)

	if !almostEqual(val, expected) {
		t.Logf("Answer should be %f, but got %f", expected, val)
		t.Fail()
	}
}

func TestMode(t *testing.T) {
	expected := 3
	val := Mode(vals)

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	expected := 58
	val := Sum(vals...)

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}
}
