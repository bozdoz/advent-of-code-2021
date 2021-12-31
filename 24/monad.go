package main

import (
	"fmt"
	"regexp"
)

type Instruction interface {
	Exec(a, b int) int
	String() string
}

type Command struct {
	left, right string
}

type Inp Command

func (inp *Inp) Exec(a, b int) int {
	return a
}

type Add Command

func (add *Add) Exec(a, b int) int {
	return a + b
}

type Mul Command

func (mul *Mul) Exec(a, b int) int {
	return a * b
}

type Div Command

func (div *Div) Exec(a, b int) int {
	return a / b
}

type Mod Command

func (mod *Mod) Exec(a, b int) int {
	return a % b
}

type Eql Command

func (eql *Eql) Exec(a, b int) int {
	if a == b {
		return 1
	}

	return 0
}

type Program struct {
	w, x, y, z   int
	instructions []Instruction
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
			program.instructions[i] = &Inp{
				left:  a,
				right: b,
			}
		case "add":
			program.instructions[i] = &Add{
				left:  a,
				right: b,
			}
		case "mul":
			program.instructions[i] = &Mul{
				left:  a,
				right: b,
			}
		case "div":
			program.instructions[i] = &Div{
				left:  a,
				right: b,
			}
		case "mod":
			program.instructions[i] = &Mod{
				left:  a,
				right: b,
			}
		case "eql":
			program.instructions[i] = &Eql{
				left:  a,
				right: b,
			}
		default:
			panic("what happened")
		}
	}

	log.Println(program)

	return program
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
