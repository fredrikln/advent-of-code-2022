package main

import (
	"fmt"
	"strings"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	input := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func getPriority(character string) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", character) + 1
}

func part1(rucksacks []string) int {
	totalPriority := 0

rucksackloop:
	for _, rucksack := range rucksacks {
		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		inFirstCompartment := make(map[rune]bool)

		for _, character := range firstCompartment {
			inFirstCompartment[character] = true
		}

		for _, character := range secondCompartment {
			if _, ok := inFirstCompartment[character]; ok {
				totalPriority += getPriority(string(character))
				continue rucksackloop
			}
		}

	}

	return totalPriority
}

func findCommonItemForTwoStrings(a, b string) string {
	out := ""

	for _, char := range b {
		if strings.Index(a, string(char)) != -1 {
			out = out + string(char)
		}
	}

	// fmt.Println(out)

	return out
}

func findCommonItemForGroup(rucksacks []string) string {
	inAandB := findCommonItemForTwoStrings(rucksacks[0], rucksacks[1])

	inABandC := findCommonItemForTwoStrings(inAandB, rucksacks[2])

	return string([]rune(inABandC)[0])
}

func splitIntoGroupsOfThree(items []string) [][]string {
	out := make([][]string, 0)

	for i := 0; i < len(items); i += 3 {
		out = append(out, items[i:i+3])
	}

	return out
}

func part2(rucksacks []string) int {
	totalPriority := 0

	groups := splitIntoGroupsOfThree(rucksacks)
	for _, group := range groups {
		commonItem := findCommonItemForGroup(group)

		totalPriority += getPriority(commonItem)
	}

	return totalPriority
}
