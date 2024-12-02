package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/conradj3/aoc-2024/shared"
)

// Solve solves both parts of Day 1 returning results
func Solve(inputPath string) (int, int, error) {
	leftList, rightList, err := readFile(inputPath)
	if err != nil {
		return 0, 0, err
	}

	// Part 1: Total Distance
	sort.Ints(leftList)
	sort.Ints(rightList)
	totalDistance := calculateTotalDistance(leftList, rightList)

	// Part 2: Similarity Score
	similarityScore := calculateSimilarityScore(leftList, rightList)

	return totalDistance, similarityScore, nil
}

// readFile reads the input file and returns two integer slices.
func readFile(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftNum, rightNum, err := parseLine(scanner.Text())
		if err != nil {
			return nil, nil, err
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return leftList, rightList, nil
}

// parseLine parses a line of text into two integers.
func parseLine(line string) (int, int, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid line format: %s", line)
	}

	leftNum, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting left number: %w", err)
	}

	rightNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting right number: %w", err)
	}

	return leftNum, rightNum, nil
}

// calculateTotalDistance calculates the total distance between two lists.
func calculateTotalDistance(leftList, rightList []int) int {
	totalDistance := 0
	for i := range leftList {
		totalDistance += shared.Abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

// calculateSimilarityScore calculates the similarity score between two lists.
func calculateSimilarityScore(leftList, rightList []int) int {
	rightCounts := make(map[int]int)
	for _, num := range rightList {
		rightCounts[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		if count, exists := rightCounts[num]; exists {
			similarityScore += num * count
		}
	}
	return similarityScore
}
