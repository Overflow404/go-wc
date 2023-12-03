package main

import (
	"strings"
)

type WordCounter struct{}

func (b WordCounter) Count(input string) int {
	return len(strings.Fields(input))
}
