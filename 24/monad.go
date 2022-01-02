package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction interface {
	Exec(a *int, b int)
	String() string
}

type Command struct {
	left, right string
}

type Inp struct {
	Command
}

func (inp *Inp) Exec(a *int, b int) {
	*a = b
}

type Add struct {
	Command
}

func (add *Add) Exec(a *int, b int) {
	*a += b
}

type Mul struct {
	Command
}

func (mul *Mul) Exec(a *int, b int) {
	*a *= b
}

type Div struct {
	Command
}

func (div *Div) Exec(a *int, b int) {
	*a /= b
}

type Mod struct {
	Command
}

func (mod *Mod) Exec(a *int, b int) {
	*a %= b
}

type Eql struct {
	Command
}

func (eql *Eql) Exec(a *int, b int) {
	if *a == b {
		*a = 1
	} else {
		*a = 0
	}
}

type Program struct {
	w, x, y, z   int
	instructions []Instruction
	states       [][4]int
	blocks       *[14][]Instruction
}

func parseInput(data []string) *Program {
	program := &Program{
		instructions: make([]Instruction, len(data)),
	}

	re := regexp.MustCompile(`^(\w{3})\s(\w)\s?(-?\w*?)$`)

	for i, line := range data {
		parts := re.FindStringSubmatch(line)

		if len(parts) == 0 {
			log.Println(i, line)
		}

		cmd, a, b := parts[1], parts[2], parts[3]

		switch cmd {
		case "inp":
			program.instructions[i] = &Inp{Command{
				left:  a,
				right: b,
			}}
		case "add":
			program.instructions[i] = &Add{Command{
				left:  a,
				right: b,
			}}
		case "mul":
			program.instructions[i] = &Mul{Command{
				left:  a,
				right: b,
			}}
		case "div":
			program.instructions[i] = &Div{Command{
				left:  a,
				right: b,
			}}
		case "mod":
			program.instructions[i] = &Mod{Command{
				left:  a,
				right: b,
			}}
		case "eql":
			program.instructions[i] = &Eql{Command{
				left:  a,
				right: b,
			}}
		default:
			panic("what happened")
		}
	}

	program.updateBlocks()

	return program
}

func (program *Program) saveState() {
	program.states = append(program.states, [4]int{
		program.w,
		program.x,
		program.y,
		program.z,
	})
}

func (program *Program) restoreState() {
	n := len(program.states) - 1
	last := (program.states)[n]

	program.states = (program.states)[:n]

	program.w = last[0]
	program.x = last[1]
	program.y = last[2]
	program.z = last[3]
}

func (program *Program) doCommand(cmd Instruction, input int) {
	var a *int
	var b int
	vars := map[string]*int{
		"w": &program.w,
		"x": &program.x,
		"y": &program.y,
		"z": &program.z,
	}

	exec := func(v Command) {
		var err error

		a = vars[v.left]

		switch {
		case strings.ContainsAny(v.right, "wxyz"):
			b = *vars[v.right]
		default:
			b, err = strconv.Atoi(v.right)
			if err != nil {
				panic(fmt.Sprint("can't parse v.right", v.right))
			}
		}

		cmd.Exec(a, b)
	}

	// TODO: a little painful
	switch v := cmd.(type) {
	case *Inp:
		a := vars[v.left]
		cmd.Exec(a, input)
	case *Add:
		exec(v.Command)
	case *Mul:
		exec(v.Command)
	case *Div:
		exec(v.Command)
	case *Mod:
		exec(v.Command)
	case *Eql:
		exec(v.Command)
	}
}

func (program *Program) reset() {
	program.w = 0
	program.x = 0
	program.y = 0
	program.z = 0
}

func (program *Program) input(ints ...int) {
	var i, j, input int

	for i = 0; i < len(program.instructions); i++ {
		inst := program.instructions[i]

		_, isInput := inst.(*Inp)

		if isInput {
			// get next input
			input = ints[j]
			j++
		}

		program.doCommand(inst, input)
	}
}

var Largest = [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

var Smallest = [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

func (program *Program) updateBlocks() {
	inputBlocks := [14][]Instruction{}

	var i, j int

	// first instruction is input, so start by incrementing to 0
	j = -1

	for i = 0; i < len(program.instructions); i++ {
		_, isInput := program.instructions[i].(*Inp)

		if isInput {
			j++
		}

		inputBlocks[j] = append(inputBlocks[j], program.instructions[i])
	}

	program.blocks = &inputBlocks
}

func (program *Program) decrement(i int) (solved bool) {
	block := program.blocks[i]

	z := program.z

	if failed[i] != nil && failed[i][z] {
		return false
	}

	program.saveState()
	for j := 9; j > 0; j-- {
		program.z = z
		for _, inst := range block {
			program.doCommand(inst, j)
		}
		if i < 13 && program.decrement(i+1) {
			Largest[i] = j
			return true
		} else {
			// finished (2.4s for 1M iterations)
			if program.z == 0 {
				Largest[i] = j
				return true
			}
		}
	}
	program.restoreState()

	// cache the big ones
	if i < 10 {
		if failed[i] == nil {
			failed[i] = map[int]bool{}
		}
		// this iteration with this z value will never work
		failed[i][z] = true
	}

	return
}

var failed = map[int]map[int]bool{}

func (program *Program) decrementDirect(i int) (solved bool) {
	diffs := blockDiffs[i]

	// save state
	zPrev := program.z

	if failed[i] != nil && failed[i][zPrev] {
		return false
	}

	for j := 9; j > 0; j-- {
		program.z = block(j, zPrev, diffs[0], diffs[1], diffs[2])
		if i < 13 {
			solved = program.decrementDirect(i + 1)
			if solved {
				Largest[i] = j
				return true
			}
		} else {
			// finished (19ms for 1M iterations)
			if program.z == 0 {
				Largest[i] = j
				return true
			}
		}
	}

	// restore
	program.z = zPrev

	// cache the big ones
	if i < 10 {
		if failed[i] == nil {
			failed[i] = map[int]bool{}
		}
		// this iteration with this z value will never work
		failed[i][zPrev] = true
	}

	return
}

func (program *Program) solveLargest() [14]int {
	program.decrementDirect(0)

	return Largest
}

func (program *Program) incrementDirect(i int) (solved bool) {
	diffs := blockDiffs[i]

	// save state
	zPrev := program.z

	if failed[i] != nil && failed[i][zPrev] {
		return false
	}

	for j := 1; j < 10; j++ {
		program.z = block(j, zPrev, diffs[0], diffs[1], diffs[2])
		if i < 13 && program.incrementDirect(i+1) {
			Smallest[i] = j
			return true
		} else {
			// finished
			if program.z == 0 {
				Smallest[i] = j
				return true
			}
		}
	}

	// restore
	program.z = zPrev

	// cache the big ones
	if i < 10 {
		if failed[i] == nil {
			failed[i] = map[int]bool{}
		}
		// this iteration with this z value will never work
		failed[i][zPrev] = true
	}

	return
}

func (program *Program) solveSmallest() [14]int {
	program.incrementDirect(0)

	return Smallest
}

//
// String reps
//

func (program *Program) String() (output string) {
	output += fmt.Sprintf("w: %d x: %d y: %d z: %d\n", program.w, program.x, program.y, program.z)

	for _, instruction := range program.instructions {
		output += fmt.Sprintf("%s -> ", instruction)
	}

	return
}

func (this Inp) String() string {
	return fmt.Sprintf("inp %v", this.left)
}

func (this Add) String() string {
	return fmt.Sprintf("add %v %v", this.left, this.right)
}

func (this Mul) String() string {
	return fmt.Sprintf("mul %v %v", this.left, this.right)
}

func (this Div) String() string {
	return fmt.Sprintf("div %v %v", this.left, this.right)
}

func (this Mod) String() string {
	return fmt.Sprintf("mod %v %v", this.left, this.right)
}

func (this Eql) String() string {
	return fmt.Sprintf("eql %v %v", this.left, this.right)
}
