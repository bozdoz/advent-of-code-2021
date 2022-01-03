package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/bozdoz/advent-of-code-2021/utils"
)

// different puzzles require different file loaders
var FileLoader = utils.LoadAsLines

// custom logger extended from the "log" package
var log = utils.Logger()

func init() {
	// disable logs when running (enabled in _test)
	log.SetOutput(ioutil.Discard)
}

func PartOne(content []string) (output int, err error) {
	goal := 1000
	game := startGame(content, goal)

	dice := &DeterministicDice{
		sides:   100,
		current: 1,
	}

	for {
		if game.playDeterministic(dice) {
			// game over
			break
		}
	}

	score := dice.rolls

	// get losing player's score
	for _, player := range game.players {
		if player.score < 1000 {
			score *= player.score
			break
		}
	}

	return score, nil
}

func PartTwo(content []string) (output int, err error) {
	goal := 21
	game := startGame(content, goal)

	// player one starts
	wins := game.playQuantum(PLAYER_ONE)

	winner := math.Max(
		float64(wins[PLAYER_ONE]),
		float64(wins[PLAYER_TWO]),
	)

	fmt.Println("cache hit", cacheHit)

	return int(winner), nil
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
