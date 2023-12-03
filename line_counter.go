package main

import (
	"strings"
)

type LineCounter struct{}

func (b LineCounter) Count(input string) int {
	lines := strings.Split(input, "\n")

	if lines[len(lines)-1] == "" {
		return len(lines) - 1
	}

	return len(lines)
}
