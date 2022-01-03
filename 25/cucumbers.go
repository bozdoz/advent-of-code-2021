package main

type Cucumber rune

const (
	EMPTY Cucumber = '.'
	LEFT  Cucumber = '>'
	DOWN  Cucumber = 'v'
)

type Grid struct {
	cucumbers     [][]Cucumber
	width, height int
}

func parseInput(data []string) *Grid {
	height := len(data)
	width := len(data[0])

	grid := &Grid{
		width:  width,
		height: height,
	}

	grid.cucumbers = make([][]Cucumber, height)

	for i, line := range data {
		grid.cucumbers[i] = make([]Cucumber, width)
		for j, char := range line {
			grid.cucumbers[i][j] = Cucumber(char)
		}
	}

	return grid
}

// utility func to wrap around grid
func (grid *Grid) get(i, j int) Cucumber {
	return grid.cucumbers[i%grid.height][j%grid.width]
}

// utility func to wrap around grid
func (grid *Grid) set(i, j int, cucumber Cucumber) {
	grid.cucumbers[i%grid.height][j%grid.width] = cucumber
}

func (grid *Grid) isEmpty(i, j int) bool {
	return grid.get(i, j) == EMPTY
}

func (grid *Grid) moveLeft(i, j int) {
	grid.set(i, j+1, grid.get(i, j))
	grid.set(i, j, EMPTY)
}

func (grid *Grid) moveDown(i, j int) {
	grid.set(i+1, j, grid.get(i, j))
	grid.set(i, j, EMPTY)
}

func (grid *Grid) step() (moves int) {
	for i := 0; i < grid.height; i++ {
		// need to ignore wrap around cucumbers
		max := grid.width
		// if the wrap around has cucumbers of same type,
		// then alter `max` to ignore them
		if grid.cucumbers[i][0] == LEFT {
			for k := max - 1; k > 0; k-- {
				if grid.cucumbers[i][k] == LEFT {
					// avoid iterating this cucumber again
					max--
				} else {
					break
				}
			}
		}

		for j := 0; j < max; j++ {
			if grid.cucumbers[i][j] == LEFT && grid.isEmpty(i, j+1) {
				grid.moveLeft(i, j)
				// skip next iteration
				j++
				moves++
			}
		}
	}

	// iterate width, then height for the down-trodden
	for j := 0; j < grid.width; j++ {
		// need to ignore wrap around cucumbers (see if i == 0 below)
		max := grid.height
		// if the wrap around has cucumbers of same type,
		// then alter `max` to ignore them
		if grid.cucumbers[0][j] == DOWN {
			for k := max - 1; k > 0; k-- {
				if grid.cucumbers[k][j] == DOWN {
					// avoid iterating this cucumber again
					max--
				} else {
					break
				}
			}
		}

		for i := 0; i < max; i++ {
			if grid.cucumbers[i][j] == DOWN && grid.isEmpty(i+1, j) {
				grid.moveDown(i, j)
				// skip next iteration
				i++
				moves++
			}
		}
	}

	return
}

func (grid *Grid) stepsTilStopped() (steps int) {
	for {
		moves := grid.step()

		// example says stops after 58 steps, but clearly stops after
		// 57 steps; always incrementing steps here before
		// the moves == 0 check for that reason
		steps++

		if moves == 0 {
			break
		}
	}

	return
}

//
// string reps
//

func (grid *Grid) String() (output string) {
	for _, row := range grid.cucumbers {
		output += "\n"
		for _, val := range row {
			output += string(val)
		}
	}

	return
}
