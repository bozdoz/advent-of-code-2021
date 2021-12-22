package main

import (
	"fmt"
	"os"
	"testing"

	"bozdoz.com/aoc-2021/19/scanner2d"
	"bozdoz.com/aoc-2021/19/scanner3d"
	"bozdoz.com/aoc-2021/types"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 0,
	2: 0,
}

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func Test2d(t *testing.T) {
	scanners := scanner2d.ParseScanners(FileLoader("example2d.txt"))

	unmatched := scanners[0].CompareScanner(scanners[1])

	if len(unmatched) != 1 {
		t.Logf("expected %v, got %v", 1, len(unmatched))
		t.Fail()
	}

	composite := scanners[0]

	composite.AddBeacons(unmatched)

	log.Println(composite)
}

// Scanners 0 and 1 have overlapping detection cubes;
// the 12 beacons they both detect (relative to scanner 0) are
func Test3dFirst(t *testing.T) {
	scanners := scanner3d.ParseScanners(FileLoader("example3d.txt"))

	_, matchedBeacons := scanners[0].CompareScanner(scanners[1])

	expected := 12

	if matchedBeacons != expected {
		t.Logf("expected %v beacons to match, but got %v", expected, matchedBeacons)
		t.Fail()
	}
}

func Test3dSecond(t *testing.T) {
	scanners := scanner3d.ParseScanners(FileLoader("example3d.txt"))

	composite := scanners[1]

	_, matchedBeacons := composite.CompareScanner(scanners[4])

	expected := 12

	if matchedBeacons != expected {
		t.Logf("expected %v beacons to match, but got %v", expected, matchedBeacons)
		t.Fail()
	}
}

func needsReworking3dEqual(t *testing.T) {
	scanners := scanner3d.ParseScanners(FileLoader("exampleEqual.txt"))
	expected := 6

	composite := scanners[0]

	queue := types.Queue[scanner3d.Scanner]{}

	for _, scanner := range scanners[1:] {
		queue.Push(scanner)
		fmt.Println(scanner)
	}

	for len(queue) > 0 {
		scanner := queue.Shift()
		fmt.Println("comparing", scanner.Name)
		newBeacons, count := composite.CompareScanner(scanner)

		if count > 0 {
			composite.AddBeacons(newBeacons)
		} else {
			queue.Push(scanner)
		}
	}

	if len(composite.Beacons) != expected {
		t.Logf("expected %v beacons in total, but got %v", expected, len(composite.Beacons))
		t.Fail()
	}
}

func Test3dFull(t *testing.T) {
	// TODO: Day 19.1
	t.SkipNow()

	scanners := scanner3d.ParseScanners(FileLoader("example3d.txt"))
	expected := 79

	composite := scanners[0]

	queue := types.Queue[scanner3d.Scanner]{}

	for _, scanner := range scanners[1:] {
		queue.Push(scanner)
	}

	lastScanner := composite

	for len(queue) > 0 {
		scanner := queue.Shift()

		fmt.Println("comparing", scanner.Name)

		if scanner == lastScanner {
			fmt.Println("repeat scanner found beacons:", len(scanner.Beacons))
			break
		}

		lastScanner = scanner

		newBeacons, count := composite.CompareScanner(scanner)

		fmt.Println("found", count)

		if count > 0 {
			composite.AddBeacons(newBeacons)
			fmt.Println("total", len(composite.Beacons))
		} else {
			queue.Push(scanner)
		}
	}

	if len(composite.Beacons) != expected {
		t.Logf("expected %v beacons in total, but got %v", expected, len(composite.Beacons))
		t.Fail()
	}
}
