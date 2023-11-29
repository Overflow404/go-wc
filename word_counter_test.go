package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestWordCounter_Count_OpenSucceed(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedBytes int64
		expectedError error
	}{
		{
			"it should count the words if the file exist",
			args{
				filename: "assets/test/default.txt",
			},
			58164,
			nil,
		},
		{
			"it should count zero words if the file is empty",
			args{
				filename: "assets/test/empty.txt",
			},
			0,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wordCounter := WordCounter{}

			actualWords, countError := wordCounter.Count(tt.args.filename)

			if countError != nil {
				t.Errorf("Count method returned an error: %v", countError)
			}

			if actualWords != tt.expectedBytes {
				t.Errorf("Expected %d words, got %d words", tt.expectedBytes, actualWords)
			}
		})
	}
}

func TestWordCounter_Count_OpenFails(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedWords int64
		expectedError error
	}{
		{
			"it should count zero words and return error if the file does not exist",
			args{
				filename: "assets/test/_.txt",
			},
			0,
			errors.New(fmt.Sprintf("open %s: no such file or directory", "assets/test/_.txt")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wordCounter := WordCounter{}

			actualWords, actualError := wordCounter.Count(tt.args.filename)

			if actualError.Error() != tt.expectedError.Error() {
				t.Fatalf("Expected error {%s}, got {%s}", tt.expectedError, actualError)
			}

			if actualWords != tt.expectedWords {
				t.Errorf("Expected %d words, got %d words", tt.expectedWords, actualWords)
			}
		})
	}
}
