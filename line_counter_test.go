package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestLineCounter_Count_OpenSucceed(t *testing.T) {
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
			"it should count the lines if the file exist",
			args{
				filename: "assets/test/default.txt",
			},
			7145,
			nil,
		},
		{
			"it should count zero lines if the file is empty",
			args{
				filename: "assets/test/empty.txt",
			},
			0,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lineCounter := LineCounter{}

			actualLines, countError := lineCounter.Count(tt.args.filename)

			if countError != nil {
				t.Errorf("Count method returned an error: %v", countError)
			}

			if actualLines != tt.expectedBytes {
				t.Errorf("Expected %d lines, got %d lines", tt.expectedBytes, actualLines)
			}
		})
	}
}

func TestLineCounter_Count_OpenFails(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedLines int64
		expectedError error
	}{
		{
			"it should count zero lines and return error if the file does not exist",
			args{
				filename: "assets/test/_.txt",
			},
			0,
			errors.New(fmt.Sprintf("open %s: no such file or directory", "assets/test/_.txt")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lineCounter := LineCounter{}

			actualLines, actualError := lineCounter.Count(tt.args.filename)

			if actualError.Error() != tt.expectedError.Error() {
				t.Fatalf("Expected error {%s}, got {%s}", tt.expectedError, actualError)
			}

			if actualLines != tt.expectedLines {
				t.Errorf("Expected %d lines, got %d lines", tt.expectedLines, actualLines)
			}
		})
	}
}
