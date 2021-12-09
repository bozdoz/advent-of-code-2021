#!/bin/bash

NEW_DAY=$1
NEW_DAY_NAME=$2
TEMPLATE=${3:-'09'}

if [ -z $NEW_DAY ]; then
  echo "Provide ## for new day directory"
  exit 1
fi

if [ -z $NEW_DAY_NAME ]; then
  echo "Provide go filename for new day: one, two, three"
  exit 1
fi

cp -r $TEMPLATE $NEW_DAY

cd $NEW_DAY

# get name from old day
for f in *.go; do
  if [ $f != "*_test*" ]; then
    OLD_DAY_NAME=$(echo "${f/.go/}")
    break
  fi
done;

# rename all go files from old name to new name
for f in *.go; do
  mv $f ${f/${OLD_DAY_NAME}/${NEW_DAY_NAME}}
done;

# overwrite main go file
cat > $NEW_DAY_NAME.go << EOF
package main

import (
	"fmt"

	"bozdoz.com/aoc-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsString

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