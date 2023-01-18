// Package day04 - Solution for Advent of Code 2021/04
// Problem Link: http://adventofcode.com/2021/day/04
package day04

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	bingo, choices := parse(input)
	return playBingo(bingo, choices, true), playBingo(bingo, choices, false)
}

func playBingo(bingo *Bingo, choices []int, firstWinnerOnly bool) int {
	for _, n := range choices {
		bingo.draw(n)
		if firstWinnerOnly && bingo.score != 0 {
			return bingo.score
		}
	}
	return bingo.score
}

func parse(input string) (*Bingo, []int) {
	inputGroups := strings.Split(input, "\n\n")
	choices := utils.SplitByInts(inputGroups[0], ",")
	bingo := &Bingo{}

	for _, boardInput := range inputGroups[1:] {
		boardRows := strings.Split(boardInput, "\n")
		bingo.boards = append(bingo.boards, NewBoard(boardRows))
	}

	return bingo, choices
}

type InvertedGrid map[int][2]int
type DimSet map[int]map[int]bool

type Bingo struct {
	boards []*Board
	score  int
}

func (bingo *Bingo) draw(number int) {
	for _, board := range bingo.boards {
		if !board.hasWon {
			row, col := board.grid[number][1], board.grid[number][0]

			if _, exists := board.grid[number]; exists && !board.markedRows[row][col] {
				board.totalScore -= number
				board.markedRows[row][col], board.markedCols[col][row] = true, true

				if len(board.markedRows[row]) == 5 || len(board.markedCols[col]) == 5 {
					bingo.score = number * board.totalScore
					board.hasWon = true
				}
			}
		}
	}
}

type Board struct {
	grid       InvertedGrid
	markedRows DimSet
	markedCols DimSet
	totalScore int
	hasWon     bool
}

func NewBoard(rows []string) *Board {
	grid := make(InvertedGrid)
	markedRows, markedCols := getMarkedDimensions()
	initialScore := 0

	for i, line := range rows {
		row := utils.SplitByInts(line, " ")
		for j, cell := range row {
			grid[cell] = [2]int{i, j}
			initialScore += cell
		}
	}

	return &Board{grid, markedRows, markedCols, initialScore, false}
}

func getMarkedDimensions() (DimSet, DimSet) {
	markedRows, markedCols := DimSet{}, DimSet{}

	for i := 0; i < 5; i++ {
		markedRows[i], markedCols[i] = map[int]bool{}, map[int]bool{}
	}

	return markedRows, markedCols
}
