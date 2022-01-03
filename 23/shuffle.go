package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/bozdoz/advent-of-code-2021/types"
	"github.com/bozdoz/advent-of-code-2021/utils"
)

const (
	HALLWAY_SPOTS = 7
)

type AmphipodType byte

const (
	A AmphipodType = 'A'
	B AmphipodType = 'B'
	C AmphipodType = 'C'
	D AmphipodType = 'D'
)

var costs = map[AmphipodType]int{
	A: 1,
	B: 10,
	C: 100,
	D: 1000,
}

var rooms = map[AmphipodType]int8{
	A: 2,
	B: 4,
	C: 6,
	D: 8,
}

type Amphipod struct {
	x, y  int8
	_type AmphipodType
}

type Grid [5][11]*Amphipod

type Burrow struct {
	amphipods []*Amphipod
	cost      int
	grid      *Grid
	states    []string
}

func parseInput(data string) *Burrow {
	burrow := &Burrow{
		amphipods: make([]*Amphipod, 0, 16),
		grid:      &Grid{},
	}

	re := regexp.MustCompile("[ABCD]")

	for i, match := range re.FindAllString(data, 16) {
		_type := match[0]

		room := (i%4)*2 + 2
		position := (i / 4) + 1

		pod := &Amphipod{
			x:     int8(room),
			y:     int8(position),
			_type: AmphipodType(_type),
		}

		burrow.grid[pod.y][pod.x] = pod

		burrow.amphipods = append(burrow.amphipods, pod)
	}

	return burrow
}

// for debugging
func burrowFromString(str string) *Burrow {
	burrow := &Burrow{
		grid: &Grid{},
	}

	lines := strings.Split(str, "\n")

	for y, line := range lines {
		for x, char := range line {
			if char == ' ' || char == '.' {
				continue
			}
			pod := &Amphipod{
				x:     int8(x),
				y:     int8(y),
				_type: AmphipodType(char),
			}
			burrow.grid[y][x] = pod
			burrow.amphipods = append(burrow.amphipods, pod)
		}
	}

	return burrow
}

func ordered[T types.Numeric](i, j T) (T, T) {
	if i < j {
		return i, j
	}

	return j, i
}

func (burrow *Burrow) isHallwayClear(i, j int8) bool {
	min, max := ordered(i, j)

	for _, pod := range burrow.amphipods {
		if pod.y == 0 && pod.x >= min && pod.x <= max {
			return false
		}
	}

	return true
}

func sideRoomHasOtherTypes(grid *Grid, _type AmphipodType) bool {
	room := rooms[_type]

	for i := 1; i < 5; i++ {
		pod := grid[i][room]
		if pod != nil && pod._type != _type {
			return true
		}
	}

	return false
}

func sideRoomComplete(grid *Grid, _type AmphipodType) bool {
	room := rooms[_type]
	hasAnyPods := grid[1][room] != nil

	if !hasAnyPods {
		return false
	}

	for i := 1; i < 5; i++ {
		pod := grid[i][room]
		if pod == nil {
			// ignore nil pods if we're only dealing with 8 pods total
			break
		}
		if pod._type != _type {
			return false
		}
	}

	return true
}

func nextPodsAreSameType(grid *Grid, pod *Amphipod) bool {
	room := rooms[pod._type]

	for i := pod.y + 1; i < 5; i++ {
		ref := grid[i][room]
		if ref == nil {
			// ignore nil pods if we're only dealing with 8 pods total
			break
		}
		if ref._type != pod._type {
			return false
		}
	}

	return true
}

// 1. which pods can move
func (burrow *Burrow) getActivePods() []*Amphipod {
	// TODO: how many could possibly be active
	activePods := make([]*Amphipod, 0, 5)
	grid := burrow.grid

	// check hallway for pods that can go "home"
	for _, pod := range grid[0] {
		if pod == nil {
			continue
		}

		toRoom := rooms[pod._type]

		// check room is ready:
		roomHasOtherTypes := sideRoomHasOtherTypes(
			grid,
			pod._type,
		)

		// check path is clear
		podX := pod.x

		// don't include index of pod
		if toRoom < podX {
			podX--
		} else {
			podX++
		}
		isHallwayClear := burrow.isHallwayClear(podX, toRoom)

		canGoHome := !roomHasOtherTypes && isHallwayClear

		if canGoHome {
			activePods = append(activePods, pod)
		}
	}

	// check siderooms for pods that can go into the hallway
	var x, y int8
	for x = 2; x < 10; x += 2 {
		for y = 1; y < 5; y++ {
			pod := grid[y][x]

			if pod == nil {
				// dig deeper
				continue
			}

			isHome := x == rooms[pod._type]

			if isHome && nextPodsAreSameType(grid, pod) {
				// this pod/room is inactive
				break
			}

			// hallway has some room (left or right)
			isHallwayClear := burrow.isHallwayClear(x-1, x) || burrow.isHallwayClear(x, x+1)

			if isHallwayClear {
				activePods = append(activePods, pod)
			}

			// stop at first pod in room (subsequent pods are buried)
			break
		}
	}

	return activePods
}

func (burrow *Burrow) Copy() *Burrow {
	copy := &Burrow{
		amphipods: make([]*Amphipod, 0, 16),
		grid:      &Grid{},
		states:    make([]string, 0, len(burrow.states)+1),
	}

	copy.states = append(copy.states, burrow.states...)

	for _, pod := range burrow.amphipods {
		newPod := pod.Copy()
		copy.amphipods = append(copy.amphipods, newPod)
		copy.grid[pod.y][pod.x] = newPod
	}

	copy.cost = burrow.cost

	return copy
}

func (pod *Amphipod) Copy() *Amphipod {
	if pod == nil {
		return nil
	}

	return &Amphipod{
		x:     pod.x,
		y:     pod.y,
		_type: pod._type,
	}
}

func (burrow *Burrow) movePodTo(pod *Amphipod, x, y int8) {
	// update grid
	burrow.grid[y][x] = pod
	burrow.grid[pod.y][pod.x] = nil

	dx := utils.AbsInt(pod.x - x)
	dy := utils.AbsInt(pod.y - y)

	// update pod
	pod.x = x
	pod.y = y

	// update cost of travel
	burrow.cost += int(dx+dy) * costs[pod._type]
}

func (burrow *Burrow) getValidHallwayPositionsFromRoom(room int8) *[]int8 {
	var start int8 = 0
	var end int8 = 10

	//  reverse for start position
	for i := room; i >= 0; i-- {
		if burrow.grid[0][i] != nil {
			// blocked
			start = i + 1
			break
		}
	}

	for i := room; i < 11; i++ {
		if burrow.grid[0][i] != nil {
			// blocked
			end = i - 1
			break
		}
	}

	positions := make([]int8, 0, HALLWAY_SPOTS)

	for i := start; i <= end; i++ {
		// hallway room positions are invalid (2,4,6,8)
		if i%2 == 1 || i == 0 || i == 10 {
			positions = append(positions, i)
		}
	}

	return &positions
}

func (burrow *Burrow) sendPodToRoom(pod *Amphipod) {
	var y int8
	room := rooms[pod._type]
	var depth int8 = 2
	if len(burrow.amphipods) == 16 {
		depth = 4
	}

	// go in reverse for pod-homing
	for y = depth; y >= 1; y-- {
		space := burrow.grid[y][room]

		if space == nil {
			burrow.movePodTo(pod, room, y)
			return
		}
	}
}

// 2. where can each pod move
func (burrow *Burrow) getNextStates() *[]*Burrow {
	activePods := burrow.getActivePods()
	// the worst that could happen is the initial state, where
	// each pod could move out into the hallway (7 valid spots)
	nextStates := make([]*Burrow, 0, 7*len(activePods))

	for _, pod := range activePods {
		// if pod in hallway, pod moves to sideroom
		if pod.y == 0 {
			// single new state
			copy := burrow.Copy()
			newPod := copy.grid[pod.y][pod.x]
			copy.sendPodToRoom(newPod)

			nextStates = append(nextStates, copy)

			continue
		}

		// if pod in ANY(!!!) sideroom, pod moves anywhere in hallway
		positions := burrow.getValidHallwayPositionsFromRoom(pod.x)

		for _, hallwayX := range *positions {
			copy := burrow.Copy()
			newPod := copy.grid[pod.y][pod.x]
			// y=0 is hallway
			copy.movePodTo(newPod, hallwayX, 0)

			nextStates = append(nextStates, copy)
		}
	}

	return &nextStates
}

func (burrow *Burrow) isComplete() bool {
	grid := burrow.grid

	return sideRoomComplete(grid, A) &&
		sideRoomComplete(grid, B) &&
		sideRoomComplete(grid, C) &&
		sideRoomComplete(grid, D)
}

var cacheHits int
var iterations int
var cachedStates = map[string]int{}

func resetCaches() {
	cacheHits = 0
	iterations = 0
	cachedStates = map[string]int{}
}

func (this *Burrow) saveState() {
	this.states = append(this.states, this.String())
}

func (this *Burrow) play() int {
	resetCaches()
	pq := types.PriorityQueue[Burrow]{}
	pq.PushNewItem(this, 0)
	min := math.MaxInt
	// bestMoves := this

	for pq.Len() > 0 {
		burrow := pq.Get()
		iterations++

		// enabling this changes runtime from 40s to 1m17s
		// burrow.saveState()

		// if iterations%100000 == 0 {
		// 	log.Println("iterations", iterations, cacheHits)
		// }

		// log.Println("burrow", burrow)

		nextStates := burrow.getNextStates()

		if len(*nextStates) == 0 {
			// complete? failed?
			if burrow.isComplete() && burrow.cost < min {
				min = burrow.cost
				// bestMoves = burrow
				// log.Println("min so far", min, iterations, cacheHits)
			}

			continue
		}

		// log.Println("next", nextStates)

		// 3. push state to priority queue with cost
		for _, state := range *nextStates {
			if state.cost > min {
				continue
			}
			key := state.hash()
			cachedCost, ok := cachedStates[key]
			if !ok || cachedCost > state.cost {
				// update cost of state
				cachedStates[key] = state.cost
			} else if cachedCost <= state.cost {
				// there's a cheaper (OR EQUAL) sequence for this state
				cacheHits++
				continue
			}

			pq.PushNewItem(state, state.cost)
		}
	}

	fmt.Println("iterations", iterations)
	fmt.Println("cache hits", cacheHits)
	// log.Println("best moves", bestMoves.states)

	return min
}

//
// string representations
//

// TODO: is this optimized? or insane?
func (burrow *Burrow) hash() string {
	grid := burrow.grid
	toString := make([]byte, 11+4*4)

	for i, pod := range grid[0] {
		toString[i] = pod.String()[0]
	}

	i := 11

	for _, row := range grid[1:] {
		for j, pod := range row[2:9] {
			if j%2 == 0 {
				toString[i] = pod.String()[0]
				i++
			}
		}
	}

	return string(toString)
}

func (burrow *Burrow) String() string {
	var out strings.Builder
	grid := burrow.grid

	fmt.Fprintf(&out, "\ncost: %d\n", burrow.cost)

	for _, pod := range grid[0] {
		out.WriteString(pod.String())
	}

	count := len(burrow.amphipods)
	max := 3
	if count == 16 {
		max = 5
	}

	for _, row := range grid[1:max] {
		out.WriteString("\n  ")
		for i, pod := range row[2:9] {
			if i%2 == 0 {
				fmt.Fprintf(&out, "%s ", pod.String())
			}
		}
	}

	return out.String()
}

func (pod *Amphipod) String() (output string) {
	if pod == nil {
		return "."
	}
	return string(pod._type)
}
