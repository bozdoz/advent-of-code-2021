package main

import (
	"sort"
	"strings"
)

// Path is a route through connected caves
type Path []*Cave

// Paths is a list of all viable paths through connected caves
type Paths []Path

func (path *Path) has(cave *Cave) bool {
	for _, ref := range *path {
		// pointers for the win!?
		if cave == ref {
			return true
		}
	}

	return false
}

// all paths you find should visit small caves at most once
// ...except Part Two where a single small cave
// can be visited at most twice!
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

// custom string representation of Path
func (path *Path) String() (output string) {
	paths := []string{}
	for _, cave := range *path {
		paths = append(paths, cave.name)
	}

	return strings.Join(paths, ",")
}

// custom string representation of Paths
func (paths *Paths) String() (output string) {
	lines := []string{}
	for _, path := range *paths {
		lines = append(lines, path.String())
	}

	sort.Strings(lines)

	return strings.Join(lines, "\n")
}
