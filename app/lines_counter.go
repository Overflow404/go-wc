package app

import (
	"bufio"
)

type LineCounter struct{}

func (b LineCounter) Count(filename string, fileWrapper FileWrapper) (int64, error) {
	file, openError := fileWrapper.Open(filename)

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
