package scanner3d

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/bozdoz/advent-of-code-2021/types"
	"github.com/bozdoz/advent-of-code-2021/utils"
)

const minSharedBeacons = 12

type Beacon3d struct {
	position types.Vector3d
	edges    map[string][]types.Vector3d
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
				position: types.NewVector3d(x, y, z),
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
		a.edges = map[string][]types.Vector3d{}
		for _, b := range scanner.Beacons {
			if a == b {
				continue
			}

			addEdge(a, b)
		}
	}
}

// removes negative signs so we can get the un-oriented numbers
// for somewhat more efficient searching
func getVectorKey(vector types.Vector3d) string {
	sorted := []int{utils.Abs(vector.X), utils.Abs(vector.Y), utils.Abs(vector.Z)}
	sort.Ints(sorted)
	return fmt.Sprint(sorted[0], sorted[1], sorted[2])
}

// edges could have different orientations: +-x,+-y,+-z
// AND the number could appear anywhere
func (this *Beacon3d) findEdge(edge types.Vector3d) ([]types.Vector3d, bool) {
	edgeStr := getVectorKey(edge)

	edges, ok := this.edges[edgeStr]

	return edges, ok
}

func (this *Beacon3d) sharesAtLeastNEdges(n int, beacon *Beacon3d) bool {
	shared := 0

	for _, arr := range this.edges {
		for _, edge := range arr {
			_, ok := beacon.findEdge(edge)

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

func (this *Beacon3d) getOrientationDiff(alt *Beacon3d) (signs types.Vector3d, order [3]int) {
	for _, arr := range this.edges {
		for _, edge := range arr {
			edges, ok := alt.findEdge(edge)

			if ok && len(edges) == 1 {
				// this edge has all the same absolute numbers in some orientation
				altEdge := edges[0]
				reOrderedEdge := types.Vector3d{}

				axes := [3]int{edge.X, edge.Y, edge.Z}
				altAxes := [3]int{altEdge.X, altEdge.Y, altEdge.Z}

				for i, axis := range axes {
					for j, altAxis := range altAxes {
						if utils.Abs(axis) == utils.Abs(altAxis) {
							// reorder y,z,x to x,y,z for example
							order[i] = j

							switch i {
							case 0:
								reOrderedEdge.X = altAxis
							case 1:
								reOrderedEdge.Y = altAxis
							case 2:
								reOrderedEdge.Z = altAxis
							}
							break
						}
					}
				}

				// signs is what to multiply origin vector by
				// to get the same signs:
				// {1,2,3} -> {-1,2,-3} = {-1,1,-1}
				signs = edge.Divide(reOrderedEdge)

				// TODO: lazy, grabbing just the first?
				return
			}
		}
	}

	return
}

func align(vec types.Vector3d, order [3]int, signs types.Vector3d) types.Vector3d {
	arr := [3]int{vec.X, vec.Y, vec.Z}

	vec = types.NewVector3d(
		arr[order[0]],
		arr[order[1]],
		arr[order[2]],
	)

	vec = vec.Multiply(signs)

	return vec
}

// returns new Beacons in the correct projection if enough matched
func (this *Scanner) CompareScanner(scanner *Scanner) (
	newBeacons []*Beacon3d,
	shared int,
	scannerPosition types.Vector3d,
) {
	// make a copy to alter the list within the loop
	unmatched := make([]*Beacon3d, len(scanner.Beacons))
	copy(unmatched, scanner.Beacons)
	remaining := len(this.Beacons)

	var selfBeacon, altBeacon *Beacon3d

	// compare all scanners' beacons' edges
	for _, a := range this.Beacons {
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
		signs, order := selfBeacon.getOrientationDiff(altBeacon)
		alt := align(altBeacon.position, order, signs)

		// need to transform unmatched
		for _, b := range unmatched {
			// compare unmatched to an edge in the same orientation
			aligned := align(b.position, order, signs)
			diff := aligned.Subtract(alt)
			newPosition := selfBeacon.position.Add(diff)

			newBeacons = append(newBeacons, &Beacon3d{
				position: newPosition,
			})
		}

		scannerPosition = selfBeacon.position.Subtract(alt)
	}

	return
}

func addEdge(a *Beacon3d, b *Beacon3d) {
	edge := a.position.Subtract(b.position)

	key := getVectorKey(edge)

	if a.edges == nil {
		// TODO: still no idea when this happens
		a.edges = map[string][]types.Vector3d{}
	}

	a.edges[key] = append(a.edges[key], edge)
}

func (scanner *Scanner) AddBeacons(beacons []*Beacon3d) {
	scanner.Beacons = append(scanner.Beacons, beacons...)

	// update edges just for new beacons
	for _, a := range beacons {
		a.edges = map[string][]types.Vector3d{}
		for _, b := range scanner.Beacons {
			if a == b {
				continue
			}
			addEdge(a, b)
			addEdge(b, a)
		}
	}
}

func MergeScanners(content []string) (
	composite *Scanner,
	relativePositions []types.Vector3d,
	err error,
) {
	scanners := ParseScanners(content)
	relativePositions = make([]types.Vector3d, 0, len(scanners))

	composite = scanners[0]

	// scanner 0 is at 0,0,0, relatively
	relativePositions = append(relativePositions, types.NewVector3d(0, 0, 0))

	queue := types.Queue[Scanner]{}

	for _, scanner := range scanners[1:] {
		queue.Push(scanner)
	}

	lastScanner := composite

	for len(queue) > 0 {
		scanner := queue.Shift()

		if scanner == lastScanner {
			err = errors.New(fmt.Sprint("repeat scanner found beacons:", len(scanner.Beacons)))
			return
		}

		lastScanner = scanner

		newBeacons, count, relativeScanner := composite.CompareScanner(scanner)

		if count > 0 {
			composite.AddBeacons(newBeacons)
			relativePositions = append(relativePositions, relativeScanner)
		} else {
			queue.Push(scanner)
		}
	}

	return
}

func ManhattanDistance(a types.Vector3d, b types.Vector3d) int {
	x := a.X - b.X
	y := a.Y - b.Y
	z := a.Z - b.Z

	return utils.Abs(x) + utils.Abs(y) + utils.Abs(z)
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
