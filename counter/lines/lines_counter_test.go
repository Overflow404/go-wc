package lines

import (
	"errors"
	"go-wc/wrapper"
	"io"
	"os"
	"testing"
)

var linesCounter = Counter{}

func TestLinesCounter_Count_ShouldReturnErrorIfOpenFails(t *testing.T) {
	openError := errors.New("open failed")

	mockedFileWrapper := wrapper.MockedFileWrapper{
		OpenFunc: func(_ string) (wrapper.File, error) {
			return nil, openError
		},
	}

	count, countError := linesCounter.Count("filename", mockedFileWrapper)

	if count != 0 {
		t.Fatalf("Expected 0 count upon an error, got {%d}", count)
	}

	if !errors.Is(countError, openError) {
		t.Fatalf("Expected error {%s}, got {%s}", openError, countError)
	}
}

func TestLinesCounter_Count_ShouldReturnCountIfOpenSucceeds(t *testing.T) {
	expectedLines := int64(7145)
	testFile, _ := os.Open("resources/test/default.txt")

	mockedFileWrapper := wrapper.MockedFileWrapper{
		OpenFunc: func(_ string) (wrapper.File, error) {
			return testFile, nil
		},
	}

	count, countError := linesCounter.Count("filename", mockedFileWrapper)

	if count != expectedLines {
		t.Fatalf("Expected {%d} lines, got {%d} lines", expectedLines, count)
	}

	if countError != nil {
		t.Fatalf("Expected no error. Got {%s}", countError)
	}

}

func TestLinesCounter_Count_ShouldReturnZeroCountIfFileIsEmpty(t *testing.T) {
	expectedLines := int64(0)
	testFile, _ := os.Open("resources/test/empty.txt")

	mockedFileWrapper := wrapper.MockedFileWrapper{
		OpenFunc: func(_ string) (wrapper.File, error) {
			return testFile, nil
		},
	}

	count, countError := linesCounter.Count("filename", mockedFileWrapper)

	if count != expectedLines {
		t.Fatalf("Expected {%d} lines, got {%d} lines", expectedLines, count)
	}

	if countError != nil {
		t.Fatalf("Expected no error. Got {%s}", countError)
	}

}

func TestLinesCounter_Count_ShouldReturnErrorIfOpenSucceedsButCloseFails(t *testing.T) {
	closeError := errors.New("close error")

	mockedFileWrapper := wrapper.MockedFileWrapper{
		OpenFunc: func(_ string) (wrapper.File, error) {
			return wrapper.MockFile{
				ReadFunc: func(p []byte) (int, error) {
					return 2, io.EOF
				},
				CloseFunc: func() error {
					return closeError
				},
			}, nil
		},
	}

	count, countError := linesCounter.Count("filename", mockedFileWrapper)

	if count != 0 {
		t.Fatalf("Expected 0 count upon an error, got {%d}", count)
	}

	if !errors.Is(countError, closeError) {
		t.Fatalf("Expected error {%s}, got {%s}", closeError, countError)
	}

}
