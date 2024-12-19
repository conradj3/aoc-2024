package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the input file and returns two integer slices.
func ReadFile(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftNum, rightNum, err := ParseLine(scanner.Text())
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

// ParseLine parses a line of text into two integers.
func ParseLine(line string) (int, int, error) {
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
