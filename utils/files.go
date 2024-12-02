package utils

import (
	"bufio"
	"strings"
)

func NewAocScanner(input string) *bufio.Scanner {
	reader := strings.NewReader(input)
	return bufio.NewScanner(reader)
}
