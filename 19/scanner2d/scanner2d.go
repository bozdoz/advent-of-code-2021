package scanner2d

import (
	"fmt"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

type Beacon2d struct {
	position utils.Vector[int]
	edges    map[string]bool
}

type Scanner struct {
	beacons []*Beacon2d
	name    string
}

// custom logger extended from the "log" package
var log = utils.Logger("2d")

func ParseScanners(data []string) []*Scanner {
	scanners := []*Scanner{}
	var curScanner *Scanner

	for _, line := range data {
		switch {
		case strings.Contains(line, "---"):
			curScanner = &Scanner{
				name: strings.Trim(line, "- "),
			}
			scanners = append(scanners, curScanner)
		case strings.TrimSpace(line) == "":
			continue
		default:
			var x, y int
			count, err := fmt.Sscanf(line, "%d,%d", &x, &y)

			if count != 2 || err != nil {
				panic(fmt.Sprint("Could not parse beacon x,y,z: ", err))
			}

			curScanner.beacons = append(curScanner.beacons, &Beacon2d{
				position: utils.Vector[int]{X: x, Y: y},
			})
		}
	}

	for _, scanner := range scanners {
		scanner.updateEdges()
	}

	return scanners
}

func (scanner *Scanner) updateEdges() {
	for _, a := range scanner.beacons {
		a.edges = map[string]bool{}
		for _, b := range scanner.beacons {
			if a == b {
				continue
			}

			addEdge(a, b)
		}
	}
}

const minSharedBeacons = 3

func (this *Beacon2d) sharesAtLeastNEdges(n int, beacon *Beacon2d) bool {
	shared := 0
	remaining := len(this.edges)

	for edge := range this.edges {
		remaining--
		_, ok := beacon.edges[edge]

		if ok {
			shared++
		}

		if shared == n {
			return true
		}

		if (remaining + shared) < n {
			// can't possibly find enough
			return false
		}
	}

	return false
}

func (this *Scanner) CompareScanner(scanner *Scanner) []*Beacon2d {
	shared := 0

	// make a copy to alter the list within the loop
	unmatched := make([]*Beacon2d, len(scanner.beacons))
	copy(unmatched, scanner.beacons)
	remaining := len(this.beacons)

	var selfBeacon, altBeacon *Beacon2d

	// compare all scanners' beacons' edges
	for _, a := range this.beacons {
		if (remaining + shared) < minSharedBeacons {
			// not enough beacons to be a match
			log.Println("not enough beacons to match", remaining, shared, minSharedBeacons)
			break
		}
		remaining--
		for key, b := range unmatched {
			// does this beacon have a crossover of edges
			// if scanners share 3 beacons, the beacons will share 2 edges
			if a.sharesAtLeastNEdges(minSharedBeacons-1, b) {
				shared++
				log.Println("found shared beacon", a, b)
				// remove from list
				unmatched = append(unmatched[:key], unmatched[key+1:]...)

				// save to find transformation later
				selfBeacon = a
				altBeacon = b

				break
			}
		}
	}

	log.Println("found shared beacons: ", shared)

	didMatch := shared >= minSharedBeacons
	newBeacons := []*Beacon2d{}

	if didMatch {
		// need to transform unmatched
		for _, b := range unmatched {
			// compare unmatched to an edge in the same orientation
			diff := b.position.Subtract(altBeacon.position)
			newPosition := selfBeacon.position.Add(diff)

			log.Println("unmatched", b.position, newPosition)

			newBeacons = append(newBeacons, &Beacon2d{
				position: newPosition,
			})
		}
	}

	// returns new Beacons in the correct projection if enough matched
	return newBeacons
}

func addEdge(a *Beacon2d, b *Beacon2d) {
	edge := a.position.Subtract(b.position)
	key := edge.ToString()

	a.edges[key] = true
}

func (scanner *Scanner) AddBeacons(beacons []*Beacon2d) {
	scanner.beacons = append(scanner.beacons, beacons...)

	// update edges just for new beacons
	for _, a := range beacons {
		a.edges = map[string]bool{}
		for _, b := range scanner.beacons {
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
	output += fmt.Sprintf("\n--- %s ---", scanner.name)

	for _, beacon := range scanner.beacons {
		output += fmt.Sprintf("\n %v", beacon.position)
		output += fmt.Sprintf("\n %v", beacon.edges)
	}
	output += "\n"

	return
}
