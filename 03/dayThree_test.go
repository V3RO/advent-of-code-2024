package main

import (
	_ "embed"
	"testing"
)

//go:embed input/test_1.txt
var test1 string

//go:embed input/test_2.txt
var test2 string

func TestPartOneExampleData(t *testing.T) {
	input := test1
	expected := 161
	result := SolvePartOne(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPartTwoExampleData(t *testing.T) {
	input := test2
	expected := 48
	result := SolvePartTwo(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
