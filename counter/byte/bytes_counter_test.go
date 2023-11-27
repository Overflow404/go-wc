package byte

import (
	"errors"
	"go-wc/wrapper"
	"os"
	"testing"
)

var bytesCounter = Counter{}

func TestBytesCounter_Count_ShouldReturnErrorIfStatFails(t *testing.T) {
	statError := errors.New("stat failed")

	mockedFileWrapper := wrapper.MockedFileWrapper{
		StatFunc: func(_ string) (os.FileInfo, error) {
			return nil, statError
		},
	}

	count, countError := bytesCounter.Count("filename", mockedFileWrapper)

	if count != 0 {
		t.Fatalf("Expected 0 count upon an error, got {%d}", count)
	}

	if !errors.Is(countError, statError) {
		t.Fatalf("Expected error {%s}, got {%s}", statError, countError)
	}
}

func TestBytesCounter_Count_ShouldReturnCountIfStatSucceeds(t *testing.T) {
	expectedBytes := int64(342190)

	testFile, _ := os.Open("resources/test/default.txt")
	testStat, _ := testFile.Stat()

	mockedFileWrapper := wrapper.MockedFileWrapper{
		StatFunc: func(_ string) (os.FileInfo, error) {
			return testStat, nil
		},
	}

	count, countError := bytesCounter.Count("filename", mockedFileWrapper)

	if count != expectedBytes {
		t.Fatalf("Expected {%d} lines, got {%d} lines", expectedBytes, count)
	}

	if countError != nil {
		t.Fatalf("Expected no error. Got {%s}", countError)
	}
}

func TestBytesCounter_Count_ShouldReturnZeroCountIfFileIsEmpty(t *testing.T) {
	expectedBytes := int64(0)

	testFile, _ := os.Open("resources/test/empty.txt")
	testStat, _ := testFile.Stat()

	mockedFileWrapper := wrapper.MockedFileWrapper{
		StatFunc: func(_ string) (os.FileInfo, error) {
			return testStat, nil
		},
	}

	count, countError := bytesCounter.Count("filename", mockedFileWrapper)

	if count != expectedBytes {
		t.Fatalf("Expected {%d} lines, got {%d} lines", expectedBytes, count)
	}

	if countError != nil {
		t.Fatalf("Expected no error. Got {%s}", countError)
	}
}
