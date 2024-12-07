package main

import (
	"advent-of-code/utils"
	"fmt"
)

type WordSearch struct {
	Grid [][]string
}

func newWordSearch(lines []string) *WordSearch {
	ws := &WordSearch{}
	ws.Grid = make([][]string, len(lines))

	for i, line := range lines {
		ws.Grid[i] = make([]string, len(line))
		for j, char := range line {
			ws.Grid[i][j] = string(char)
		}
	}

	return ws
}

func isValid(ws *WordSearch, row, col int, char string) bool {
	return row >= 0 && row < len(ws.Grid) && col >= 0 && col < len(ws.Grid[0]) && ws.Grid[row][col] == char
}

func (ws *WordSearch) checkAllDirections(row, col int, word string) int {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}

	var count int

	for _, direction := range directions {
		currentRow, currentCol := row, col
		match := true

		for i := 1; i < len(word); i++ {
			currentRow += direction[0]
			currentCol += direction[1]

			if ok := isValid(ws, currentRow, currentCol, string(word[i])); !ok {
				match = false
				break
			}
		}

		if match {
			count++
		}
	}

	return count
}

func (ws *WordSearch) partOne() {
	word := "XMAS"
	var occurrences int

	for row := range ws.Grid {
		for col := range ws.Grid[row] {
			if ws.Grid[row][col] == string(word[0]) {
				occurrences += ws.checkAllDirections(row, col, word)
			}
		}
	}

	fmt.Println(occurrences)
}

func (ws *WordSearch) partTwo() {
	word := "MAS"
	var occurrences int
	first := string(word[0])
	last := string(word[len(word)-1])

	for row := range ws.Grid {
		for col := range ws.Grid[row] {
			if ws.Grid[row][col] == string(word[1]) {
				topLeftBottomRightMAS := isValid(ws, row-1, col-1, first) && isValid(ws, row+1, col+1, last)
				topRightBottomLeftMAS := isValid(ws, row-1, col+1, first) && isValid(ws, row+1, col-1, last)
				topLeftBottomRightSAM := isValid(ws, row-1, col-1, last) && isValid(ws, row+1, col+1, first)
				topRightBottomLeftSAM := isValid(ws, row-1, col+1, last) && isValid(ws, row+1, col-1, first)

				if (topLeftBottomRightMAS && topRightBottomLeftMAS) ||
					(topLeftBottomRightSAM && topRightBottomLeftSAM) ||
					(topLeftBottomRightMAS && topRightBottomLeftSAM) ||
					(topLeftBottomRightSAM && topRightBottomLeftMAS) {
					occurrences++
				}
			}
		}
	}

	fmt.Println(occurrences)
}

func main() {
	lines := utils.ReadFile("day4/input.txt")

	ws := newWordSearch(lines)
	ws.partOne()
	ws.partTwo()
}
