package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1. Read the entire grid into memory
	// A slice of strings acts perfectly as a 2D grid
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// 2. Process the grid
	result := CountAccessibleRolls(grid)
	fmt.Printf("Accessible Rolls: %d\n", result)
}

func CountAccessibleRolls(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	accessibleCount := 0

	// Iterate over every cell in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// We only care if the current spot is a paper roll '@'
			if grid[r][c] == '@' {
				neighbors := countNeighbors(grid, r, c)
				if neighbors < 4 {
					accessibleCount++
				}
			}
		}
	}
	return accessibleCount
}

func countNeighbors(grid []string, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// goal is to look for offsets -1, 0, 1 for both row and column
	// This covers top-left (-1, -1) to bottom right (1, 1)
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			nr, nc := r+dr, c+dc

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if grid[nr][nc] == '@' {
					count++
				}
			}
		}
	}
	return count
}
