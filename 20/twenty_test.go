package main

import (
	"math"
	"strings"
	"testing"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 35,
	2: 3351,
}

var vals = FileLoader("example.txt")

func TestSinglePixel(t *testing.T) {
	data := FileLoader("input.txt")
	parts := utils.SplitByEmptyNewline(data)
	data = parts[0] + "\n\n" + "#"
	image, enhancer := parseInput(data)

	if len(image.pixels) != 1 {
		t.Log("image should have 1 pixel")
		t.Fail()
	}

	if image.litCount() != 1 {
		t.Log("image should have 1 lit pixel")
		t.Fail()
	}

	if image.infinitePixel != "." {
		t.Log("image infinite pixels should be '.'")
		t.Fail()
	}

	if image.nextInfinitePixel != "#" {
		t.Log("next image infinite pixels should be '#'")
		t.Fail()
	}

	newImage := image.enhance(enhancer)

	if newImage.infinitePixel != "#" {
		t.Log("newImage infinite pixels should be '#'")
		t.Fail()
	}

	if newImage.nextInfinitePixel != "." {
		t.Log("next newImage infinite pixels should be '.'")
		t.Fail()
	}

	stringified := strings.Join(strings.Split(newImage.String(), "\n"), "")

	for i, pixel := range stringified {
		index := int(math.Pow(2, float64(i)))
		char := enhancer[index]

		if char != byte(pixel) {
			t.Logf("expected %v, got %v", char, pixel)
			t.Fail()
		}
	}

	lastImage := newImage.enhance(enhancer)
	stringified = strings.Join(strings.Split(lastImage.String(), "\n"), "")
	val, _ := utils.BinaryToInt("111111110")

	if stringified[0] != enhancer[val] {
		t.Logf("expected %v, got %v", enhancer[val], stringified[0])
		t.Fail()
	}

	val, _ = utils.BinaryToInt("111111101")

	if stringified[1] != enhancer[val] {
		t.Logf("expected %v, got %v", enhancer[val], stringified[1])
		t.Fail()
	}

	val, _ = utils.BinaryToInt("111111011")

	if stringified[2] != enhancer[val] {
		t.Logf("expected %v, got %v", enhancer[val], stringified[2])
		t.Fail()
	}

	val, _ = utils.BinaryToInt("111111111")

	if stringified[3] != enhancer[val] {
		t.Logf("expected %v, got %v", enhancer[val], stringified[3])
		t.Fail()
	}

	if stringified[4] != enhancer[val] {
		t.Logf("expected %v, got %v", enhancer[val], stringified[4])
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

	val, err := PartTwo(vals)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo(vals)
	}
}
