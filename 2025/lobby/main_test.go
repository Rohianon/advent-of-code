package main

import "testing"

func TestFindMaxJoltage(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// Examples from prompt
		{"987654321111111", 98}, // 9 at start, 8 next to it
		{"811111111111119", 89}, // 8 at start, 9 at very end
		{"234234234234278", 78}, // 7 and 8 at the end
		{"818181911112111", 92}, // 9 is index 6, 2 is index 11. (92 > 81, 89, etc)

		// Edge cases
		{"12", 12},    // Only two digits
		{"55", 55},    // Duplicates
		{"12345", 45}, // Ascending order: max is last two
		{"54321", 54}, // Descending order: max is first two
	}

	for _, tc := range tests {
		got := findMaxJoltage(tc.input)
		if got != tc.expected {
			t.Errorf("findMaxJoltage(%q) = %d; want %d", tc.input, got, tc.expected)
		}
	}
}
