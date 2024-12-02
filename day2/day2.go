package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parts(tolerate bool) int {
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatalf("Cannot read file. err=%v", err)
	}

	lines := strings.Split(string(data), "\n")

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
				diff := abs(old - current)

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
