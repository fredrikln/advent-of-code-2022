package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	input := parsing.ReadFirstLineToString("input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func getStart(signal string, length int) int {
signalLoop:
	for i := 0; i < len(signal)-length; i++ {
		letters := signal[i : i+length]
		seen := make(map[rune]bool)

		for _, letter := range letters {
			_, ok := seen[letter]

			if ok {
				continue signalLoop
			}

			seen[letter] = true
		}

		return i + length
	}

	return 0
}

func part1(signal string) int {
	start := getStart(signal, 4)

	return start
}

func part2(signal string) int {
	start := getStart(signal, 14)

	return start
}
