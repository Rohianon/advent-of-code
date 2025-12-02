package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	dialSize = 100
	startPos = 50
)

func main() {
	deltas, err := parseRotations(os.Stdin)
	if err != nil {
		log.Fatalf("parse input :%v", err)
	}

	password := countZeros(deltas, startPos, dialSize)
	fmt.Println(password)
}

// ParseRotations reads lines like "L68", "R5", etc and returns
// a slice of signed deltas
// L -> negative, R -> positive.
func parseRotations(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	deltas := make([]int, 0, 1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if len(line) < 2 {
			return nil, fmt.Errorf("invalid line %q: too short", line)
		}

		dir := line[0]
		stepsStr := strings.TrimSpace(line[1:])

		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			return nil, fmt.Errorf("parsing steps in %q: %w", line, err)
		}

		var delta int
		switch dir {
		case 'L':
			delta = -steps
		case 'R':
			delta = steps
		default:
			return nil, fmt.Errorf("invalid direct %q in line %q", string(dir), line)
		}
		deltas = append(deltas, delta)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return deltas, nil
}

func countZeros(deltas []int, start, size int) int {
	pos := start
	hits := 0

	for _, d := range deltas {
		pos = (pos + d) % size
		if pos < 0 {
			pos += size
		}

		if pos == 0 {
			hits++
		}
	}

	return hits

}
