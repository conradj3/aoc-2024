package day1

import (
	"sort"

	"github.com/conradj3/aoc-2024/shared"
)

// Solve solves both parts of Day 1 returning results
func Solve(inputPath string) (int, int, error) {
	leftList, rightList, err := shared.ReadFile(inputPath)
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
