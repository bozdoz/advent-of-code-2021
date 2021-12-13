package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

const (
	start = "start"
	end   = "end"
)

type Cave struct {
	name      string
	isBig     bool
	flowsInto []*Cave
}

type Path []*Cave
type Paths []Path

type CaveSystem struct {
	caves                    map[string]*Cave
	viewSingleSmallCaveTwice bool
	paths                    Paths
}

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

func isUpperCaseLetter(str string) bool {
	return unicode.IsUpper(rune(str[0]))
}

func (caveSys *CaveSystem) getCave(name string) (cave *Cave) {
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

func (cave *Cave) addConnection(nextCave *Cave) {
	cave.flowsInto = append(cave.flowsInto, nextCave)
}

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

func newCaveSystem(data []string) *CaveSystem {
	caveSys := &CaveSystem{}
	caveSys.caves = map[string]*Cave{}

	for _, connection := range data {
		caveNames := strings.Split(connection, "-")

		first, second := caveNames[0], caveNames[1]

		firstCave := caveSys.getCave(first)
		secondCave := caveSys.getCave(second)

		firstCave.connect(secondCave)
	}

	return caveSys
}

func (path *Path) has(cave *Cave) bool {
	for _, ref := range *path {
		if cave == ref {
			return true
		}
	}

	return false
}

// all paths you find should visit small caves at most once
// ...except Part Two where a single small cave
// can be revisited at most once
func (path *Path) canCaveBeVisited(cave *Cave, oneSmallCaveCanRevisit bool) bool {
	// big caves can always be revisited
	// and caves can always be visited once
	if cave.isBig || !path.has(cave) {
		return true
	}

	if !oneSmallCaveCanRevisit {
		// small caves can only be visited once
		return false
	}

	seen := map[*Cave]int{}

	for _, cave := range *path {
		// only care about small caves
		if !cave.isBig {
			seen[cave]++
		}
	}

	for _, count := range seen {
		if count > 1 {
			// one small cave has been revisited already
			return false
		}
	}

	// revisit this small cave
	return true
}

func (path *Path) String() (output string) {
	paths := []string{}
	for _, cave := range *path {
		paths = append(paths, cave.name)
	}

	return strings.Join(paths, ",")
}

func (paths *Paths) String() (output string) {
	lines := []string{}
	for _, path := range *paths {
		lines = append(lines, path.String())
	}

	sort.Strings(lines)

	return strings.Join(lines, "\n")
}

func (caveSys *CaveSystem) traverse(prevPath Path) {
	path := make(Path, len(prevPath))

	copy(path, prevPath)

	lastCave := path[len(path)-1]

	for _, nextCave := range lastCave.flowsInto {
		if !path.canCaveBeVisited(nextCave, caveSys.viewSingleSmallCaveTwice) {
			continue
		}

		updatedPath := append(path, nextCave)
		if nextCave.name == end {
			caveSys.paths = append(caveSys.paths, updatedPath)
		} else {
			caveSys.traverse(updatedPath)
		}
	}
}

func (caveSys *CaveSystem) findAllPaths() (count int) {
	// starts at start
	caveSys.traverse(Path{caveSys.caves[start]})

	return len(caveSys.paths)
}

func PartOne(content []string) (output int, err error) {
	caveSys := newCaveSystem(content)

	output = caveSys.findAllPaths()

	return
}

func PartTwo(content []string) (output int, err error) {
	caveSys := newCaveSystem(content)

	caveSys.viewSingleSmallCaveTwice = true

	output = caveSys.findAllPaths()

	return
}

func main() {
	// safe to assume
	filename := "input.txt"

	data := FileLoader(filename)

	answer, err := PartOne(data)

	if err != nil {
		fmt.Println("failed to parse PartOne", err)
		return
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(data)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
