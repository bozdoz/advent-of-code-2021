package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	start = "start"
	end   = "end"
)

type Cave struct {
	name      string
	isBig     bool
	flowsInto []*Cave
}

type CaveSystem struct {
	caves                    map[string]*Cave
	viewSingleSmallCaveTwice bool
	paths                    Paths
}

// constructor-like function to create CaveSystem
func newCaveSystem(data []string) *CaveSystem {
	caveSys := &CaveSystem{}
	caveSys.caves = map[string]*Cave{}

	for _, connection := range data {
		caveNames := strings.Split(connection, "-")

		first, second := caveNames[0], caveNames[1]

		firstCave := caveSys.getOrCreateCave(first)
		secondCave := caveSys.getOrCreateCave(second)

		firstCave.connect(secondCave)
	}

	return caveSys
}

func isUpperCaseLetter(str string) bool {
	return unicode.IsUpper(rune(str[0]))
}

// adds cave to cave system
func (caveSys *CaveSystem) getOrCreateCave(name string) (cave *Cave) {
	cave, ok := caveSys.caves[name]

	if !ok {
		cave = &Cave{
			name:  name,
			isBig: isUpperCaseLetter(name),
		}

		caveSys.caves[name] = cave
	}

	return
}

// Caves are connected, one flows into another (usually bi-directional)
func (cave *Cave) addConnection(nextCave *Cave) {
	cave.flowsInto = append(cave.flowsInto, nextCave)
}

// usually calls addConnection bi-directionally for each cave
func (cave *Cave) connect(nextCave *Cave) {
	// no cave flows into start
	// and end doesn't flow into any cave
	if nextCave.name != start && cave.name != end {
		cave.addConnection(nextCave)
	}

	// also connect the reverse
	if cave.name != start && nextCave.name != end {
		nextCave.addConnection(cave)
	}
}

// recursive function to continually navigate through connected caves
// saves all paths that end in "end" to the CaveSystem.paths
func (caveSys *CaveSystem) traverse(path Path) {
	lastCave := path[len(path)-1]

	for _, nextCave := range lastCave.flowsInto {
		if !path.canCaveBeVisited(nextCave, caveSys.viewSingleSmallCaveTwice) {
			continue
		}

		updatedPath := append(path, nextCave)
		if nextCave.name == end {
			// nested append with the spread syntax creates a copy of the slice
			caveSys.paths = append(caveSys.paths, append(Path{}, updatedPath...))
		} else {
			caveSys.traverse(updatedPath)
		}
	}
}

// begin finding all paths in the cave system, by starting at the start
func (caveSys *CaveSystem) findAllPaths() (count int) {
	caveSys.traverse(Path{caveSys.caves[start]})

	return len(caveSys.paths)
}

// custom string representation for CaveSystem
func (caveSys *CaveSystem) String() (output string) {
	output += "CaveSystem {"
	for _, cave := range caveSys.caves {
		name := "\n  " + cave.name
		if cave.isBig {
			name += " (big)"
		}

		// TODO: can't get this to be right-aligned
		output += fmt.Sprintf("%-10s: ", name)

		flows := []string{}

		for _, nextCave := range cave.flowsInto {
			flows = append(flows, nextCave.name)
		}

		output += strings.Join(flows, ", ")
	}
	output += "\n}"
	return
}
