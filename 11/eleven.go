package main

import (
	"fmt"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadFileAsLines

const (
	size      = 10
	maxIndex  = 9
	maxEnergy = 9
)

type Cell struct {
	energy     int
	flashed    bool
	neighbours []*Cell
}

type Grid struct {
	cells [size][size]*Cell
}

// custom string representation
func (grid *Grid) String() (output string) {
	output += "[[ "
	for r, cells := range grid.cells {
		for _, cell := range cells {
			output += cell.String()
		}

		if r < size-1 {
			output += "\n   "
		}
	}
	output += " ]]"

	return
}

// custom string representation
func (cell *Cell) String() string {
	return fmt.Sprint(cell.energy)
}

// this is the constructor-like function we are using
func newGridPointer(data []string) *Grid {
	grid := &Grid{}

	for r, line := range data {
		chars := strings.Split(line, "")
		for c, char := range chars {
			num, err := strconv.Atoi(char)

			if err != nil {
				panic("could not convert char to num in grid")
			}

			grid.cells[r][c] = &Cell{energy: num}
		}
	}

	grid.updateNeighbours()

	return grid
}

// IGNORE: this was added just for benchmarking (see BLOG.md)
func newGrid(data []string) (grid Grid) {
	for r, line := range data {
		chars := strings.Split(line, "")
		for c, char := range chars {
			num, err := strconv.Atoi(char)

			if err != nil {
				panic("could not convert char to num in grid")
			}

			grid.cells[r][c] = &Cell{energy: num}
		}
	}

	grid.updateNeighbours()

	return grid
}

func (grid *Grid) updateNeighbours() {
	for r, row := range grid.cells {
		for c, cell := range row {
			indices := [][]int{
				{r - 1, c - 1},
				{r - 1, c},
				{r - 1, c + 1},
				{r, c - 1},
				{r, c + 1},
				{r + 1, c - 1},
				{r + 1, c},
				{r + 1, c + 1},
			}

			for _, coords := range indices {
				r, c := coords[0], coords[1]

				if r < 0 || c < 0 || r > maxIndex || c > maxIndex {
					continue
				}

				cell.neighbours = append(cell.neighbours, grid.cells[r][c])
			}
		}
	}
}

// This increases the energy level of all adjacent octopuses by 1
// If this causes an octopus to have an energy level greater than 9, it also flashes.
func (cell *Cell) flash() {
	cell.flashed = true

	for _, neighbour := range cell.neighbours {
		neighbour.energy++
		if neighbour.energy > maxEnergy && !neighbour.flashed {
			neighbour.flash()
		}
	}
}

// any octopus that flashed during has its
// energy level set to 0
func (cell *Cell) resolveFlash() {
	cell.flashed = false
	cell.energy = 0
}

func (grid *Grid) getAllCells() (cells []*Cell) {
	for _, row := range grid.cells {
		for _, cell := range row {
			cells = append(cells, cell)
		}
	}

	return
}

func (grid *Grid) update() (flashes int) {
	// this is making me second guess my 2d array
	cells := grid.getAllCells()

	// First, the energy level of each octopus increases by 1.
	for _, cell := range cells {
		cell.energy++
	}

	// any octopus with an energy level greater than 9 flashes
	for _, cell := range cells {
		// check for cell.flashed because
		// it may have been flashed by a neighbour
		// ...scandalous!
		if cell.energy > maxEnergy && !cell.flashed {
			cell.flash()
		}
	}

	// wanted to make this a deferred statement, but no
	for _, cell := range cells {
		if cell.flashed {
			flashes++
			cell.resolveFlash()
		}
	}

	return
}

// How many total flashes are there after 100 steps?
func PartOne(content []string) (output int, err error) {
	grid := newGridPointer(content)

	for i := 0; i < 100; i++ {
		output += grid.update()
	}

	return
}

func PartTwo(content []string) (output int, err error) {
	grid := newGridPointer(content)

	// set an upper bound of 1K?
	for i := 0; i < 1000; i++ {
		flashes := grid.update()

		if flashes == size*size {
			// return value is not 0-indexed
			return i + 1, nil
		}
	}

	return
}

func PartTwoValue(content []string) (output int, err error) {
	grid := newGrid(content)

	// set an upper bound of 1K?
	for i := 0; i < 1000; i++ {
		flashes := grid.update()

		if flashes == size*size {
			// return value is not 0-indexed
			return i + 1, nil
		}
	}

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
