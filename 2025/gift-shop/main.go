package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Create the Scanner on Stdin
	scanner := bufio.NewScanner(os.Stdin)

	// 2. MAGIC: Tell the scanner to split on COMMAS, not Newlines.
	// This lets us process "11-22", then "95-115", one by one.
	scanner.Split(scanCommaSeparated)

	totalSum := 0

	// 3. Loop through every "token" (range string) found
	for scanner.Scan() {
		// Get the text, e.g., "11-22"
		rawRange := scanner.Text()

		// Clean it up (remove newlines if they exist at the end of file)
		rawRange = strings.TrimSpace(rawRange)
		if rawRange == "" {
			continue
		}

		// Process this single range
		val, err := processRange(rawRange)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping invalid range '%s': %v\n", rawRange, err)
			continue
		}
		totalSum += val
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}

// scanCommaSeparated is a custom SplitFunc for bufio.Scanner
// It tells the scanner how to chop up the stream of bytes.
func scanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// If we're at EOF and have no data, stop.
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Look for a comma in the current chunk of data
	if i := bytes.IndexByte(data, ','); i >= 0 {
		// we found a comma at index 'i'
		// Return the data upto the comma as the token.
		// Tell scanner to advance past the comma (i + 1).
		return i + 1, data[0:i], nil
	}

	// If we're at EOF, we have a final chunk of data without a comma. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// request mode data
	return 0, nil, nil
}

func processRange(r string) (int, error) {
	bounds := strings.Split(r, "-")
	if len(bounds) != 2 {
		return 0, fmt.Errorf("bad format")
	}

	start, err1 := strconv.Atoi(bounds[0])
	end, err2 := strconv.Atoi(bounds[1])

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("bad numbers")
	}

	sum := 0
	for i := start; i <= end; i++ {
		if isMultiRepeatedPattern(i) {
			sum += i
		}
		// PART 1 solution
		// if isRepeatedPattern(i) {
		// 	sum += i
		// }
	}

	return sum, nil
}

func isMultiRepeatedPattern(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	// Try every possible chunk size 'k'
	// The chunk mus be at least length 1
	// The chunk cannot be larger than length / 2 (because we need atleast 2 chunks.)
	for k := 1; k <= length/2; k++ {
		// Math check: The total length must be divisible by the chunk size
		if length%k != 0 {
			continue
		}

		// extract the candidate pattern (the first k characters)
		pattern := s[:k]

		// Rebuild what the string should look like
		repeats := length / k
		rebuilt := strings.Repeat(pattern, repeats)

		// compare
		if rebuilt == s {
			return true
		}
	}
	return false
}

func isRepeatedPattern(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	if length%2 != 0 {
		return false
	}

	mid := length / 2
	return s[:mid] == s[mid:]
}
