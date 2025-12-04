package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totalJoltage int64 = 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// calculate max joltage for this specific bank (line)
		// bankMax := findMaxJoltage(line)
		bankMax := calculateStaticFriction(line)
		totalJoltage += bankMax

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Output Joltage: %d\n", totalJoltage)
}

func findMaxJoltage(s string) int {
	maxVal := 0
	n := len(s)

	if n < 2 {
		return 0
	}

	for i := 0; i < n-1; i++ {
		// Inner loop picks the second digit (Ones place)
		// Goes from i+1 to the end
		for j := i + 1; j < n; j++ {
			// Get the integer value of the char digits
			// In Go, s[i] is a byte (ASCII value)
			// Subtracting '0' converst ASCII '9' (57) to int 9.
			d1 := int(s[i] - '0')
			d2 := int(s[j] - '0')

			// combine them: if d1=9, d2=2 -> 90 + 2 = 92
			val := (d1 * 10) + d2

			if val > maxVal {
				maxVal = val
			}

		}
	}
	return maxVal
}

func calculateStaticFriction(s string) int64 {
	// if the string is shorter than 12 chars, we can't select 12.
	if len(s) < 12 {
		return 0
	}

	var resultBuilder strings.Builder

	digitsNeeded := 12
	currentSearchStart := 0
	length := len(s)

	for digitsNeeded > 0 {
		// We must leave enough room for the remaining digits.
		// For example, if we need 12 digits total, we can't pick the very last char as our first digit.
		// The furthers we can look is length - digitsneeded
		maxSearchIndex := length - digitsNeeded
		bestDigit := -1
		bestIndex := -1

		// Scan the valid window for the largest digit
		for i := currentSearchStart; i <= maxSearchIndex; i++ {
			digit := int(s[i] - '0')

			// Optimization: if we find a 9, we can't get better.
			// Pick it and stop scanning this window
			if digit == 9 {
				bestDigit = 9
				bestIndex = i
				break
			}

			if digit > bestDigit {
				bestDigit = digit
				bestIndex = i
			}

		}
		// Append the winner to our result
		resultBuilder.WriteByte(s[bestIndex])

		// Move our search start to right after the digit we just picked
		currentSearchStart = bestIndex + 1
		digitsNeeded--
	}

	// Convert the 12-char string to an int64
	val, _ := strconv.ParseInt(resultBuilder.String(), 10, 64)
	return val
}
