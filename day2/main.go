package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func parts(tolerate bool) int {
	lines := utils.ReadFile("day2/input.txt")

	var res int

loop:
	for _, line := range lines {
		fields := strings.Fields(line)

		var increasing, decreasing bool
		var old int

		canTolerate := true

		for i, field := range fields {
			if i > 0 {
				current, _ := strconv.Atoi(field)
				diff := utils.Abs(old - current)

				if diff > 3 {
					if tolerate && canTolerate {
						canTolerate = false
						increasing = false
						increasing = false
						goto end
					}
					continue loop
				}

				if current == old {
					if tolerate && canTolerate {
						canTolerate = false
						goto end
					}
					continue loop
				} else if current > old {
					if decreasing {
						if tolerate && canTolerate {
							canTolerate = false
							goto end
						}
						continue loop
					}
					increasing = true
				} else {
					if increasing {
						if tolerate && canTolerate {
							canTolerate = false
							goto end
						}
						continue loop
					}
					decreasing = true
				}

			end:
				if i == len(fields)-1 {
					res++
				}
			}

			old, _ = strconv.Atoi(field)
		}
	}

	return res
}

func main() {
	fmt.Println(parts(false))
	fmt.Println(parts(true))
}
