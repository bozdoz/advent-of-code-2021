package main

import (
	"fmt"
	"strings"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

type FoldInstruction struct {
	// "x" | "y"
	axis string
	line int
}

type Paper struct {
	dots             map[int]map[int]bool
	foldInstructions []FoldInstruction
}

func (paper *Paper) drawDot(x, y int) {
	if paper.dots[x] == nil {
		paper.dots[x] = map[int]bool{}
	}
	paper.dots[x][y] = true
}

func newPaper(data string) (paper Paper) {
	parts := utils.SplitByEmptyNewline(data)
	dotCoords, instructions := parts[0], parts[1]

	paper.dots = map[int]map[int]bool{}

	for _, coordStr := range strings.Split(dotCoords, "\n") {
		var x, y int

		fmt.Sscanf(coordStr, "%d,%d", &x, &y)

		paper.drawDot(x, y)
	}

	for _, instructionStr := range strings.Split(instructions, "\n") {
		var axis string
		var line int

		// lucky caught that this should be %1s
		fmt.Sscanf(instructionStr, "fold along %1s=%d", &axis, &line)

		paper.foldInstructions = append(paper.foldInstructions, FoldInstruction{
			axis, line,
		})
	}

	return
}

type Coords [2]int

func (paper *Paper) foldFunc(fn func(x, y int) Coords) {
	nextDots := []Coords{}

	for x, row := range paper.dots {
		for y := range row {
			nextDots = append(nextDots, fn(x, y))
		}
	}

	// brand new board
	paper.dots = map[int]map[int]bool{}

	for _, coords := range nextDots {
		x, y := coords[0], coords[1]

		paper.drawDot(x, y)
	}
}

func (paper *Paper) foldUp(lineNum int) {
	paper.foldFunc(func(x, y int) Coords {
		if y > lineNum {
			diff := y - lineNum
			y = lineNum - diff
		}

		return Coords{x, y}
	})
}

func (paper *Paper) foldLeft(lineNum int) {
	paper.foldFunc(func(x, y int) Coords {
		if x > lineNum {
			diff := x - lineNum
			x = lineNum - diff
		}

		return Coords{x, y}
	})
}

func (paper *Paper) fold(instruction FoldInstruction) {
	if instruction.axis == "y" {
		paper.foldUp(instruction.line)
	} else {
		paper.foldLeft(instruction.line)
	}
}

func (paper *Paper) countDots() (count int) {
	for _, row := range paper.dots {
		count += len(row)
	}

	return
}

// output a board similar to adventofcode.com/2021/day/13
func (paper *Paper) Board() (output string) {
	width := 0
	height := 0

	for x, row := range paper.dots {
		for y := range row {
			// 0-indexed
			x++
			y++

			if x > width {
				width = x
			}

			if y > height {
				height = y
			}
		}
	}

	board := make([][]string, height)

	for r := range board {
		board[r] = strings.Split(strings.Repeat(".", width), "")
	}

	for x, row := range paper.dots {
		for y := range row {
			board[y][x] = "#"
		}
	}

	output += "\n\n"

	for _, row := range board {
		output += strings.Join(row, "") + "\n"
	}

	return
}
