package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input/input.txt
var input string

func SolvePartOne(input string) int {
	scanner := utils.NewAocScanner(input)

	var listA, listB []int

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, "   ")

		a, _ := strconv.Atoi(lineSplit[0])
		b, _ := strconv.Atoi(lineSplit[1])

		listA, listB = append(listA, a), append(listB, b)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	result := 0
	for i := range listA {
		result += utils.Abs(listA[i] - listB[i])
	}

	return result
}

func SolvePartTwo(input string) int {
	scanner := utils.NewAocScanner(input)

	var listA, listB []int

	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, "   ")

		a, _ := strconv.Atoi(lineSplit[0])
		b, _ := strconv.Atoi(lineSplit[1])

		listA, listB = append(listA, a), append(listB, b)
	}

	result := 0
	for _, v := range listA {
		result += utils.Abs(v * utils.Count(
			listB,
			func(i int) bool {
				return i == v
			}))
	}

	return result
}

func main() {
	fmt.Println("Part 1: ", SolvePartOne(input))
	fmt.Println("Part 2: ", SolvePartTwo(input))
}
