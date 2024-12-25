package main

import (
	"advent-of-code/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type data struct {
	rules   map[int][]int
	updates [][]int
}

func newData(lines []string) *data {
	data := data{
		rules:   map[int][]int{},
		updates: [][]int{},
	}

	for _, line := range lines {
		if strings.Contains(line, "|") {
			splits := strings.Split(line, "|")
			
			x, _ := strconv.Atoi(splits[0])
			y, _ := strconv.Atoi(splits[1])

			data.rules[x] = append(data.rules[x], y)
		} else if strings.Contains(line, ",") {
			splits := strings.Split(line, ",")
			update := []int{}

			for _, split := range splits {
				x, _ := strconv.Atoi(split)
				update = append(update, x)
			}

			data.updates = append(data.updates, update)
		}
	}

	return &data
}

func (d *data) partOne() [][]int {
	var sum int
	var badUpdates [][]int

	for _, update := range d.updates {
		skip := false

		for i, n := range update {
			mustBefore := d.rules[n]
			numbersBefore := update[:i]

			for _, x := range mustBefore {
				if slices.Contains(numbersBefore, x) {
					badUpdates = append(badUpdates, update)
					skip = true
					break
				}
			}

			if skip {
				break
			}

			if i == len(update)-1 {
				sum += update[len(update)/2]
			}
		}
	}

	fmt.Println(sum)
	return badUpdates
}

func (d *data) partTwo(badUpdates [][]int) {
	var permutations int

	for _, update := range badUpdates {
		for i := 0; i < len(update); i++ {
			mustBefore := d.rules[update[i]]

			for j := 0; j < i; j++ {
				if slices.Index(mustBefore, update[j]) != -1 {
					update[i], update[j] = update[j], update[i]
					permutations++
				}
			}
		}
	}

	if permutations > 0 {
		d.partTwo(badUpdates)
	} else {
		var res int

		for _, update := range badUpdates {
			res += update[len(update)/2]
		}

		fmt.Println(res)
	}
}

func main() {
	lines := utils.ReadFile("day5/input.txt")

	data := newData(lines)
	badUpdates := data.partOne()
	data.partTwo(badUpdates)
}
