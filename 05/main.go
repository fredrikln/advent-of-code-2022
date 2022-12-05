package main

import (
	"fmt"
	"strings"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	instructions := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

func parseStacks(stacks []string) [][]string {
	firstLine := stacks[0] + " "
	out := make([][]string, len(firstLine)/4)

	for i := len(stacks) - 2; i >= 0; i-- {
		for j := 0; j < len(stacks[i])-1; j += 4 {

			box := strings.TrimSpace(stacks[i][j : j+3])

			if box == "" {
				continue
			}

			out[j/4] = append(out[j/4], string(box[1]))
		}
	}

	return out
}

func part1(input []string) string {
	stacks := make([]string, 0)
	moves := make([]string, 0)

	isMoves := false

	for _, row := range input {
		if row == "" {
			isMoves = true
			continue
		}

		if !isMoves {
			stacks = append(stacks, row)
		} else {
			moves = append(moves, row)
		}
	}

	parsedStacks := parseStacks(stacks)

	for _, move := range moves {
		var count, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)

		for i := 0; i < count; i++ {
			item := parsedStacks[from-1][len(parsedStacks[from-1])-1:]

			parsedStacks[from-1] = parsedStacks[from-1][:len(parsedStacks[from-1])-1]
			parsedStacks[to-1] = append(parsedStacks[to-1], item...)
		}
	}

	out := ""

	for _, stack := range parsedStacks {
		letter := stack[len(stack)-1]
		out += letter
	}

	return out
}

func part2(input []string) string {
	stacks := make([]string, 0)
	moves := make([]string, 0)

	isMoves := false

	for _, row := range input {
		if row == "" {
			isMoves = true
			continue
		}

		if !isMoves {
			stacks = append(stacks, row)
		} else {
			moves = append(moves, row)
		}
	}

	parsedStacks := parseStacks(stacks)

	for _, move := range moves {
		var count, from, to int
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)

		item := parsedStacks[from-1][len(parsedStacks[from-1])-count:]

		parsedStacks[from-1] = parsedStacks[from-1][:len(parsedStacks[from-1])-count]
		parsedStacks[to-1] = append(parsedStacks[to-1], item...)
	}

	out := ""

	for _, stack := range parsedStacks {
		letter := stack[len(stack)-1]
		out += letter
	}

	return out
}
