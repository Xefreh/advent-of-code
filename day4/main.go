package main

import (
	"advent-of-code/utils"
	"fmt"
)

type wordSearch struct {
	grid [][]rune
}

func newWordSearch(lines []string) *wordSearch {
	ws := &wordSearch{}
	ws.grid = make([][]rune, len(lines))

	for i, line := range lines {
		ws.grid[i] = make([]rune, len(line))
		for j, char := range line {
			ws.grid[i][j] = char
		}
	}

	return ws
}

func isValid(ws *wordSearch, row, col int, char rune) bool {
	return row >= 0 && row < len(ws.grid) && col >= 0 && col < len(ws.grid[0]) && ws.grid[row][col] == char
}

func (ws *wordSearch) checkAllDirections(row, col int, word string) int {
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

			if ok := isValid(ws, currentRow, currentCol, rune(word[i])); !ok {
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

func (ws *wordSearch) partOne() {
	word := "XMAS"
	var occurrences int

	for row := range ws.grid {
		for col := range ws.grid[row] {
			if ws.grid[row][col] == rune(word[0]) {
				occurrences += ws.checkAllDirections(row, col, word)
			}
		}
	}

	fmt.Println(occurrences)
}

func (ws *wordSearch) partTwo() {
	word := "MAS"
	var occurrences int
	first := rune(word[0])
	last := rune(word[len(word)-1])

	for row := range ws.grid {
		for col := range ws.grid[row] {
			if ws.grid[row][col] == rune(word[1]) {
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
