package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
)

func partOne(lines []string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
 
	var matches [][]string
	for _, line := range lines {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	var res int
	for _, match := range matches {
		if len(match) > 2 && match[1] != "" && match[2] != "" {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			res += x * y
		}
	}

	return res
}

func partTwo(lines []string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	var matches [][]string
	for _, line := range lines {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	var res int
	canMultiply := true
	for _, match := range matches {
		if match[0] == "do()" {
			canMultiply = true
		} else if match[0] == "don't()" {
			canMultiply = false
		} else if len(match) > 2 && match[1] != "" && match[2] != "" && canMultiply {
			x, _ := strconv.Atoi(match[1])	
			y, _ := strconv.Atoi(match[2])

			res += x * y
		}
	}

	return res
}

func main() {
	lines := utils.ReadFile("day3/input.txt")

	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}
