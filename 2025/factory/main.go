package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalPresses := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		presses := solveMachine(line)
		totalPresses += presses
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Printf("Total Minimum Presses: %d\n", totalPresses)

}

func solveMachine(line string) int {

	parts := strings.Fields(line)

	diagramRaw := parts[0]
	diagramRaw = strings.Trim(diagramRaw, "[]")
	numLights := len(diagramRaw)

	target := make([]int, numLights)
	for i, char := range diagramRaw {
		if char == '#' {
			target[i] = 1
		} else {
			target[i] = 0
		}
	}

	var buttons [][]int
	for i := 1; i < len(parts); i++ {
		token := parts[i]

		// Stop when we hit the joltage section
		if strings.HasPrefix(token, "{") {
			break
		}

		// Clean parens (1,3) -> 1,3
		token = strings.Trim(token, "()")

		// Build the button vector
		btnVector := make([]int, numLights)
		indices := strings.Split(token, ",")
		for _, idxStr := range indices {
			idx, err := strconv.Atoi(idxStr)
			if err == nil && idx >= 0 && idx < numLights {
				btnVector[idx] = 1
			}
		}
		buttons = append(buttons, btnVector)
	}

	numVars := len(buttons)
	matrix := make([][]int, numLights)
	for r := 0; r < numLights; r++ {
		matrix[r] = make([]int, numVars+1)
		for c := 0; c < numVars; c++ {
			matrix[r][c] = buttons[c][r]
		}
		matrix[r][numVars] = target[r]
	}

	pivotRow := 0
	pivotCols := make([]int, numLights) // Stores which col is pivot for each row
	for i := range pivotCols {
		pivotCols[i] = -1
	}

	isPivotCol := make([]bool, numVars) // To identify free variables later

	for col := 0; col < numVars && pivotRow < numLights; col++ {
		// Find pivot in current column
		sel := -1
		for row := pivotRow; row < numLights; row++ {
			if matrix[row][col] == 1 {
				sel = row
				break
			}
		}

		if sel == -1 {
			continue // No pivot in this column, it's a free variable
		}

		// Swap rows
		matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]

		// Mark pivot
		pivotCols[pivotRow] = col
		isPivotCol[col] = true

		// Eliminate other rows
		for row := 0; row < numLights; row++ {
			if row != pivotRow && matrix[row][col] == 1 {
				// XOR row with pivot row
				for k := col; k <= numVars; k++ {
					matrix[row][k] ^= matrix[pivotRow][k]
				}
			}
		}
		pivotRow++
	}

	for r := pivotRow; r < numLights; r++ {
		if matrix[r][numVars] == 1 {
			// Impossible to solve.
			// (Advent of Code usually guarantees valid inputs, but good to handle)
			return 0
		}
	}

	var freeVars []int
	for c := 0; c < numVars; c++ {
		if !isPivotCol[c] {
			freeVars = append(freeVars, c)
		}
	}

	minPresses := math.MaxInt32

	numFree := len(freeVars)
	limit := 1 << numFree

	for i := 0; i < limit; i++ {
		// Construct a potential solution vector
		solution := make([]int, numVars)
		currentPresses := 0

		// Set free variables based on bits of i
		for k := 0; k < numFree; k++ {
			if (i>>k)&1 == 1 {
				solution[freeVars[k]] = 1
				currentPresses++
			}
		}

		// Solve for dependent variables (Pivots)
		// We process rows from bottom up (back substitution style,
		// though RREF makes it direct)
		for r := pivotRow - 1; r >= 0; r-- {
			pCol := pivotCols[r]
			if pCol == -1 {
				continue
			}

			// val = Target XOR (Sum of known variables in this row)
			val := matrix[r][numVars]
			for c := pCol + 1; c < numVars; c++ {
				if matrix[r][c] == 1 {
					val ^= solution[c]
				}
			}
			solution[pCol] = val
			if val == 1 {
				currentPresses++
			}
		}

		if currentPresses < minPresses {
			minPresses = currentPresses
		}
	}

	return minPresses

}
