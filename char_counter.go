package main

import (
	"bufio"
	"os"
	"unicode/utf8"
)

type CharCounter struct{}

func (b CharCounter) Count(filename string) (int64, error) {
	file, openError := os.Open(filename)

	if openError != nil {
		return 0, openError
	}

	scanner := bufio.NewScanner(file)

	chars := int64(0)

	for scanner.Scan() {
		chars = chars + int64(utf8.RuneCountInString(scanner.Text()))
	}

	if readingError := scanner.Err(); readingError != nil {
		return 0, readingError
	}

	closeError := file.Close()

	if closeError != nil {
		return 0, closeError
	}

	return chars, nil
}
