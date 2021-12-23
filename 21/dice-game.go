package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type PlayerType int

const (
	PLAYER_ONE PlayerType = iota
	PLAYER_TWO
)

type Player struct {
	position, score int
}

// 1-indexed auto-incremented die
type DeterministicDice struct {
	sides, current, rolls int
}

type Game struct {
	players []*Player
	goal    int
}

func startGame(data []string, goal int) *Game {
	game := &Game{goal: goal}

	for _, line := range data {
		re := regexp.MustCompile(`\d*$`)
		val := string(re.Find([]byte(line)))

		if val == "" {
			panic(fmt.Sprint("couldn't parse val from line", line))
		}

		num, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		game.players = append(game.players, &Player{
			position: num,
		})
	}

	return game
}

func (player *Player) won(goal int) bool {
	return player.score >= goal
}

func (game *Game) isGameOver() bool {
	for _, player := range game.players {
		if player.won(game.goal) {
			return true
		}
	}

	return false
}

// each player takes turns
func (game *Game) playDeterministic(dice *DeterministicDice) bool {
	if game.isGameOver() {
		return true
	}

	for _, player := range game.players {
		// player takes their turn
		player.playDeterministic(dice)

		if player.won(game.goal) {
			// game over
			return true
		}
	}

	return false
}

// inspired by @Torakushi
func (game *Game) playQuantum(current PlayerType) []int {
	wins := make([]int, 2)

	for roll, universes := range *getAllPossibleUniverses() {
		// each universe gets a copy of the board state
		altGame := game.Copy()
		currentPlayer := altGame.players[current]
		currentPlayer.update(roll)
		if currentPlayer.won(game.goal) {
			// winner
			wins[current] += universes
		} else {
			// next players turn (alternate from 0 to 1)
			var nextPlayer PlayerType = 1 - current

			// get all the wins from the subsequent universes
			altWins := altGame.playQuantumWithCache(nextPlayer)

			// returns the sum of all the wins in all subsequent universes
			wins[0] += universes * altWins[0]
			wins[1] += universes * altWins[1]
		}
	}

	return wins
}

var quantumCache = map[string][]int{}

// cache gets hit 96,257 times
var cacheHit int

// cache of playQuantum (saves us ~26 seconds in the test)
func (game *Game) playQuantumWithCache(current PlayerType) []int {
	p1 := game.players[PLAYER_ONE]
	p2 := game.players[PLAYER_TWO]
	key := fmt.Sprintf("%v-%v-%v", current, p1, p2)

	wins, ok := quantumCache[key]

	if !ok {
		wins = game.playQuantum(current)
		quantumCache[key] = wins
	} else {
		cacheHit++
	}

	return wins
}

var cachedPossibleUniverses *map[int]int

func getAllPossibleUniverses() *map[int]int {
	if cachedPossibleUniverses == nil {
		cachedPossibleUniverses = &map[int]int{}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					(*cachedPossibleUniverses)[i+j+k+3]++
				}
			}
		}
	}

	return cachedPossibleUniverses
}

func (game *Game) Copy() *Game {
	players := []*Player{}

	for _, player := range game.players {
		players = append(players, &Player{
			position: player.position,
			score:    player.score,
		})
	}

	return &Game{
		players: players,
		goal:    game.goal,
	}
}

func (player *Player) update(newPosition int) {
	// board wraps around
	position := (newPosition + player.position) % 10

	if position == 0 {
		// 0 means 10, since it's 1-indexed
		position = 10
	}

	player.position = position
	player.score += position
}

func (player *Player) playDeterministic(dice *DeterministicDice) {
	move := 0
	for i := 0; i < 3; i++ {
		move += dice.roll()
	}

	player.update(move)
}

func (dice *DeterministicDice) roll() int {
	current := dice.current

	if current == dice.sides {
		// loop back to 1
		dice.current = 1
	} else {
		dice.current++
	}

	dice.rolls++

	return current
}

//
// string representations
//

func (game *Game) String() (output string) {
	for i, player := range game.players {
		output += "\nPlayer "
		output += fmt.Sprint(i + 1)
		output += ": "
		output += "\n  score:    "
		output += fmt.Sprint(player.score)
		output += "\n  position: "
		output += fmt.Sprint(player.position)
	}

	return
}
