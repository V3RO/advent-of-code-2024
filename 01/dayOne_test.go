package main

import (
	_ "embed"
	"testing"
)

//go:embed input/test.txt
var test string

func TestPartOneExampleData(t *testing.T) {
	input := test
	expected := 11
	result := SolvePartOne(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPartTwoExampleData(t *testing.T) {
	input := test
	expected := 31
	result := SolvePartTwo(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
