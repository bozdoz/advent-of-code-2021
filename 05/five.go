package main

import (
	"fmt"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

type coords [2]int

type line struct {
	from, to coords
}

type grid struct {
	// 1d representation of 2d space
	space         []int
	width         int
	lines         []line
	checkDiagonal bool
}

func (g *grid) load(data []string) (err error) {
	maxX := 0
	maxY := 0

	for _, row := range data {
		var x1, y1, x2, y2 int

		_, err := fmt.Fscanf(strings.NewReader(row), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if err != nil {
			return err
		}

		g.lines = append(g.lines, line{
			from: [2]int{x1, y1},
			to:   [2]int{x2, y2},
		})

		if x1 > maxX {
			maxX = x1
		}

		if x2 > maxX {
			maxX = x2
		}

		if x1 > maxX {
			maxX = x1
		}

		if y1 > maxY {
			maxY = y1
		}

		if y2 > maxY {
			maxY = y2
		}
	}

	g.width = maxX + 1
	g.space = make([]int, g.width*(maxY+1))

	return
}

func (g *grid) draw(x, y int) {
	g.space[g.width*y+x]++
}

func (g *grid) drawVertical(l line) {
	x := l.from[0]
	y := l.from[1]
	y2 := l.to[1]

	// beginning coord should be counted
	g.draw(x, y)

	for y != y2 {
		if y < y2 {
			y++
		} else {
			y--
		}
		g.draw(x, y)
	}
}

func (g *grid) drawHorizontal(l line) {
	y := l.from[1]
	x := l.from[0]
	x2 := l.to[0]

	// beginning coord should be counted
	g.draw(x, y)

	for x != x2 {
		if x < x2 {
			x++
		} else {
			x--
		}
		g.draw(x, y)
	}
}

func (g *grid) drawDiagonal(l line) {
	y := l.from[1]
	x := l.from[0]
	x2 := l.to[0]
	y2 := l.to[1]

	// beginning coord should be counted
	g.draw(x, y)

	for x != x2 {
		if x < x2 {
			x++
		} else {
			x--
		}

		if y < y2 {
			y++
		} else {
			y--
		}

		g.draw(x, y)
	}
}

func (g *grid) drawLines() {
	for _, l := range g.lines {
		// only draw horizontal/vertical lines
		if l.from[0] == l.to[0] {
			g.drawVertical(l)
		} else if l.from[1] == l.to[1] {
			g.drawHorizontal(l)
		} else if g.checkDiagonal {
			g.drawDiagonal(l)
		}
	}
}

func PartOne(content []string) (output int, err error) {
	grid := grid{}

	grid.load(content)
	grid.drawLines()

	for _, num := range grid.space {
		if num > 1 {
			output++
		}
	}

	return
}

func PartTwo(content []string) (output int, err error) {
	grid := grid{
		checkDiagonal: true,
	}

	grid.load(content)
	grid.drawLines()

	for _, num := range grid.space {
		if num > 1 {
			output++
		}
	}

	return
}

func main() {
	// safe to assume
	filename := "input.txt"

	vals := utils.LoadFile(filename)

	answer, err := PartOne(vals)

	if err != nil {
		fmt.Println("failed to parse PartOne", err)
		return
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(vals)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
