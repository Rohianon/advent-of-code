package main

import (
	"strings"
	"testing"
)

const exampleInput = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestExamplePassword(t *testing.T) {
	deltas, err := parseRotations(strings.NewReader(exampleInput))
	if err != nil {
		t.Fatalf("parseRoations() err = %v", err)
	}

	got := countZeros(deltas, startPos, dialSize)
	want := 3

	if got != want {
		t.Fatalf("countZeros() = %d, want %d", got, want)
	}

}
