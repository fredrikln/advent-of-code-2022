package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	input := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func mapToInt(values []string) []int {
	out := make([]int, 0, len(values))
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		out = append(out, v)
	}

	return out
}

func checkFullOverlap(a0, a1, b0, b1 int) bool {
	if b0 >= a0 && b1 <= a1 {
		return true
	}

	return false
}

func checkPairFullOverlap(pair string) bool {
	elves := strings.Split(pair, ",")
	a := mapToInt(strings.Split(elves[0], "-"))
	b := mapToInt(strings.Split(elves[1], "-"))

	if checkFullOverlap(a[0], a[1], b[0], b[1]) || checkFullOverlap(b[0], b[1], a[0], a[1]) {
		return true
	}

	return false
}

func checkPartialOverlap(a0, a1, b0, b1 int) bool {
	if b0 <= a0 && b1 >= a0 {
		return true
	}

	if b0 <= a1 && b1 >= a1 {
		return true
	}

	if b0 <= a0 && b1 >= a1 {
		return true
	}

	if b0 >= a0 && b1 <= a1 {
		return true
	}

	return false
}

func checkPairPartialOverlap(pair string) bool {
	elves := strings.Split(pair, ",")
	a := mapToInt(strings.Split(elves[0], "-"))
	b := mapToInt(strings.Split(elves[1], "-"))

	if checkPartialOverlap(a[0], a[1], b[0], b[1]) || checkPartialOverlap(b[0], b[1], a[0], a[1]) {
		return true
	}

	return false
}

func part1(pairs []string) int {
	sum := 0

	for _, pair := range pairs {
		if checkPairFullOverlap(pair) {
			sum += 1
		}
	}

	return sum
}

func part2(pairs []string) int {
	sum := 0

	for _, pair := range pairs {
		if checkPairPartialOverlap(pair) {
			sum += 1
		}
	}

	return sum
}
