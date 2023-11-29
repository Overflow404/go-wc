package main

import (
	"bufio"
	"os"
)

type LineCounter struct{}

func (b LineCounter) Count(filename string) (int64, error) {
	file, openError := os.Open(filename)

	if openError != nil {
		return 0, openError
	}

	scanner := bufio.NewScanner(file)

	lines := int64(0)

	for scanner.Scan() {
		lines++
	}

	if readingError := scanner.Err(); readingError != nil {
		return 0, readingError
	}

	closeError := file.Close()

	if closeError != nil {
		return 0, closeError
	}

	return lines, nil
}
