package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestCharCounter_Count_OpenSucceed(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedChars int64
		expectedError error
	}{
		{
			"it should count the chars if the file exist",
			args{
				filename: "assets/test/default.txt",
			},
			325002,
			nil,
		},
		{
			"it should count zero chars if the file is empty",
			args{
				filename: "assets/test/empty.txt",
			},
			0,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charCounter := CharCounter{}

			actualChars, countError := charCounter.Count(tt.args.filename)

			if countError != nil {
				t.Errorf("Count method returned an error: %v", countError)
			}

			if actualChars != tt.expectedChars {
				t.Errorf("Expected %d chars, got %d chars", tt.expectedChars, actualChars)
			}
		})
	}
}

func TestCharCounter_Count_OpenFails(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name          string
		args          args
		expectedChars int64
		expectedError error
	}{
		{
			"it should count zero chars and return error if the file does not exist",
			args{
				filename: "assets/test/_.txt",
			},
			0,
			errors.New(fmt.Sprintf("open %s: no such file or directory", "assets/test/_.txt")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			charCounter := CharCounter{}

			actualChars, actualError := charCounter.Count(tt.args.filename)

			if actualError.Error() != tt.expectedError.Error() {
				t.Fatalf("Expected error {%s}, got {%s}", tt.expectedError, actualError)
			}

			if actualChars != tt.expectedChars {
				t.Errorf("Expected %d chars, got %d chars", tt.expectedChars, actualChars)
			}
		})
	}
}
