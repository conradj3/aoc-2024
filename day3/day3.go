package day3

import (
	"os"
	"regexp"
	"strconv"
)

// Solve processes the input file and returns the sum of all valid mul(X,Y) results and the count of such instructions.
func Solve(filePath string) (int, int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, 0, err
	}

	values, err := extractMulValues(string(content))
	if err != nil {
		return 0, 0, err
	}

	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum, len(values), nil
}

// SolveWithControl processes the input file and returns the sum of all valid mul(X,Y) results and the count of such instructions, considering do() and don't() controls.
func SolveWithControl(filePath string) (int, int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return 0, 0, err
	}

	values, err := extractMulValuesWithControl(string(content))
	if err != nil {
		return 0, 0, err
	}

	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum, len(values), nil
}

// extractMulValues extracts valid mul(X,Y) instructions from the input string and returns their results.
func extractMulValues(input string) ([]int, error) {
	var results []int
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) == 3 {
			a, err1 := strconv.Atoi(match[1])
			b, err2 := strconv.Atoi(match[2])
			if err1 == nil && err2 == nil {
				results = append(results, a*b)
			}
		}
	}

	return results, nil
}

// extractMulValuesWithControl extracts valid mul(X,Y) instructions from the input string and returns their results, considering do() and don't() controls.
func extractMulValuesWithControl(input string) ([]int, error) {
	var results []int
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	mulEnabled := true
	pos := 0

	for pos < len(input) {
		doMatch := reDo.FindStringIndex(input[pos:])
		dontMatch := reDont.FindStringIndex(input[pos:])
		mulMatch := reMul.FindStringSubmatchIndex(input[pos:])

		if doMatch != nil && (dontMatch == nil || doMatch[0] < dontMatch[0]) && (mulMatch == nil || doMatch[0] < mulMatch[0]) {
			mulEnabled = true
			pos += doMatch[1]
		} else if dontMatch != nil && (doMatch == nil || dontMatch[0] < doMatch[0]) && (mulMatch == nil || dontMatch[0] < mulMatch[0]) {
			mulEnabled = false
			pos += dontMatch[1]
		} else if mulMatch != nil {
			if mulEnabled {
				a, err1 := strconv.Atoi(input[pos+mulMatch[2] : pos+mulMatch[3]])
				b, err2 := strconv.Atoi(input[pos+mulMatch[4] : pos+mulMatch[5]])
				if err1 == nil && err2 == nil {
					results = append(results, a*b)
				}
			}
			pos += mulMatch[1]
		} else {
			break
		}
	}

	return results, nil
}
