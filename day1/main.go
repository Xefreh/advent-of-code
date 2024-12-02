package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parseInput(fileName string) [][]int {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read file. err=%v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

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

func partOne(input [][]int) int {
	var totalDistance int

	for i := range input[0] {
		totalDistance += abs(input[0][i] - input[1][i])
	}

	return totalDistance
}

func partTwo(input [][]int) int {
	var similarityScore int
	freqMap := make(map[int]int)

	for _, v := range input[1] {
		freqMap[v]++
	}

	for _, v := range input[0] {
		count := freqMap[v]
		similarityScore += v * count
	}

	return similarityScore
}

func main() {
	input := parseInput("day01/testinput.txt")
	if input == nil {
		return
	}

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
