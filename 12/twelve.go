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

type CaveSystem struct {
	caves map[string]*Cave
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

type Path []Cave

func (path *Path) has(cave Cave) bool {
	for _, ref := range *path {
		if cave.name == ref.name {
			return true
		}
	}

	return false
}

type Paths []Path

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

func (paths *Paths) addCave(prevPath Path) {
	path := make(Path, len(prevPath))

	copy(path, prevPath)

	lastCave := path[len(path)-1]

	for _, cave := range lastCave.flowsInto {
		nextCave := *cave
		// all paths you find should visit small caves at most once
		if !nextCave.isBig && path.has(nextCave) {
			continue
		}

		updatedPath := append(path, nextCave)
		fmt.Println("updatedPath", updatedPath.String())
		if nextCave.name == end {
			if len(*paths) < 1 {
				func(paths *Paths, updatedPath Path) {
					*paths = append(*paths, updatedPath)
				}(paths, updatedPath)
			} else {
				return
			}
			fmt.Println("---", paths, "---")
		} else {
			paths.addCave(updatedPath)
		}
	}
}

func (caveSys *CaveSystem) findAllPaths() (count int) {
	paths := &Paths{}
	// starts at start
	paths.addCave(Path{*caveSys.caves[start]})

	fmt.Println("---", paths, "---")

	return
}

func PartOne(content []string) (output int, err error) {
	caveSys := newCaveSystem(content)

	fmt.Println(caveSys)

	caveSys.findAllPaths()

	return
}

func PartTwo(content []string) (output int, err error) {
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
