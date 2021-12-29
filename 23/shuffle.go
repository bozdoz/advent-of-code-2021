package main

import (
	"fmt"
	"math"
	"regexp"

	"bozdoz.com/aoc-2021/types"
	"bozdoz.com/aoc-2021/utils"
)

type AmphipodType string

type Amphipod struct {
	_type AmphipodType
}

const (
	AMBER  AmphipodType = "A"
	BRONZE AmphipodType = "B"
	COPPER AmphipodType = "C"
	DESERT AmphipodType = "D"
)

var costs = map[AmphipodType]int{
	AMBER:  1,
	BRONZE: 10,
	COPPER: 100,
	DESERT: 1000,
}

var rooms = map[AmphipodType]int{
	AMBER:  0,
	BRONZE: 1,
	COPPER: 2,
	DESERT: 3,
}

type Burrow struct {
	hallway   [11]*Amphipod
	siderooms [4][]*Amphipod
	cost      int
}

func parseInput(data string) *Burrow {
	burrow := &Burrow{}

	re := regexp.MustCompile("[ABCD]")

	for i, match := range re.FindAllString(data, -1) {
		_type := AmphipodType(match)

		pod := &Amphipod{_type}

		room := i % 4

		burrow.siderooms[room] = append(burrow.siderooms[room], pod)
	}

	return burrow
}

// room should have same type or be empty
func sideRoomHasOtherTypes(sideroom []*Amphipod, _type AmphipodType) bool {
	for _, pod := range sideroom {
		if pod != nil && pod._type != _type {
			return true
		}
	}

	return false
}

func sideRoomIsComplete(sideroom []*Amphipod, _type AmphipodType) bool {
	for _, pod := range sideroom {
		if pod == nil || pod._type != _type {
			return false
		}
	}

	return true
}

func getFirstPodInSideroom(sideroom []*Amphipod) *Amphipod {
	for _, pod := range sideroom {
		if pod != nil {
			return pod
		}
	}

	return nil
}

func (burrow *Burrow) isHallwayClear(i, j int) bool {
	min := utils.MinInt(i, j)
	max := utils.MaxInt(i, j)
	for _, pod := range burrow.hallway[min : max+1] {
		if pod != nil {
			return false
		}
	}

	return true
}

func (burrow *Burrow) swapSideRoomtoHallway(room, pos, hallway int) {
	burrow.hallway[hallway] = burrow.siderooms[room][pos]
	burrow.siderooms[room][pos] = nil
}

func (burrow *Burrow) swapHallwaytoSideRoom(hallway, room, pos int) {
	burrow.siderooms[room][pos] = burrow.hallway[hallway]
	burrow.hallway[hallway] = nil
}

// pods nestled as deep as possible in sideroom
// with no other types deeper are inactive
func nextPodsAreSameType(sideroom []*Amphipod, pod *Amphipod) bool {
	startChecking := false
	for _, ref := range sideroom {
		if ref == pod {
			startChecking = true
			continue
		}
		if startChecking && ref._type != pod._type {
			return false
		}
	}

	return true
}

// 1. which pods can move
func (burrow *Burrow) getActivePods() (activePods []*Amphipod) {
	// check hallway for pods that can go "home"
	for i, pod := range burrow.hallway {
		if pod != nil {
			// check room is ready:
			room := rooms[pod._type]
			roomHasOtherTypes := sideRoomHasOtherTypes(
				burrow.siderooms[room],
				pod._type,
			)

			// check path is clear
			toRoom := (room + 1) * 2
			fromPod := i

			// don't include index of pod
			if toRoom < fromPod {
				fromPod--
			} else {
				fromPod++
			}
			isHallwayClear := burrow.isHallwayClear(fromPod, toRoom)

			canGoHome := !roomHasOtherTypes && isHallwayClear

			if canGoHome {
				activePods = append(activePods, pod)
			}
		}
	}

	// check siderooms for pods that can go into the hallway
	for i, sideroom := range burrow.siderooms {
		pod := getFirstPodInSideroom(sideroom)

		if pod == nil {
			continue
		}

		isHome := i == rooms[pod._type]

		if isHome && nextPodsAreSameType(sideroom, pod) {
			// this pod is inactive
			continue
		}

		// hallway has some room (left or right)
		hallwayAtRoom := (i + 1) * 2
		isHallwayClear := burrow.isHallwayClear(hallwayAtRoom-1, hallwayAtRoom) || burrow.isHallwayClear(hallwayAtRoom, hallwayAtRoom+1)

		if isHallwayClear {
			activePods = append(activePods, pod)
		}
	}

	return
}

func (burrow *Burrow) Copy() *Burrow {
	copy := &Burrow{}

	for i, pod := range burrow.hallway {
		copy.hallway[i] = pod.Copy()
	}

	for i, room := range burrow.siderooms {
		for _, pod := range room {
			copy.siderooms[i] = append(copy.siderooms[i], pod.Copy())
		}
	}

	copy.cost = burrow.cost

	return copy
}

func (pod *Amphipod) Copy() *Amphipod {
	if pod == nil {
		return nil
	}
	return &Amphipod{pod._type}
}

func (burrow *Burrow) travelToHallway(room, pos, hallwayPos int) {
	pod := burrow.siderooms[room][pos]
	hallwayEnd := (room + 1) * 2
	travel := 1 + pos

	travel += utils.AbsInt(hallwayPos - hallwayEnd)
	burrow.swapSideRoomtoHallway(room, pos, hallwayPos)

	// update cost of travel
	burrow.cost += travel * costs[pod._type]
}

func (burrow *Burrow) travelToRoom(hallwayPos int) {
	pod := burrow.hallway[hallwayPos]

	// pod can only go into its own room
	room := rooms[pod._type]
	sideroom := burrow.siderooms[room]
	hallwayEnd := (room + 1) * 2
	travel := utils.AbsInt(hallwayPos - hallwayEnd)

	// find room position
	roomLen := len(sideroom)
	// iterate in reverse
	for i := roomLen - 1; i >= 0; i-- {
		position := sideroom[i]

		if position == nil {
			// first empty position is where the pod must go
			travel += i + 1
			burrow.swapHallwaytoSideRoom(hallwayPos, room, i)
			break
		}
	}

	// update cost of travel
	burrow.cost += travel * costs[pod._type]
}

func (burrow *Burrow) getValidHallwayPositionsFromRoom(room int) (positions []int) {
	start := 0
	end := 10

	hallwayPos := (room + 1) * 2

	//  reverse for start position
	for i := hallwayPos; i >= 0; i-- {
		if burrow.hallway[i] != nil {
			// blocked
			start = i + 1
		}
	}

	for i := hallwayPos; i < 11; i++ {
		if burrow.hallway[i] != nil {
			// blocked
			end = i - 1
		}
	}

	for i := start; i <= end; i++ {
		// hallway room positions are invalid (2,4,6,8)
		if i%2 == 1 || i == 0 || i == 10 {
			positions = append(positions, i)
		}
	}

	return
}

// 2. where can each pod move
func (burrow *Burrow) getNextStates() []*Burrow {
	nextStates := []*Burrow{}
	activePods := burrow.getActivePods()

outer:
	for _, pod := range activePods {
		// if pod in hallway, pod moves to sideroom
		for i, ref := range burrow.hallway {
			if pod == ref {
				// single new state
				copy := burrow.Copy()
				copy.travelToRoom(i)

				nextStates = append(nextStates, copy)

				continue outer
			}
		}

		// if pod in ANY sideroom, pod moves anywhere in hallway
		for room, sideroom := range burrow.siderooms {
			for pos, ref := range sideroom {
				if ref == pod {
					positions := burrow.getValidHallwayPositionsFromRoom(room)

					for _, hallway := range positions {
						copy := burrow.Copy()
						copy.travelToHallway(room, pos, hallway)

						nextStates = append(nextStates, copy)
					}

					break
				}
			}
		}
	}

	return nextStates
}

var amphipodRooms = []AmphipodType{
	AMBER, BRONZE, COPPER, DESERT,
}

func (burrow *Burrow) isComplete() bool {
	for i, sideroom := range burrow.siderooms {
		_type := amphipodRooms[i]
		if !sideRoomIsComplete(sideroom, _type) {
			return false
		}
	}
	return true
}

var cacheHits int
var iterations int
var cachedStates = map[string]int{}

func (this *Burrow) play() int {
	pq := types.PriorityQueue[Burrow]{}
	pq.PushNewItem(this, 0)
	min := math.MaxInt

	for pq.Len() > 0 {
		burrow := pq.Get()
		iterations++

		if iterations%100000 == 0 {
			fmt.Println("iterations", iterations, cacheHits)
		}

		// log.Println("burrow", burrow)

		nextStates := burrow.getNextStates()

		// log.Println("next", nextStates)
		if len(nextStates) == 0 {
			// complete? failed?
			if burrow.isComplete() && burrow.cost < min {
				min = burrow.cost
				fmt.Println("min so far", min, iterations, cacheHits)
			}

			continue
		}

		// 3. push state to priority queue with cost
		for _, state := range nextStates {
			key := state.hash()
			cachedCost, ok := cachedStates[key]
			if !ok || cachedCost > state.cost {
				// update cost of state
				cachedStates[key] = state.cost
			} else if ok && cachedCost < state.cost {
				// there's a cheaper sequence for this state
				cacheHits++
				continue
			}

			pq.PushNewItem(state, state.cost)
		}
	}

	log.Println("iterations", iterations)
	log.Println("cache hits", cacheHits)

	return min
}

//
// string representations
//

func (burrow *Burrow) hash() (output string) {
	for _, pod := range burrow.hallway {
		output += pod.String()
	}

	for _, sideroom := range burrow.siderooms {
		for _, pod := range sideroom {
			output += pod.String()
		}
	}

	return
}

func (burrow *Burrow) String() (output string) {
	output += "\ncost: " + fmt.Sprint(burrow.cost)
	output += "\n"
	for _, pod := range burrow.hallway {
		output += pod.String()
	}

	for j := range burrow.siderooms[0] {
		output += fmt.Sprintf("\n  ")
		for _, sideroom := range burrow.siderooms {
			output += fmt.Sprintf("%v ", sideroom[j])
		}
	}

	return
}

func (pod *Amphipod) String() (output string) {
	if pod == nil {
		return "."
	}
	return string(pod._type)
}
