package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	want := 2
	got := part1(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	want := 4
	got := part2(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCheckPairFullOverlap(t *testing.T) {
	testcases := []struct {
		pair string
		want bool
	}{
		{
			"2-4,6-8",
			false,
		},
		{
			"2-3,4-5",
			false,
		},
		{
			"5-7,7-9",
			false,
		},
		{
			"2-8,3-7",
			true,
		},
		{
			"6-6,4-6",
			true,
		},
		{
			"2-6,4-8",
			false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkPairFullOverlap(tc.pair)
			want := tc.want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func TestCheckPairPartialOverlap(t *testing.T) {
	testcases := []struct {
		pair string
		want bool
	}{
		{
			"2-4,6-8",
			false,
		},
		{
			"2-3,4-5",
			false,
		},
		{
			"5-7,7-9",
			true,
		},
		{
			"2-8,3-7",
			true,
		},
		{
			"6-6,4-6",
			true,
		},
		{
			"2-6,4-8",
			true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := checkPairPartialOverlap(tc.pair)
			want := tc.want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
