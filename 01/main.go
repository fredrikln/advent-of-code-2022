package main

import (
	"fmt"
	"strconv"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	input := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) int {
	curElf := 0
	curAmount := 0

	maxAmount := 0

	for i, line := range input {
		if line != "" {
			number, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			curAmount += number
		}

		if line == "" || i == len(input)-1 {
			if curAmount > maxAmount {
				maxAmount = curAmount
			}

			curElf++
			curAmount = 0

			continue
		}
	}

	return maxAmount
}

func part2(input []string) int {
	curAmount := 0

	maxAmount1 := 0
	maxAmount2 := 0
	maxAmount3 := 0

	for i, line := range input {
		if line != "" {
			number, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			curAmount += number
		}

		if line == "" || i == len(input)-1 {
			if curAmount > maxAmount1 {
				maxAmount3 = maxAmount2
				maxAmount2 = maxAmount1
				maxAmount1 = curAmount
			} else if curAmount > maxAmount2 {
				maxAmount3 = maxAmount2
				maxAmount2 = curAmount
			} else if curAmount > maxAmount3 {
				maxAmount3 = curAmount
			}

			curAmount = 0

			continue
		}
	}

	return maxAmount1 + maxAmount2 + maxAmount3
}
