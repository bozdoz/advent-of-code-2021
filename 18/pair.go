package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/bozdoz/advent-of-code-2021/types"
)

type Element struct {
	value int
	pair  *Pair
}

type Pair struct {
	left, right *Element
	parent      *Pair
}

func (pair *Pair) append(something any) {
	element := &Element{}

	// "something" is either value or a pair
	switch v := something.(type) {
	case int:
		element.value = v
	case *Pair:
		element.pair = v
		// also link pair as its parent
		v.parent = pair
	default:
		panic(fmt.Sprint("Not sure what this is: ", something))
	}

	// appends left first, then right, then panics
	if pair.left == nil {
		pair.left = element
	} else if pair.right == nil {
		pair.right = element
	} else {
		panic(fmt.Sprint("left and right on this pair are both full: ", pair))
	}
}

func parsePairs(data string) *Pair {
	dec := json.NewDecoder(strings.NewReader(data))
	// keep track of nested pairs
	stack := &types.Stack[Pair]{}

	var cur *Pair

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		switch t {
		case json.Delim('['):
			// start a new pair
			cur = &Pair{}
			// add it to the stack: LIFO
			stack.Push(cur)
		case json.Delim(']'):
			// signals this pair is complete
			lastPair := stack.Pop()
			// previous pair becomes current pair
			cur = stack.Peek()

			if cur != nil {
				// last pair is appended to current pair
				cur.append(lastPair)
			} else {
				// last pair should be returned?
				return lastPair
			}
		default:
			// always is a float64
			numeric, ok := t.(float64)

			if !ok {
				given := fmt.Sprintf("%T %v", t, t)
				panic(fmt.Sprint("why is this not a float: ", given))
			}

			cur.append(int(numeric))
		}
	}

	// should be returned in the switch case above
	return nil
}

/*
To reduce a snailfish number, you must repeatedly do the first action in this list that applies to the snailfish number:

1. If any pair is nested inside four pairs, the leftmost such pair explodes.

2. If any regular number is 10 or greater, the leftmost such regular number splits.
*/
func (this *Pair) reduce() {
	// 1. should I explode?
	explodeMe := this.getNestedPairAtDepth(4)

	if explodeMe != nil {
		// do as he says
		explodeMe.explode()

		log.Println("After explode:", this)

		// start again if something changed
		this.reduce()
	} else {
		// value gte 10
		splitMe := this.getPairWithValueGTE(splitVal)

		if splitMe != nil {
			splitMe.split(splitVal)

			log.Println("After split:  ", this)

			// start again if something changed
			this.reduce()
		}
	}
}

func (this *Pair) isNthChild(n int) bool {
	// check if has n number of parents
	cur := this

	for i := 0; cur != nil; i++ {
		cur = cur.parent

		if i == n {
			return true
		}
	}

	return false
}

// finds the first/left-most pair at a given depth
func (this *Pair) getNestedPairAtDepth(depth int) *Pair {
	queue := types.Queue[Pair]{this}

	for len(queue) > 0 {
		cur := queue.Shift()

		if cur == nil {
			continue
		}

		if cur.isNthChild(depth) {
			return cur
		}

		// check nested pairs
		queue.Push(cur.left.pair)
		queue.Push(cur.right.pair)
	}

	return nil
}

func (pair *Pair) explode() {
	// "Exploding pairs will always consist of two regular numbers."
	left := pair.left
	right := pair.right

	elements := pair.getRootElements()

	// "the pair's left value is added to the first regular number to the left of the exploding pair (if any)"
	for i := range elements {
		if elements[i] == left && i > 0 {
			// get node to the left
			elements[i-1].value += left.value
		} else if elements[i] == right && i < len(elements)-1 {
			// get node to the right
			elements[i+1].value += right.value
		}
	}

	// current pair is "replaced with the regular number 0"
	pair.replaceWithInt(0)
}

func (this *Pair) replaceWithInt(val int) {
	parent := this.parent
	// need too figure out if this the left or right element
	if this == parent.left.pair {
		parent.left.pair = nil
		parent.left.value = val
	} else {
		parent.right.pair = nil
		parent.right.value = val
	}

	this = nil
}

// TODO: this seems expensive
func (pair *Pair) getRootElements() []*Element {
	// get root
	root := pair
	for root.parent != nil {
		root = root.parent
	}

	// get all elements
	elements := root.traverseElements()

	return elements
}

func (pair *Pair) traverseElements() []*Element {
	elements := []*Element{}

	if pair.left.pair != nil {
		elements = pair.left.pair.traverseElements()
	} else {
		elements = append(elements, pair.left)
	}

	if pair.right.pair != nil {
		elements = append(elements, pair.right.pair.traverseElements()...)
	} else {
		elements = append(elements, pair.right)
	}

	return elements
}

func (this *Pair) getFirstPairFunc(fnc func(candidate *Element) bool) (ret *Pair) {

	if fnc(this.left) {
		return this
	}

	if this.left.pair != nil {
		ret = this.left.pair.getFirstPairFunc(fnc)
		if ret != nil {
			return ret
		}
	}

	if fnc(this.right) {
		return this
	}

	if this.right.pair != nil {
		ret = this.right.pair.getFirstPairFunc(fnc)
		if ret != nil {
			return ret
		}
	}

	return nil
}

var splitVal = 10

func (this *Pair) getPairWithValueGTE(val int) *Pair {
	// find left-most val 10 or greater
	return this.getFirstPairFunc(func(candidate *Element) bool {
		return candidate.value >= val
	})
}

func (parent *Pair) split(splitIf int) {
	// the left element of the pair should be the regular number
	// divided by two and rounded down, while the right element
	// of the pair should be the regular number divided by two and
	// rounded up

	// find which side has the value
	side := parent.right

	if parent.left.value >= splitIf {
		side = parent.left
	}

	val := side.value
	remainder := val % 2
	left := &Element{value: val / 2}
	right := &Element{value: val/2 + remainder}

	replacement := &Pair{
		left,
		right,
		parent,
	}

	// replace value with pair
	side.value = 0
	side.pair = replacement
}

func (this *Pair) add(pair *Pair) *Pair {
	newPair := &Pair{}
	this.parent = newPair
	pair.parent = newPair

	newPair.left = &Element{pair: this}
	newPair.right = &Element{pair: pair}

	newPair.reduce()

	return newPair
}

func (this *Element) getMagnitude() int {
	if this.pair != nil {
		return this.pair.getMagnitude()
	}

	return this.value
}

func (this *Pair) getMagnitude() int {
	return this.left.getMagnitude()*3 + this.right.getMagnitude()*2
}

//
// String representations
//

// string representation to prove I'm parsing it correctly
func (element *Element) String() (output string) {
	if element.pair == nil {
		return fmt.Sprint(element.value)
	}

	return element.pair.String()
}

// string representation to prove I'm parsing it correctly
func (pair *Pair) String() (output string) {
	output += "["
	output += pair.left.String()
	output += ","
	output += pair.right.String()
	output += "]"

	return
}
