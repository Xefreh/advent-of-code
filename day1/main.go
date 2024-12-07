package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInput(path string) [][]int {
	lines := utils.ReadFile(path)

	res := make([][]int, 2)
	for _, line := range lines {
		v := strings.Fields(line)

		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])

		res[0] = append(res[0], x)
		res[1] = append(res[1], y)
	}

	slices.Sort(res[0])
	slices.Sort(res[1])
	return res
}

func partOne(list [][]int) int {
	var totalDistance int

	for i := range list[0] {
		totalDistance += utils.Abs(list[0][i] - list[1][i])
	}

	return totalDistance
}

func partTwo(list [][]int) int {
	var similarityScore int
	freqMap := make(map[int]int)

	for _, v := range list[1] {
		freqMap[v]++
	}

	for _, v := range list[0] {
		count := freqMap[v]
		similarityScore += v * count
	}

	return similarityScore
}

func main() {
	list := parseInput("day1/input.txt")
	if list == nil {
		return
	}

	fmt.Println(partOne(list))
	fmt.Println(partTwo(list))
}
