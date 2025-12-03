package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalJoltage := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// calculate max joltage for this specific bank (line)
		bankMax := findMaxJoltage(line)
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
