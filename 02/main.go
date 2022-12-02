package main

import (
	"fmt"
	"strings"

	"github.com/fredrikln/advent-of-code-2022/parsing"
)

func main() {
	data := parsing.ReadRowsToStringSlice("input.txt")

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}

func part1(strategy []string) int {
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	toWin := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	shape := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	score := 0

	for _, match := range strategy {
		shapes := strings.Split(match, " ")
		opponent := shapes[0]
		me := shapes[1]

		score += scores[me]

		if shape[opponent] == shape[me] {
			score += 3
			fmt.Println(opponent, me, "draw")
			continue
		}

		if me == toWin[opponent] {
			score += 6
			fmt.Println(opponent, me, "win")
			continue
		}

		fmt.Println(opponent, me, "loss")
	}

	return score
}

func part2(strategy []string) int {
	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	strategyMap := map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	howToLose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	howToDraw := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	howToWin := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	score := 0

	for _, match := range strategy {
		shapes := strings.Split(match, " ")
		opponent := shapes[0]
		whatToDo := strategyMap[shapes[1]]

		if whatToDo == "lose" {
			me := howToLose[opponent]
			score += scores[me]
		} else if whatToDo == "draw" {
			me := howToDraw[opponent]
			score += scores[me] + 3
		} else {
			me := howToWin[opponent]
			score += scores[me] + 6
		}
	}

	return score
}
