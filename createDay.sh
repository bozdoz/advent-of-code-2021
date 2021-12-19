#!/bin/bash

NEW_DAY=$1
NEW_DAY_NAME=$2

if [ -z $NEW_DAY ]; then
  echo "Provide ## for new day directory"
  exit 1
fi

if [ -z $NEW_DAY_NAME ]; then
  echo "Provide go filename for new day: one, two, three"
  exit 1
fi

mkdir $NEW_DAY

cd $NEW_DAY

# start touching things
touch README.md
touch input.txt
touch example.txt

# create main go file
cat > $NEW_DAY_NAME.go << EOF
package main

import (
	"fmt"
	"io/ioutil"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString
// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content string) (output int, err error) {
	return
}

func PartTwo(content string) (output int, err error) {
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
EOF

# create test file
cat > ${NEW_DAY_NAME}_test.go << EOF
package main

import (
	"os"
	"testing"
)

// fill in the answers for each part (as they come)
var answers = map[int]int{
	1: 0,
	2: 0,
}

var vals = FileLoader("example.txt")

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestExampleOne(t *testing.T) {
	expected, ok := answers[1]

	if !ok {
		return
	}

	val, err := PartOne(vals)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}

func TestExampleTwo(t *testing.T) {
	expected, ok := answers[2]

	if !ok {
		return
	}

	val, err := PartTwo(vals)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if val != expected {
		t.Logf("Answer should be %v, but got %v", expected, val)
		t.Fail()
	}
}
EOF