package scanner3d

import (
	"fmt"
	"sort"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

const minSharedBeacons = 12

type Beacon3d struct {
	position utils.Vector3d
	edges    map[string][]utils.Vector3d
}

type Scanner struct {
	Beacons []*Beacon3d
	Name    string
}

// custom logger extended from the "log" package
var log = utils.Logger("3d")

func ParseScanners(data []string) []*Scanner {
	scanners := []*Scanner{}
	var curScanner *Scanner

	for _, line := range data {
		switch {
		case strings.Contains(line, "---"):
			curScanner = &Scanner{
				Name: strings.Trim(line, "- "),
			}
			scanners = append(scanners, curScanner)
		case strings.TrimSpace(line) == "":
			continue
		default:
			var x, y, z int
			count, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

			if count != 3 || err != nil {
				panic(fmt.Sprint("Could not parse beacon x,y,z: ", err))
			}

			curScanner.Beacons = append(curScanner.Beacons, &Beacon3d{
				position: utils.Vector3d{X: x, Y: y, Z: z},
			})
		}
	}

	for _, scanner := range scanners {
		scanner.updateEdges()
	}

	return scanners
}

func (scanner *Scanner) updateEdges() {
	for _, a := range scanner.Beacons {
		a.edges = map[string][]utils.Vector3d{}
		for _, b := range scanner.Beacons {
			if a == b {
				continue
			}

			addEdge(a, b)
		}
	}
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// removes negative signs so we can get the un-oriented numbers
// for somewhat more efficient searching
func getVectorKey(vector utils.Vector3d) string {
	sorted := []int{absInt(vector.X), absInt(vector.Y), absInt(vector.Z)}
	sort.Ints(sorted)
	return fmt.Sprint(sorted[0], sorted[1], sorted[2])
}

// edges could have different orientations: +-x,+-y,+-z
// AND the number could appear anywhere
func (this *Beacon3d) findEdge(edge utils.Vector3d) bool {
	edgeStr := getVectorKey(edge)
	_, ok := this.edges[edgeStr]

	return ok
}

func (this *Beacon3d) sharesAtLeastNEdges(n int, beacon *Beacon3d) bool {
	shared := 0

	for _, arr := range this.edges {
		for _, edge := range arr {
			ok := beacon.findEdge(edge)

			if ok {
				shared++
			}

			if shared == n {
				return true
			}
		}
	}

	return false
}

// returns new Beacons in the correct projection if enough matched
func (this *Scanner) CompareScanner(scanner *Scanner) (newBeacons []*Beacon3d, shared int) {
	// make a copy to alter the list within the loop
	unmatched := make([]*Beacon3d, len(scanner.Beacons))
	copy(unmatched, scanner.Beacons)
	remaining := len(this.Beacons)

	var selfBeacon, altBeacon *Beacon3d

	// compare all scanners' beacons' edges
	for _, a := range this.Beacons {
		// if a.position.X != -618 {
		// 	continue
		// }
		if (remaining + shared) < minSharedBeacons {
			// not enough beacons to be a match
			break
		}
		remaining--
		for key, b := range unmatched {
			// does this beacon have a crossover of edges
			// if scanners share 12 beacons, the beacons will share 11 edges
			if a.sharesAtLeastNEdges(minSharedBeacons-1, b) {
				shared++
				// remove from list
				unmatched = append(unmatched[:key], unmatched[key+1:]...)

				// save to find transformation later
				selfBeacon = a
				altBeacon = b

				break
			}
		}
	}

	didMatch := shared >= minSharedBeacons

	if didMatch {
		// need to transform unmatched
		for _, b := range unmatched {
			// compare unmatched to an edge in the same orientation
			diff := b.position.Subtract(altBeacon.position)
			newPosition := selfBeacon.position.Add(diff)

			newBeacons = append(newBeacons, &Beacon3d{
				position: newPosition,
			})
		}
	}

	return
}

func addEdge(a *Beacon3d, b *Beacon3d) {
	edge := a.position.Subtract(b.position)

	key := getVectorKey(edge)

	if a.edges == nil {
		// TODO: still no idea when this happens
		a.edges = map[string][]utils.Vector3d{}
	}

	a.edges[key] = append(a.edges[key], edge)
}

func (scanner *Scanner) AddBeacons(beacons []*Beacon3d) {
	scanner.Beacons = append(scanner.Beacons, beacons...)

	// update edges just for new beacons
	for _, a := range beacons {
		a.edges = map[string][]utils.Vector3d{}
		for _, b := range scanner.Beacons {
			if a == b {
				continue
			}
			addEdge(a, b)
			addEdge(b, a)
		}
	}
}

//
// String Representations
//

func (scanner *Scanner) String() (output string) {
	output += fmt.Sprintf("\n--- %s ---", scanner.Name)

	for _, beacon := range scanner.Beacons {
		output += beacon.String()
	}
	output += "\n"

	return
}

func (beacon *Beacon3d) String() (output string) {
	output += "\n"
	output += fmt.Sprintf("position: %v\n", beacon.position)
	output += "edges:\n"

	for _, edge := range beacon.edges {
		output += fmt.Sprintf("%v ", edge)
	}

	return
}
