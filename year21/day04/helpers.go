package day04

import (
	"strconv"
	"strings"
)

func bingoToString(bingo *Bingo, choices []int, showGridView bool) string {
	var sb strings.Builder

	sb.WriteString("Choices:\n")

	for _, v := range choices {
		sb.WriteString(strconv.Itoa(v))
	}

	sb.WriteString("\nBoards:\n")

	for _, board := range bingo.boards {
		sb.WriteString(boardToString(board, showGridView))
	}

	return sb.String()
}

func boardToString(board *Board, showGridView bool) string {
	var sb strings.Builder

	sb.WriteString("Marked rows:\n")
	for row, cols := range board.markedRows {
		for col := range cols {
			sb.WriteString("(" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ")")
		}
	}

	sb.WriteString("\n")
	sb.WriteString("Marked Columns:\n")
	for col, rows := range board.markedRows {
		for row := range rows {
			sb.WriteString("(" + strconv.Itoa(col) + "," + strconv.Itoa(row) + ")")
		}
	}

	sb.WriteString("\n")
	sb.WriteString("Board:\n")

	if showGridView {
		var gridView [5][5]int

		for k, v := range board.grid {
			gridView[v[0]][v[1]] = k
		}

		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				sb.WriteString(strconv.Itoa(gridView[i][j]) + " ")
			}
			sb.WriteString("\n")
		}

	} else {
		for k, v := range board.grid {
			x := strconv.Itoa(v[0])
			y := strconv.Itoa(v[1])
			sb.WriteString(strconv.Itoa(k) + " -> " + x + "," + y + "\n")
		}
	}

	sb.WriteString("\n\n")

	return sb.String()
}
