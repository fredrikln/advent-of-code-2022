package main

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		desc   string
		signal string
		want   int
	}{
		{
			desc:   "1",
			signal: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:   5,
		},
		{
			desc:   "2",
			signal: "nppdvjthqldpwncqszvftbrmjlhg",
			want:   6,
		},
		{
			desc:   "3",
			signal: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:   10,
		},
		{
			desc:   "4",
			signal: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:   11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := part1(tC.signal)

			if got != tC.want {
				t.Errorf("Got %v want %v", got, tC.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		desc   string
		signal string
		want   int
	}{
		{
			desc:   "1",
			signal: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:   19,
		},
		{
			desc:   "2",
			signal: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:   23,
		},
		{
			desc:   "3",
			signal: "nppdvjthqldpwncqszvftbrmjlhg",
			want:   23,
		},
		{
			desc:   "4",
			signal: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:   29,
		},
		{
			desc:   "4",
			signal: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:   26,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := part2(tC.signal)

			if got != tC.want {
				t.Errorf("Got %v want %v", got, tC.want)
			}
		})
	}
}
