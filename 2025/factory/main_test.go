package main

import (
	"testing"
)

func TestSolveMachine(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			// Example 1: Min presses is 2 (last two buttons)
			"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
			2,
		},
		{
			// Example 2: Min presses is 3
			"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
			3,
		},
		{
			// Example 3: Min presses is 2
			"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
			2,
		},
	}

	for _, tc := range tests {
		got := solveMachine(tc.input)
		if got != tc.expected {
			t.Errorf("Input: %s\nExpected %d, got %d", tc.input, tc.expected, got)
		}
	}
}

