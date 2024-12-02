package main

import (
	"fmt"
	"os"

	"github.com/conradj3/aoc-2024/day1"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	switch day {
	case "1":
		totalDistance, similarityScore, err := day1.Solve("day1/input.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Day 1 - Part 1: %d\n", totalDistance)
		fmt.Printf("Day 1 - Part 2: %d\n", similarityScore)

	default:
		fmt.Printf("Day %s is not implemented yet\n", day)
	}
}
