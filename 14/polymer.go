package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

// stores all the puzzle inputs
type Polymer struct {
	// CBKH
	template string
	// CB -> H
	insertionRules map[string]string
}

func newPolymer(data string) *Polymer {
	polymer := &Polymer{
		insertionRules: map[string]string{},
	}

	parts := utils.SplitByEmptyNewline(data)

	polymer.template = parts[0]

	for _, instruction := range strings.Split(parts[1], "\n") {
		var key, val string

		// for griff
		if count, err := fmt.Sscanf(instruction, "%2s -> %1s", &key, &val); count != 2 || err != nil {
			if err != nil {
				log.Println("failed to parse instruction", instruction)
				log.Println(err)
				// maybe EOF?
				continue
			}
			panic("could not initialize all values of " + instruction)
		}

		polymer.insertionRules[key] = val
	}

	return polymer
}

type Elements struct {
	pairs, charCount map[string]int
}

func (elements *Elements) merge(otherElements Elements) {
	for key, val := range otherElements.pairs {
		elements.pairs[key] += val
	}

	for key, val := range otherElements.charCount {
		elements.charCount[key] += val
	}
}

func newElements(template string) Elements {
	elements := Elements{
		pairs:     map[string]int{},
		charCount: map[string]int{},
	}

	for i := 0; i < len(template); i++ {
		if i < len(template)-1 {
			pair := template[i : i+2]
			elements.pairs[pair]++
		}

		char := template[i : i+1]
		elements.charCount[char]++
	}

	return elements
}

func (polymer *Polymer) pairInsertion(steps int) Elements {
	elements := newElements(polymer.template)

	// decrement steps
	for ; steps > 0; steps-- {
		// get empty Elements
		nextElements := newElements("")

		// ? iterating each instruction, but maybe it would be
		// better to iterate each elements.pairs instead
		// (since elements.pairs is likely to be smaller)
		for ref, replacement := range polymer.insertionRules {
			// could check count > 0, but why waste a decent "ok"?
			count, ok := elements.pairs[ref]

			if !ok {
				// it's ok not to be ok
				continue
			}

			// increment char count for replacement
			nextElements.charCount[replacement] += count

			// the replacement creates two new pairs
			// NN -> C creates NC & CN
			// * NOTE: ref[0] is a byte, ref[0:1] is a string
			firstNewPair := ref[0:1] + replacement
			secondNewPair := replacement + ref[1:2]

			nextElements.pairs[firstNewPair] += count
			nextElements.pairs[secondNewPair] += count

			// remove pair from pairs so that we can merge
			// any unmatched pairs (they should persist)
			delete(elements.pairs, ref)
		}

		// merge any unmatched pairs
		elements.merge(nextElements)
	}

	return elements
}

func (elements *Elements) getMinMax() (int, int) {
	// first time using Inf
	max := math.Inf(-1)
	min := math.Inf(1)

	for _, count := range elements.charCount {
		floated := float64(count)
		if floated > max {
			max = floated
		}
		if floated < min {
			min = floated
		}
	}

	return int(min), int(max)
}
