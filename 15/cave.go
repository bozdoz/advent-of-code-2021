package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

type Cell struct {
	// risk factor
	value int
	// tentative distance (Dijkstra's algorithm)
	distance int
	// we're not coming back
	visited    bool
	neighbours []*Cell
}

type Row []*Cell
type Grid []Row

type Cave struct {
	grid          Grid
	height, width int
	start, end    *Cell
}

func newCave(data []string) (cave Cave) {
	cave.height = len(data)
	cave.width = len(data[0])
	// math.Inf is awful to work with
	startingDistance := cave.height * cave.width * 10

	// "the area you originally scanned is just one tile in a 5x5 tile area that forms the full map"
	startingDistance *= 5

	for r, line := range data {
		cave.grid = append(cave.grid, make(Row, cave.width))
		for c, char := range strings.Split(line, "") {
			value, err := strconv.Atoi(char)

			if err != nil {
				panic("could not convert char to int")
			}

			cave.grid[r][c] = &Cell{
				value:    value,
				distance: startingDistance,
			}
		}
	}

	cave.start = cave.grid[0][0]
	// start is 0 distance away from start
	cave.start.distance = 0
	cave.end = cave.grid[cave.height-1][cave.width-1]

	cave.updateNeighbours()

	return
}

// copied from eleven.go
func (cave *Cave) updateNeighbours() {
	maxRow := cave.height - 1
	maxCol := cave.width - 1

	for r, row := range cave.grid {
		for c, cell := range row {
			indices := [][]int{
				{r + 1, c},
				{r, c + 1},
				{r - 1, c},
				{r, c - 1},
			}

			for _, coords := range indices {
				r, c := coords[0], coords[1]

				if r < 0 || c < 0 || r > maxRow || c > maxCol {
					continue
				}

				cell.neighbours = append(cell.neighbours, cave.grid[r][c])
			}
		}
	}
}

func (cave *Cave) findAllPaths() {
	pq := make(PriorityQueue, cave.height*cave.width)

	for r, row := range cave.grid {
		for c, cell := range row {
			index := r*cave.width + c
			pq[index] = &Item{
				value:    cell,
				priority: cell.distance,
				index:    index,
			}
		}
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Item).value

		for i := range cell.neighbours {
			neighbour := cell.neighbours[i]

			if neighbour.visited {
				continue
			}

			// update distance of neighbours
			neighbour.distance = utils.MinInt(
				neighbour.value+cell.distance,
				neighbour.distance,
			)

			// hide the magic in the priority queue
			pq.update(neighbour)
		}

		cell.visited = true
	}
}

//
// -- String representations --
//

// custom string representation
func (cave *Cave) String() (output string) {
	output += "[[ "
	height := cave.height
	for r, cells := range cave.grid {
		for _, cell := range cells {
			output += fmt.Sprint(cell.value)
		}

		if r < height-1 {
			output += "\n   "
		}
	}
	output += " ]]"

	return
}

// custom string representation (fix recursion)
func (cell *Cell) String() string {
	return fmt.Sprint(cell.value)
}
