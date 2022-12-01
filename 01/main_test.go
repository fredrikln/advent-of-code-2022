package main

import "testing"

func TestPart1(t *testing.T) {
	data := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	got := part1(data)
	want := 24000

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	data := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	got := part2(data)
	want := 45000

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
