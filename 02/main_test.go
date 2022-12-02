package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testcase := struct {
		input []string
		want  int
	}{
		input: []string{
			"A Y",
			"B X",
			"C Z",
		},
		want: 15,
	}

	got := part1(testcase.input)
	want := testcase.want

	if got != want {
		t.Errorf("got: %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	testcase := struct {
		input []string
		want  int
	}{
		input: []string{
			"A Y",
			"B X",
			"C Z",
		},
		want: 12,
	}

	got := part2(testcase.input)
	want := testcase.want

	if got != want {
		t.Errorf("got: %d, want %d", got, want)
	}
}
