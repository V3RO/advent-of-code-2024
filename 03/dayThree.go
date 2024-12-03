package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input/input.txt
var input string

func SolvePartOne(input string) int {
	scanner := utils.NewAocScanner(input)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		validMulsRegex := regexp.MustCompile("mul\\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\\)")
		numbersRegex := regexp.MustCompile("[0-9][0-9]?[0-9]?")

		for _, mul := range validMulsRegex.FindAllString(line, -1) {
			numbers := utils.StringSliceToIntSlice(numbersRegex.FindAllString(mul, -1))
			result += numbers[0] * numbers[1]
		}
	}
	return result
}

func SolvePartTwo(input string) int {
	input = strings.Replace(input, "\n", "", -1)
	scanner := utils.NewAocScanner(input)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		// This way to hacky but tbh I can't be bothered right now, might revisit it when I have more time
		validMulsRegex := regexp.MustCompile("mul\\([0-9][0-9]?[0-9]?,[0-9][0-9]?[0-9]?\\)")
		doNotToDoRegex := regexp.MustCompile("(don't\\(\\)).*?(do\\(\\))")
		remainingDoNotRegex := regexp.MustCompile("(don't\\(\\)).*")
		numbersRegex := regexp.MustCompile("[0-9][0-9]?[0-9]?")

		doNotMuls := doNotToDoRegex.FindAllString(line, -1)
		for _, mul := range doNotMuls {
			line = strings.Replace(line, mul, "", -1)
		}

		remainingDoNot := remainingDoNotRegex.FindAllString(line, -1)
		for _, mul := range remainingDoNot {
			line = strings.Replace(line, mul, "", -1)
		}

		for _, mul := range validMulsRegex.FindAllString(line, -1) {
			numbers := utils.StringSliceToIntSlice(numbersRegex.FindAllString(mul, -1))
			result += numbers[0] * numbers[1]
		}
	}
	return result
}

func main() {
	fmt.Println("Part 1: ", SolvePartOne(input))
	fmt.Println("Part 2: ", SolvePartTwo(input))
}
