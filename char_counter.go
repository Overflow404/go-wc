package main

import (
	"unicode/utf8"
)

type CharCounter struct{}

func (b CharCounter) Count(input string) int {
	return utf8.RuneCountInString(input)
}
