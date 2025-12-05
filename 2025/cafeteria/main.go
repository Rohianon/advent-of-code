package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Range struct to hold start and end (inclusive)
type Range struct {
	Start int
	End   int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var ranges []Range
	readingRanges := true // State flag
	freshCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 1. Handle the "Section Switch"
		// If we hit any empty line, we are done with ranges and moving to IDs
		if line == "" {
			readingRanges = false
			continue
		}

		if readingRanges {
			// Parsing ranges (Top Section)
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				continue
			}
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			ranges = append(ranges, Range{Start: start, End: end})
		} else {
			// Parsing IDs (Bottom Section)
			// Format: "5"
			id, err := strconv.Atoi(line)
			if err != nil {
				continue
			}

			if isIngredientFresh(id, ranges) {
				freshCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Fresh Ingredients: %d\n", freshCount)
}

// isIngredientFresh checks if the ID falls into at leas one range
func isIngredientFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}
