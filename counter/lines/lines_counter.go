package lines

import (
	"bufio"
	"go-wc/wrapper"
)

type Counter struct{}

func (b Counter) Count(filename string, fileWrapper wrapper.FileWrapper) (int64, error) {
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
