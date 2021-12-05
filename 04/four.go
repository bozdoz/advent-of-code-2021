package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// maybe utility
func LoadAsString(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

type Bingo [5][5]int

// bingo board
func formatBoard(boardStr string) (board Bingo, err error) {
	rows := strings.Split(boardStr, "\n")

	for r, row := range rows {
		cols := strings.Fields(row)
		for c, val := range cols {
			board[r][c], err = strconv.Atoi(val)

			if err != nil {
				return board, err
			}
		}
	}

	return
}

func formatBoards(bulk []string) (boards []Bingo, err error) {
	for _, board := range bulk {
		formatted, err := formatBoard(board)

		if err != nil {
			return boards, err
		}

		boards = append(boards, formatted)
	}

	return
}

type Scored [5][5]bool

func all(arr [5]bool) bool {
	for _, val := range arr {
		if !val {
			return false
		}
	}

	return true
}

func checkcol(board Scored, col int) bool {
	for _, row := range board {
		if !row[col] {
			return false
		}
	}

	return true
}

// let's mark the number down and then check
func checkScoreBoards(boards []Bingo, scoreboards *[]Scored, check int) map[int]bool {
	winners := make(map[int]bool)

	// label to skip to the next board
boards:
	for b, board := range boards {
		for r, row := range board {
			for c, val := range row {
				if val == check {
					// mark the scoreboard!
					(*scoreboards)[b][r][c] = true
					scoreboard := (*scoreboards)[b]

					// check row & col
					if checkcol(scoreboard, c) || all(scoreboard[r]) {
						// winner winner
						winners[b] = true
					}

					// skip the rest of the checks for this board
					continue boards
				}
			}
		}
	}

	return winners
}

func sumUnmarked(board Bingo, scoreboard Scored) (sum int) {
	for r, row := range board {
		for c, val := range row {
			if !scoreboard[r][c] {
				// unmarked
				sum += val
			}
		}
	}

	return
}

// which board will win first
// what's the final score?
func PartOne(content string) (score int, err error) {
	parts := strings.Split(content, "\n\n")

	numbers := strings.Split(parts[0], ",")
	boards, err := formatBoards(parts[1:])

	if err != nil {
		return
	}

	scoreboards := make([]Scored, len(boards))

	for _, val := range numbers {
		num, err := strconv.Atoi(val)

		if err != nil {
			return -1, err
		}

		// let's mark the number down and then check
		winners := checkScoreBoards(boards, &scoreboards, num)

		if len(winners) > 0 {
			// I think we only care about the first winner
			for i := range winners {
				boardSum := sumUnmarked(boards[i], scoreboards[i])
				score := boardSum * num

				return score, nil
			}
		}
	}

	return
}

func PartTwo(content string) (output int, err error) {
	parts := strings.Split(content, "\n\n")

	numbers := strings.Split(parts[0], ",")
	boards, err := formatBoards(parts[1:])

	if err != nil {
		return
	}

	scoreboards := make([]Scored, len(boards))

	for _, val := range numbers {
		num, err := strconv.Atoi(val)

		if err != nil {
			return -1, err
		}

		// let's mark the number down and then check
		winners := checkScoreBoards(boards, &scoreboards, num)

		// no winners means we don't remove any boards from play
		if len(winners) == 0 {
			continue
		}

		// one board left means it finally won and it's over
		if len(boards) == 1 {
			for i := range winners {
				boardSum := sumUnmarked(boards[i], scoreboards[i])
				score := boardSum * num

				// happy path
				return score, nil
			}
		}

		// filter out the winning boards
		filteredBoards := []Bingo{}
		for i, board := range boards {
			if !winners[i] {
				filteredBoards = append(filteredBoards, board)
			}
		}

		filteredScoreboards := []Scored{}
		for i, board := range scoreboards {
			if !winners[i] {
				filteredScoreboards = append(filteredScoreboards, board)
			}
		}

		// swap boards as a way to filter
		boards = filteredBoards
		scoreboards = filteredScoreboards
	}

	return
}

func main() {
	// safe to assume
	filename := "input.txt"

	vals := LoadAsString(filename)

	answer, err := PartOne(vals)

	if err != nil {
		fmt.Println("failed to parse PartOne", err)
		return
	}

	fmt.Printf("Part One: %d \n", answer)

	answer2, err := PartTwo(vals)

	if err != nil {
		fmt.Println("failed to parse PartTwo", err)
		return
	}

	fmt.Printf("Part Two: %d \n", answer2)
}
