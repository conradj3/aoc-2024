package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solve processes the input file and returns the number of safe reports for both parts
func Solve(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	safeCountPart1 := 0
	safeCountPart2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)
		report := make([]int, len(levels))
		for i, level := range levels {
			report[i], _ = strconv.Atoi(level)
		}
		if isSafe(report) {
			safeCountPart1++
			safeCountPart2++ // Already safe, counts for both parts
		} else if isSafeWithDampener(report) {
			safeCountPart2++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("error reading file: %w", err)
	}
	return safeCountPart1, safeCountPart2, nil
}

// isSafe checks if a report is safe based on the problem rules
func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check the difference range
		if diff < -3 || diff > 3 {
			return false
		}

		// Check consistency in direction
		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false
		}
	}

	return true
}

// isSafeWithDampener checks if a report is safe after removing a single level
func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true // Already safe without any removals
	}

	// Try removing each level one at a time
	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if isSafe(modifiedReport) {
			return true
		}
	}

	return false
}
