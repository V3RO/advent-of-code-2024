package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input/input.txt
var input string

func CheckIfLevelIsValid(report []int) bool {
	ascending := sort.SliceIsSorted(report, func(i, j int) bool {
		prev := report[i]
		curr := report[j]
		distance := utils.Abs(prev-curr) <= 3
		return !(prev < curr && distance)
	})

	descending := sort.SliceIsSorted(report, func(i, j int) bool {
		prev := report[i]
		curr := report[j]
		distance := utils.Abs(prev-curr) <= 3
		return !(prev > curr && distance)
	})

	return ascending || descending
}

func RemoveIndexFromSlice(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func SolvePartOne(input string) int {
	scanner := utils.NewAocScanner(input)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := utils.StringSliceToIntSlice(strings.Split(line, " "))

		if CheckIfLevelIsValid(report) {
			result++
		}
	}

	return result
}

func SolvePartTwo(input string) int {
	scanner := utils.NewAocScanner(input)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := utils.StringSliceToIntSlice(strings.Split(line, " "))

		if CheckIfLevelIsValid(report) {
			result++
		} else {
			for i := range report {
				dampendResult := RemoveIndexFromSlice(report, i)
				if CheckIfLevelIsValid(dampendResult) {
					result++
					break
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Part 1: ", SolvePartOne(input))
	fmt.Println("Part 2: ", SolvePartTwo(input))
}
