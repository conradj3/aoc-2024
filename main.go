package main

import (
	"fmt"
	"os"

	"github.com/conradj3/aoc-2024/day1"
	"github.com/conradj3/aoc-2024/day2"
	"github.com/conradj3/aoc-2024/day3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	switch day {
	case "1":
		part1, part2, err := day1.Solve("day1/input.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Day 1 - Part 1: %d\n", part1)
		fmt.Printf("Day 1 - Part 2: %d\n", part2)

	case "2":
		part1, part2, err := day2.Solve("day2/input.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Day 2 - Part 1: Number of Safe Reports: %d\n", part1)
		fmt.Printf("Day 2 - Part 2: Number of Safe Reports with Dampener: %d\n", part2)

	case "3":
		part1, _, err := day3.Solve("day3/input.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Day 3 - Part 1: %d\n", part1)
		part2, _, err := day3.SolveWithControl("day3/input.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Day 3 - Part 2: %d\n", part2)

	default:
		fmt.Printf("Day %s is not implemented yet\n", day)
	}
}
