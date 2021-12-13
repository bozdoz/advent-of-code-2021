package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

type heightmap [][]int

func newHeightMap(data []string) (heights heightmap) {
	for i, line := range data {
		heights = append(heights, []int{})
		row := strings.Split(line, "")

		for _, val := range row {
			rowint, err := strconv.Atoi(val)

			if err != nil {
				panic("could not create heightmap")
			}

			heights[i] = append(heights[i], rowint)
		}
	}

	return
}

func (heights *heightmap) findNeighbours(row, col int) (neighbours [][2]int, err error) {
	indices := [][]int{
		{col, row - 1},
		{col, row + 1},
		{col - 1, row},
		{col + 1, row},
	}

	rowmax := len(*heights) - 1
	colmax := len((*heights)[0]) - 1

	if row > rowmax || col > colmax {
		return neighbours, errors.New("col and row out of range")
	}

	for _, coords := range indices {
		c, r := coords[0], coords[1]

		if c < 0 || r < 0 || c > colmax || r > rowmax {
			continue
		}

		neighbours = append(neighbours, [2]int{r, c})
	}

	return
}

func (heights *heightmap) findNeighbouringValues(row, col int) (vals []int, err error) {
	neighbours, err := heights.findNeighbours(row, col)

	for _, coords := range neighbours {
		r, c := coords[0], coords[1]

		vals = append(vals, (*heights)[r][c])
	}

	return
}

func MinInt(nums ...int) int {
	min := nums[0]

	for _, val := range nums {
		if val < min {
			min = val
		}
	}

	return min
}

func (heights *heightmap) getLowPoints() (lowpoints [][2]int, err error) {
	for r, col := range *heights {
		for c, val := range col {
			neighbours, err := heights.findNeighbouringValues(r, c)

			if err != nil {
				return lowpoints, err
			}

			lowest := MinInt(neighbours...)

			if lowest > val {
				lowpoints = append(lowpoints, [2]int{r, c})
			}
		}
	}

	return
}

// find the low points
func PartOne(content []string) (output int, err error) {
	heights := newHeightMap(content)
	lowpoints, err := heights.getLowPoints()
	lowpointvals := []int{}

	for _, point := range lowpoints {
		r, c := point[0], point[1]
		lowpointvals = append(lowpointvals, heights[r][c])
	}

	// sum of 1 + height of each lowpoint
	output = utils.Sum(lowpointvals...) + len(lowpointvals)

	return
}

type basin struct {
	// row, col
	included   map[int]map[int]bool
	heights    heightmap
	rows, cols int
}

func (b *basin) isIncluded(r, c int) bool {
	return b.included[r] != nil && b.included[r][c]
}

func (b *basin) search(r, c int) {
	// we only search cells we know are within the basin
	if b.included[r] == nil {
		b.included[r] = make(map[int]bool, b.cols)
	}

	b.included[r][c] = true

	neighbours, err := b.heights.findNeighbours(r, c)

	if err != nil {
		panic(fmt.Sprintf("failed to find neighbours for %d, %d", r, c))
	}

	for _, coords := range neighbours {
		r, c = coords[0], coords[1]
		included := b.isIncluded(r, c)

		if included || b.heights[r][c] == 9 {
			continue
		}

		// search again!
		b.search(r, c)
	}
}

func (heights *heightmap) newBasin(lowpoint [2]int) (b basin) {
	b.heights = *heights
	b.rows = len((*heights)[0])
	b.cols = len(*heights)
	b.included = make(map[int]map[int]bool, b.rows)

	b.search(lowpoint[0], lowpoint[1])

	return
}

func (b *basin) getSize() (size int) {
	for _, vals := range b.included {
		size += len(vals)
	}

	return
}

// basins stem from lowpoints and encompass all cumulative neighbours
// until a neighbour is a 9
func (heights *heightmap) getBasinSizes() (sizes []int, err error) {
	lowpoints, err := heights.getLowPoints()

	if err != nil {
		return sizes, err
	}

	for _, coords := range lowpoints {
		b := heights.newBasin(coords)
		sizes = append(sizes, b.getSize())
	}

	return
}

// Find the three largest basins and multiply their sizes together
func PartTwo(content []string) (output int, err error) {
	h := newHeightMap(content)

	sizes, err := h.getBasinSizes()

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// top three
	output = sizes[0] * sizes[1] * sizes[2]

	return output, err
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
