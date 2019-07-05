package main

import "testing"

func TestMaze(t *testing.T) {
	cases := []struct {
		in   maze
		want string
	}{
		{maze{
			Name:             "Maze Test 1",
			StartingPosition: []int{4, 3},
			EndingPosition:   []int{3, 6},
			Map: [][]string{
				{" ", " ", "X", " ", " ", " ", "X", " ", "X", "X"},
				{" ", "X", " ", " ", " ", " ", " ", " ", " ", " "},
				{" ", "X", " ", "X", "X", "X", "X", "X", "X", " "},
				{" ", "X", " ", " ", "A", " ", " ", " ", "X", " "},
				{" ", "X", "X", "X", "X", "X", "X", "X", " ", " "},
				{"X", " ", " ", " ", "X", " ", " ", " ", "X", " "},
				{" ", " ", "X", "B", "X", " ", "X", " ", "X", " "},
				{" ", " ", "X", " ", "X", " ", "X", " ", " ", " "},
				{"X", " ", "X", "X", "X", "X", "X", " ", "X", "X"},
				{"X", " ", " ", " ", " ", " ", " ", " ", "X", "X"},
			},
			Width:  10,
			Height: 10,
		}, "WWNNEEEEEEESSSSSSWWSSWWWWWWNNNNEES"},
	}
	for _, c := range cases {
		got := solve(c.in)
		if got != c.want {
			t.Errorf("solve(%v) == %v, want %v", c.in.Name, got, c.want)
		}
	}
}
