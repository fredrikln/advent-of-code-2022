package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	got := part1(input)
	want := "CMZ"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	got := part2(input)
	want := "MCD"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
