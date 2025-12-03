package main

import "testing"

func TestProcessRange(t *testing.T) {
	input := "11-22"
	expected := 33

	result, err := processRange(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Range 11-22: expected sum %d, got %d", expected, result)
	}
}

func TestIsRepeatedPattern(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{55, true},
		{123123, true},
		{101, false},
	}

	for _, tc := range tests {
		if res := isRepeatedPattern(tc.input); res != tc.expected {
			t.Errorf("isRepeatedPattern(%d) = %v; want %v", tc.input, res, tc.expected)
		}
	}
}

func TestIsMultiRepeatedPattern(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
		reason   string
	}{
		// Part 1 cases (still valid in Part 2)
		{11, true, "1 repeated twice"},
		{222222, true, "2 repeated 6 times OR 222 repeated twice"},

		// Part 2 NEW cases
		{123123123, true, "123 repeated 3 times"},
		{1212121212, true, "12 repeated 5 times"},
		{111, true, "1 repeated 3 times"},

		// Invalid cases
		{1234, false, "No repeating pattern"},
		{101, false, "Cannot divide 3 digits into equal repeating parts"},
		{123123124, false, "Almost repeating, but last digit differs"},
	}

	for _, tc := range tests {
		result := isMultiRepeatedPattern(tc.input)
		if result != tc.expected {
			t.Errorf("Input %d (%s): expected %v, got %v", tc.input, tc.reason, tc.expected, result)
		}
	}
}

func TestPart2ExampleSum(t *testing.T) {
	// The prompt gives a new expected sum for the example input
	// input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	expected := 4174379265

	// NOTE: We need to adapt the test to use the scanner or expose processRange logic
	// Since processRange handles a single range, we iterate here manually for the test

	// Split manually just for this test verification
	ranges := []string{
		"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659",
		"824824821-824824827", "2121212118-2121212124",
	}

	totalSum := 0
	for _, r := range ranges {
		val, _ := processRange(r)
		totalSum += val
	}

	if totalSum != expected {
		t.Errorf("Example sum failed. Expected %d, got %d", expected, totalSum)
	}
}
