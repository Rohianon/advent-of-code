package main

import (
	"strings"
	"testing"
)

func TestCountAccessibleRolls(t *testing.T) {
	// The example input from the prompt
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	// Convert the single string into the slice format our function expects
	grid := strings.Split(input, "\n")

	// The prompt says the answer for this example is 13
	expected := 13

	result := CountAccessibleRolls(grid)

	if result != expected {
		t.Errorf("Expected %d accessible rolls, got %d", expected, result)
	}
}
