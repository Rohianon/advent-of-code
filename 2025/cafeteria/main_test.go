package main

import "testing"

func TestIsIngredientFresh(t *testing.T) {
	// Define the ranges from the example
	ranges := []Range{
		{Start: 3, End: 5},
		{Start: 10, End: 14},
		{Start: 16, End: 20},
		{Start: 12, End: 18}, // Note: this overlaps with 10-14 and 16-20
	}

	tests := []struct {
		id       int
		expected bool
	}{
		{1, false},  // Out of range
		{5, true},   // Inside 3-5
		{8, false},  // Out of range
		{11, true},  // Inside 10-14
		{17, true},  // Inside 16-20 AND 12-18
		{32, false}, // Out of range
		{3, true},   // Boundary check (start)
		{20, true},  // Boundary check (end)
	}

	for _, tc := range tests {
		result := isIngredientFresh(tc.id, ranges)
		if result != tc.expected {
			t.Errorf("ID %d: expected %v, got %v", tc.id, tc.expected, result)
		}
	}
}
