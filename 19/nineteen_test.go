package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/bozdoz/advent-of-code-2021/19/scanner2d"
	"github.com/bozdoz/advent-of-code-2021/19/scanner3d"
	"github.com/bozdoz/advent-of-code-2021/types"
)

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

	composite := scanners[0]

	newBeacons, matchedBeacons, scannerPos := composite.CompareScanner(scanners[1])

	expected := 12

	if matchedBeacons != expected {
		t.Logf("expected %v beacons to match, but got %v", expected, matchedBeacons)
		t.Fail()
	}

	if scannerPos.X != 68 || scannerPos.Y != -1246 || scannerPos.Z != -43 {
		t.Errorf("expected %v got %v", []int{68, -1246, -43}, scannerPos)
	}

	composite.AddBeacons(newBeacons)

	newBeacons, matchedBeacons, scannerPos = composite.CompareScanner(scanners[4])

	if matchedBeacons != expected {
		t.Logf("expected %v beacons to match scanner 4, but got %v", expected, matchedBeacons)
		t.Fail()
	}

	if scannerPos.X != -20 || scannerPos.Y != -1133 || scannerPos.Z != 1061 {
		t.Errorf("[s4] expected %v got %v", []int{-20, -1133, 1061}, scannerPos)
	}
}

func Test3dFull(t *testing.T) {
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

		newBeacons, count, _ := composite.CompareScanner(scanner)

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

func TestManhattan(t *testing.T) {
	a := types.NewVector3d(1105, -1205, 1229)
	b := types.NewVector3d(-92, -2380, -20)

	distance := scanner3d.ManhattanDistance(a, b)

	if distance != 3621 {
		t.Errorf("expected %v, got %v", 3621, distance)
	}
}

func TestPartTwo(t *testing.T) {
	vals := FileLoader("example3d.txt")

	answer, err := PartTwo(vals)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if answer != 3621 {
		t.Errorf("expected %v, got %v", 3621, answer)
	}
}
