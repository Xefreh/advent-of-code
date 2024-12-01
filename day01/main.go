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
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot open file. err=%v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

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

	for i := 0; i < len(input[0]); i++ {
		totalDistance += abs(input[0][i] - input[1][i])
	}

	return totalDistance
}

func partTwo(input [][]int) int {
	var similarityScore int
	cache := make(map[int]int)

	for _, v := range input[0] {
		if _, ok := cache[v]; ok {
			similarityScore += cache[v]
			continue
		}

		var count int
		for _, v2 := range input[1] {
			if v == v2 {
				count++
			}
		}

		res := v * count
		cache[v] = res
		similarityScore += res
	}

	return similarityScore
}

func main() {
	input := parseInput("input.txt")
	if input == nil {
		return
	}

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
